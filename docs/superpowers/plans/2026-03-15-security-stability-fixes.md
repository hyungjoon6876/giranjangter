# 백엔드 보안/안정성 보류 이슈 구현 계획

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 보안/안정성 리뷰에서 보류된 6개 이슈를 모두 해결한다.

**Architecture:** DB 스키마 변경(마이그레이션) → 핸들러 수정 → 테스트 → 커밋 순서. 리프레시 토큰은 Redis 없이 PostgreSQL `refresh_tokens` 테이블로 관리. favorite_count는 서브쿼리 기반으로 전환.

**Tech Stack:** Go 1.25, Gin, PostgreSQL 16, pgx/v5

---

## Task 1: 리프레시 토큰 서버사이드 관리 (L-2)

Redis 없이 PostgreSQL에 `refresh_tokens` 테이블을 만들어 토큰 발급/검증/폐기를 관리한다.

**Files:**
- Create: `backend/db/migrations/005_refresh_tokens.sql`
- Modify: `backend/cmd/server/handlers_auth.go`
- Modify: `backend/cmd/server/main.go` (로그아웃 라우트 추가)

### 마이그레이션

- [ ] **Step 1: 마이그레이션 파일 생성**

```sql
-- 005: Refresh token server-side storage (DB-based, no Redis)
CREATE TABLE IF NOT EXISTS refresh_tokens (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_hash TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_refresh_tokens_user ON refresh_tokens(user_id);
CREATE INDEX IF NOT EXISTS idx_refresh_tokens_hash ON refresh_tokens(token_hash);
```

- [ ] **Step 2: repository/db.go에 마이그레이션 경로 추가**

`005_refresh_tokens.sql`을 마이그레이션 paths에 추가.

### 핸들러 수정

- [ ] **Step 3: handleLogin — 리프레시 토큰 DB 저장**

로그인 성공 후 리프레시 토큰을 생성하고, SHA-256 해시를 `refresh_tokens` 테이블에 저장한다.

```go
import "crypto/sha256"
import "encoding/hex"

// 토큰 생성 후
tokenHash := sha256.Sum256([]byte(refreshToken))
hashStr := hex.EncodeToString(tokenHash[:])
db.Exec("INSERT INTO refresh_tokens (id, user_id, token_hash, expires_at) VALUES ($1, $2, $3, $4)",
    uuid.New().String(), userID, hashStr, time.Now().Add(cfg.JWTRefreshTTL))
```

- [ ] **Step 4: handleRefresh — DB에서 토큰 존재 확인**

리프레시 시 토큰 해시가 DB에 있는지 확인. 있으면 삭제(rotation) 후 새 토큰 발급+저장.

```go
tokenHash := sha256.Sum256([]byte(req.RefreshToken))
hashStr := hex.EncodeToString(tokenHash[:])
var tokenID string
err := db.QueryRow("SELECT id FROM refresh_tokens WHERE token_hash = $1 AND expires_at > NOW()", hashStr).Scan(&tokenID)
if err != nil {
    // 토큰 없음 또는 만료 → 401
    return
}
// 기존 토큰 삭제 (rotation)
db.Exec("DELETE FROM refresh_tokens WHERE id = $1", tokenID)
// 새 토큰 발급 + DB 저장
```

- [ ] **Step 5: 로그아웃 핸들러 추가**

```go
func handleLogout(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := middleware.GetUserID(c)
        // 해당 유저의 모든 리프레시 토큰 삭제
        db.Exec("DELETE FROM refresh_tokens WHERE user_id = $1", userID)
        c.Status(http.StatusNoContent)
    }
}
```

main.go에 `write.POST("/auth/logout", handleLogout(db))` 라우트 추가.

- [ ] **Step 6: 만료 토큰 정리 — 주기적 cleanup**

서버 시작 시 만료된 토큰 삭제:
```go
db.Exec("DELETE FROM refresh_tokens WHERE expires_at < NOW()")
```

- [ ] **Step 7: 빌드 확인 + 커밋**

```bash
cd backend && go build ./cmd/server/ && go test ./...
git add -A && git commit -m "security: 리프레시 토큰 DB 기반 관리 — 발급/검증/폐기/로그아웃"
```

---

## Task 2: favorite_count 레이스 컨디션 해결 (H-3)

denormalized `favorite_count` 컬럼 대신 서브쿼리로 실시간 계산한다.

**Files:**
- Modify: `backend/cmd/server/handlers_listing.go` (favorite/unfavorite + 목록/상세 쿼리)

### 핸들러 수정

- [ ] **Step 1: handleFavoriteListing 수정**

