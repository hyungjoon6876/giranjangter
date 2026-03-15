# 에러 핸들링 및 로깅 컨벤션

## Backend (Go)

### 에러 응답 형식

모든 에러 응답은 통일된 구조를 따른다:

```json
{
  "error": {
    "code": "ERROR_CODE",
    "message": "사람이 읽을 수 있는 설명"
  }
}
```

### 에러 코드 매핑

| HTTP 상태 | 에러 코드 | 용도 |
|-----------|-----------|------|
| 400 | `VALIDATION_ERROR` | 요청 파싱/검증 실패 |
| 401 | `UNAUTHORIZED` | JWT 없음 또는 만료 |
| 403 | `FORBIDDEN` | 권한 부족 (RBAC) |
| 404 | `NOT_FOUND` | 리소스 미존재 |
| 409 | `CONFLICT` | 상태 전이 충돌, 중복 |
| 500 | `INTERNAL_ERROR` | 서버 내부 오류 |

### 핸들러 에러 처리 패턴

```go
// 입력 검증
if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
        "error": gin.H{"code": "VALIDATION_ERROR", "message": err.Error()},
    })
    return
}

// DB 에러
result, err := db.ExecContext(ctx, query, args...)
if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
        "error": gin.H{"code": "INTERNAL_ERROR", "message": "처리 중 오류가 발생했습니다"},
    })
    return
}

// 상태 전이 에러
if err := guard.ValidateListingTransition(current, target); err != nil {
    c.JSON(http.StatusConflict, gin.H{
        "error": gin.H{"code": "CONFLICT", "message": err.Error()},
    })
    return
}
```

### DB 트랜잭션 패턴

```go
tx, err := db.BeginTx(ctx, nil)
if err != nil { /* 에러 응답 */ }
defer tx.Rollback()

// 쿼리 실행...

if err := tx.Commit(); err != nil { /* 에러 응답 */ }
```

---

## Frontend (Dart/Flutter)

### API 호출 에러 처리

```dart
try {
  final response = await api.createListing(data);
  // 성공 처리
} on DioException catch (e) {
  // 에러 코드별 분기
  final code = e.response?.data['error']?['code'];
  // UI 피드백
}
```

### 토큰 갱신

Dio 인터셉터가 401 응답 시 자동으로 토큰을 갱신한다.
갱신 실패 시 로그인 화면으로 리다이렉트.

### 로딩/에러 상태

- `FutureProvider`: loading/error/data 3상태 자동 관리
- `ConsumerStatefulWidget`: `setState`로 `_isLoading` 플래그 관리

---

## 로깅 (향후 도입)

현재 `fmt.Printf` 기반. 향후 구조화 로깅 도입 시:

- **레벨**: DEBUG, INFO, WARN, ERROR
- **형식**: JSON 구조화 로그
- **민감 정보**: 로그에 JWT, 비밀번호, 개인정보 포함 금지
