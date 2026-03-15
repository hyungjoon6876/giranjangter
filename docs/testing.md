# 테스팅 전략

## 실행 명령

```bash
# Backend 전체 테스트
cd backend && go test ./...

# Backend 특정 패키지
cd backend && go test ./internal/guard/...

# Frontend 전체 테스트
cd frontend && flutter test

# Frontend 특정 파일
cd frontend && flutter test test/xxx_test.dart
```

## 테스트 유형

| 유형 | Backend | Frontend | 용도 |
|------|---------|----------|------|
| 단위 테스트 | `*_test.go` (같은 패키지) | `test/**/*_test.dart` | 개별 함수/클래스 검증 |
| 통합 테스트 | 핸들러 + DB 테스트 | `integration_test/` | 컴포넌트 간 연동 검증 |
| 위젯 테스트 | N/A | `test/widget/` | Flutter UI 컴포넌트 검증 |

## 파일 네이밍

- **Go**: `<원본파일>_test.go` — 같은 디렉토리에 위치 (예: `listing_guard_test.go`)
- **Dart**: `<원본파일>_test.dart` — `test/` 디렉토리에 미러링

## 테스트 우선 대상

다음 영역은 반드시 테스트를 작성한다:

1. **Guard (상태 전이)** — 허용/금지 전이를 모두 검증
2. **Middleware (인증/인가)** — JWT 검증, 역할 체크
3. **핸들러 입력 검증** — 필수 필드 누락, 잘못된 형식
4. **도메인 로직** — 비즈니스 규칙 (예: 자기 거래글에 예약 불가)

## 테스트 작성 원칙

- 테스트 함수명은 `Test<함수명>_<시나리오>` (Go) 또는 `<설명>` (Dart)
- 하나의 테스트는 하나의 동작만 검증
- 테스트 데이터는 테스트 함수 내에서 생성 (전역 상태 공유 금지)
- DB 테스트는 트랜잭션으로 감싸서 격리

## 모킹 컨벤션

- **Go**: 인터페이스 기반 모킹 또는 테스트용 SQLite in-memory DB
- **Dart**: Riverpod `ProviderScope.overrides`로 Provider 교체