favorite_count UPDATE 제거. INSERT만 수행.

```go
func handleFavoriteListing(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        userID := middleware.GetUserID(c)
        // 매물 존재 확인 (M-9 해결)
        var exists bool
        if err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM listings WHERE id = $1 AND deleted_at IS NULL)", id).Scan(&exists); err != nil || !exists {
            c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "NOT_FOUND", "message": "매물을 찾을 수 없습니다."}})
            return
        }
        db.Exec("INSERT INTO favorites (id, user_id, listing_id) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING",
            uuid.New().String(), userID, id)
        c.Status(http.StatusNoContent)
    }
}
```

- [ ] **Step 2: handleUnfavoriteListing 수정**

favorite_count UPDATE 제거. DELETE만 수행.

```go
func handleUnfavoriteListing(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        userID := middleware.GetUserID(c)
        db.Exec("DELETE FROM favorites WHERE user_id = $1 AND listing_id = $2", userID, id)
        c.Status(http.StatusNoContent)
    }
}
```

- [ ] **Step 3: handleListListings 쿼리에서 favorite_count를 서브쿼리로 변경**

기존: `l.favorite_count`
변경: `(SELECT COUNT(*) FROM favorites f WHERE f.listing_id = l.id) as favorite_count`

handleGetListing에서도 동일하게 변경.

- [ ] **Step 4: chat_count, view_count도 동일 패턴 검토**

view_count는 실시간 계산이 비효율적이므로 유지. chat_count도 유지.
favorite_count만 서브쿼리로 전환.

- [ ] **Step 5: 빌드 확인 + 커밋**

```bash
cd backend && go build ./cmd/server/ && go test ./...
git add -A && git commit -m "fix: favorite_count 레이스 컨디션 해결 — 서브쿼리 기반 실시간 계산"
```

---

## Task 3: 채팅 메시지 커서 페이지네이션 (L-3)

커서 기반 페이지네이션을 추가하여 "이전 메시지 더보기" 기능을 지원한다.

**Files:**
- Modify: `backend/cmd/server/handlers_chat.go` (handleListMessages)
- Modify: `frontend/lib/shared/api/api_client.dart` (getMessages에 cursor 파라미터 추가)

### 백엔드

- [ ] **Step 1: handleListMessages에 cursor/limit 파라미터 추가**

```go
func handleListMessages(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        chatID := c.Param("chatId")
        userID := middleware.GetUserID(c)
        cursor := c.Query("cursor")  // 마지막 메시지의 sent_at
        limitStr := c.DefaultQuery("limit", "50")
        limit, _ := strconv.Atoi(limitStr)
        if limit <= 0 || limit > 100 { limit = 50 }

        // 참여자 확인 (기존)
        ...

        // 커서 기반 쿼리
        query := "SELECT id, sender_user_id, message_type, body_text, metadata_json, sent_at FROM chat_messages WHERE chat_room_id = $1 AND deleted_at IS NULL"
        args := []interface{}{chatID}
        paramIdx := 2

        if cursor != "" {
            query += fmt.Sprintf(" AND sent_at < $%d", paramIdx)
            args = append(args, cursor)
            paramIdx++
        }

        query += fmt.Sprintf(" ORDER BY sent_at DESC LIMIT %d", limit+1)
        rows, err := db.Query(query, args...)
        ...

        // hasMore 판별
        hasMore := len(msgs) > limit
        if hasMore { msgs = msgs[:limit] }

        c.JSON(http.StatusOK, gin.H{
            "data": msgs,
            "cursor": gin.H{
                "next": nextCursor,
                "hasMore": hasMore,
            },
        })
    }
}
```

### 프론트엔드

- [ ] **Step 2: api_client.dart의 getMessages에 cursor 파라미터 추가**

```dart
Future<Map<String, dynamic>> getMessages(String chatId, {String? cursor}) async {
    final params = <String, dynamic>{};
    if (cursor != null) params['cursor'] = cursor;
    final res = await dio.get('/chats/$chatId/messages', queryParameters: params);
    return res.data;
}
```

- [ ] **Step 3: 빌드 확인 + 커밋**

```bash
cd backend && go build ./cmd/server/
cd ../frontend && flutter analyze 2>&1 | grep error
git add -A && git commit -m "feat: 채팅 메시지 커서 페이지네이션 — cursor/limit/hasMore"
```

---

## Task 4: Google 토큰 로컬 검증 (H-7)

`google.golang.org/api/idtoken` 패키지를 사용하여 Google ID Token을 로컬에서 검증한다.

**Files:**
- Modify: `backend/internal/oauth/google.go`
- Modify: `backend/go.mod` (의존성 추가)

