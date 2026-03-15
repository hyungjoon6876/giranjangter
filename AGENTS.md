# 에이전트 운영 가이드

## 역할

이 레포에서 에이전트는 **풀스택 개발자**로 동작한다.
Go 백엔드와 Flutter 프론트엔드 모두 작업한다.

## 태스크 실행 루프

1. **계획** — `tasks/todo.md`에 작업 항목 작성
2. **구현** — 코드 작성, 한 단계씩 진행
3. **검증** — 테스트 실행, 동작 확인
4. **문서화** — 변경 사항 요약, 필요시 docs/ 업데이트

## 자율성 경계

### 자율적으로 할 수 있는 것
- 버그 수정 및 테스트 작성
- 기존 패턴을 따르는 CRUD 구현
- 리팩토링 (동작 변경 없음)
- 린터/포맷터 오류 수정
- 테스트 실행 및 결과 분석

### 반드시 물어봐야 하는 것
- 새로운 상태 전이 규칙 추가 (Guard 변경)
- RBAC 권한 변경
- DB 스키마 변경 (마이그레이션)
- 새로운 외부 의존성 추가
- API 엔드포인트 추가/삭제
- 인프라 설정 변경 (Docker, Caddy)

## 파일별 주의사항

| 파일/디렉토리 | 규칙 |
|--------------|------|
| `backend/internal/guard/` | 상태 전이 로직 — 변경 시 반드시 확인 |
| `backend/db/migrations/` | 신규 마이그레이션만 추가, 기존 파일 수정 금지 |
| `docs/RBAC_ACTION_MATRIX.md` | 권한 변경은 사용자 승인 필요 |
| `docs/STARTER_DDL.md` | 스키마 참조 문서 — 마이그레이션과 동기화 필요 |
| `.env`, `.env.example` | 시크릿 노출 금지 |

## 코드 작성 원칙

- **Backend**: 새 핸들러는 `handlers_xxx.go` 패턴을 따른다
- **Frontend**: 새 기능은 `lib/features/xxx/` 하위에 생성한다
- **에러 응답**: 항상 `gin.H{"error": gin.H{"code": "...", "message": "..."}}` 형식
- **Provider**: Riverpod Provider는 `xxxProvider` 네이밍
- **상태 전이**: `AllowedXxxTransitions` 맵에 정의된 것만 허용
