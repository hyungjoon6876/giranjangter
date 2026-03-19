# 아키텍처 패턴 규칙

> 에이전트가 코드를 작성할 때 따라야 하는 설계 패턴.
> 자동 검증이 어려운 규칙을 마크다운으로 정의한다.

## 상태 전이는 Guard를 통해서만

- 상태 전이는 `internal/guard/` 패키지의 검증 함수를 거쳐야 한다
- 핸들러에서 상태 필드를 직접 UPDATE하지 않는다
- 새로운 상태 전이는 `AllowedXxxTransitions` 맵에 추가 후 사용

## Handler Factory 패턴

- 핸들러는 `func handleXxx(db *sql.DB) gin.HandlerFunc` 팩토리 함수로 정의
- 의존성(db, config, broker)은 클로저로 캡처
- 글로벌 변수를 통한 의존성 접근 금지

## API 에러 응답 형식 통일

- 에러 응답: `{"error": {"code": "XXX_ERROR", "message": "..."}}`
- 코드는 UPPER_SNAKE_CASE (예: `VALIDATION_ERROR`, `NOT_FOUND`, `UNAUTHORIZED`)
- 메시지는 사용자에게 보여줄 수 있는 한국어 문구