- [ ] **Step 1: idtoken 패키지 설치**

```bash
cd backend && go get google.golang.org/api/idtoken
```

- [ ] **Step 2: VerifyGoogleIDToken 재작성**

```go
package oauth

import (
    "context"
    "fmt"
    "google.golang.org/api/idtoken"
)

type GoogleTokenInfo struct {
    Sub   string
    Email string
    Name  string
}

func VerifyGoogleIDToken(token string, expectedClientIDs []string) (*GoogleTokenInfo, error) {
    for _, clientID := range expectedClientIDs {
        if clientID == "" { continue }
        payload, err := idtoken.Validate(context.Background(), token, clientID)
        if err != nil { continue }
        sub, _ := payload.Claims["sub"].(string)
        email, _ := payload.Claims["email"].(string)
        name, _ := payload.Claims["name"].(string)
        if sub == "" { return nil, fmt.Errorf("google token missing sub claim") }
        return &GoogleTokenInfo{Sub: sub, Email: email, Name: name}, nil
    }
    return nil, fmt.Errorf("google token validation failed for all client IDs")
}
```

- [ ] **Step 3: handlers_auth.go에서 호출 부분 확인 — 변경 불필요**

`VerifyGoogleIDToken`의 반환 타입이 동일하면 핸들러 변경 없음. `info.Sub` 접근이 동일한지 확인.

- [ ] **Step 4: 빌드 확인 + 커밋**

```bash
cd backend && go build ./cmd/server/ && go test ./...
git add -A && git commit -m "security: Google ID Token 로컬 검증 — idtoken 패키지 기반"
```

---

## Task 5: 이미지 소유권 관리 (M-3)

업로드된 이미지를 DB에 기록하고, 매물 등록 시 소유권을 검증한다.

**Files:**
- Create: `backend/db/migrations/006_image_ownership.sql`
- Modify: `backend/cmd/server/handlers_upload.go`
- Modify: `backend/cmd/server/handlers_listing.go` (createListing)

- [ ] **Step 1: 마이그레이션 — uploaded_images 테이블**

```sql
-- 006: Track uploaded image ownership
CREATE TABLE IF NOT EXISTS uploaded_images (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    filename TEXT NOT NULL,
    url TEXT NOT NULL,
    content_type TEXT,
    size_bytes BIGINT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_uploaded_images_user ON uploaded_images(user_id);
```

- [ ] **Step 2: handleUploadImage — DB에 이미지 기록**

업로드 성공 후 `uploaded_images` 테이블에 INSERT:
```go
db.Exec("INSERT INTO uploaded_images (id, user_id, filename, url, content_type, size_bytes) VALUES ($1, $2, $3, $4, $5, $6)",
    imageID, userID, filename, url, contentType, file.Size)
```

handlers_upload.go의 함수 시그니처에 `db *sql.DB` 추가 필요. main.go에서도 `handleUploadImage(cfg, db)`로 변경.

- [ ] **Step 3: handleCreateListing — imageIds 소유권 검증**

```go
if len(req.ImageIDs) > 0 {
    for _, imgID := range req.ImageIDs {
        var ownerID string
        err := db.QueryRow("SELECT user_id FROM uploaded_images WHERE id = $1", imgID).Scan(&ownerID)
        if err != nil || ownerID != userID {
            c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "FORBIDDEN", "message": "본인이 업로드한 이미지만 사용할 수 있습니다."}})
            return
        }
    }
    // listing_images 테이블에 연결
}
```

- [ ] **Step 4: 빌드 확인 + 커밋**

```bash
cd backend && go build ./cmd/server/ && go test ./...
git add -A && git commit -m "security: 이미지 업로드 DB 기록 + 매물 등록 시 소유권 검증"
```

---

## Task 6: 찜하기 매물 존재 확인 (M-9)

Task 2에서 이미 해결됨 (handleFavoriteListing에 EXISTS 체크 추가). 별도 작업 불필요.

---

## 실행 순서

1. Task 1 (리프레시 토큰) — 가장 중요, 보안 핵심
2. Task 2 (favorite_count) — 데이터 정합성
3. Task 3 (채팅 페이지네이션) — UX 개선
4. Task 4 (Google 토큰 로컬 검증) — 의존성 추가, 단독 실행
5. Task 5 (이미지 소유권) — 마이그레이션 + 핸들러

---

## 최종 검증

모든 태스크 완료 후:
```bash
cd backend && go build ./cmd/server/ && go test ./...
# 서버 시작 후 마이그레이션 정상 적용 확인
ENV=development DATABASE_URL=... go run ./cmd/server/
```
