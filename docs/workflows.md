# 개발 및 배포 워크플로우

## 로컬 개발 환경

### Backend

```bash
cd backend
cp ../.env.example .env     # 환경변수 설정
go run ./cmd/server/         # 서버 시작 (기본 :8080)
```

### Frontend

```bash
cd frontend
flutter pub get              # 의존성 설치
dart run build_runner build --delete-conflicting-outputs  # 코드 생성
flutter run -d chrome        # 웹 개발 서버
flutter run                  # 모바일 에뮬레이터
```

### 전체 스택 (Docker)

```bash
docker compose up --build    # API + Caddy 실행
```

Caddy 라우팅:
- `/api/*`, `/sse/*`, `/health` → Go API 서버
- `/uploads/*` → 파일 서빙
- `/*` → Flutter Web 정적 파일

## 배포

### 대상 환경
- **개발**: 로컬 Docker Compose
- **운영**: 개인 NAS (2GB RAM) + Cloudflare DNS

### 배포 절차

```bash
# 1. 빌드
cd frontend && flutter build web
cd backend && docker build -t lincle-api .

# 2. 배포 (NAS)
docker compose up -d
```

## 데이터베이스 마이그레이션

```
backend/db/migrations/
├── 001_initial_schema.sql
├── 002_xxx.sql
└── ...
```

- 새 마이그레이션: 다음 번호로 파일 추가
- 기존 마이그레이션 파일 수정 금지
- 롤백 SQL은 같은 파일 하단에 주석으로 포함

## 코드 생성 (Flutter)

freezed, json_serializable 모델 변경 후:

```bash
cd frontend && dart run build_runner build --delete-conflicting-outputs
```

생성 파일 (`*.g.dart`, `*.freezed.dart`)은 커밋에 포함한다.
