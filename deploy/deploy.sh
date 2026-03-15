#!/bin/bash
# 기란장터 NAS 배포 스크립트
# Usage: ./deploy.sh

set -e

NAS_HOST="jym-nas"
NAS_DOCKER="/usr/local/bin/docker"
NAS_COMPOSE="/usr/local/bin/docker-compose"
DEPLOY_DIR="/volume1/docker/lincle-deploy"
WEB_DIR="/volume1/docker/lincle-web"
PROJECT_ROOT="$(cd "$(dirname "$0")/.." && pwd)"

echo "=== 기란장터 배포 시작 ==="

# 1. Flutter Web 빌드
echo "[1/5] Flutter Web 빌드..."
cd "$PROJECT_ROOT/frontend"
flutter build web \
  --dart-define=GOOGLE_CLIENT_ID=${GOOGLE_CLIENT_ID:-""} \
  --release
echo "  ✓ Flutter Web 빌드 완료"

# 2. 빌드 파일을 NAS로 전송
echo "[2/5] 파일 전송..."
# 배포 설정 파일
ssh $NAS_HOST "mkdir -p $DEPLOY_DIR $WEB_DIR"
scp "$PROJECT_ROOT/deploy/docker-compose.yml" "$NAS_HOST:$DEPLOY_DIR/"
scp "$PROJECT_ROOT/deploy/Caddyfile" "$NAS_HOST:$DEPLOY_DIR/"
scp "$PROJECT_ROOT/deploy/.env" "$NAS_HOST:$DEPLOY_DIR/" 2>/dev/null || echo "  (no .env file, using defaults)"

# Flutter Web 빌드 결과물
rsync -az --delete "$PROJECT_ROOT/frontend/build/web/" "$NAS_HOST:$WEB_DIR/"

# Backend 소스 (Docker build on NAS)
rsync -az --delete \
  --exclude='data/' --exclude='uploads/' --exclude='*.db*' \
  "$PROJECT_ROOT/backend/" "$NAS_HOST:$DEPLOY_DIR/backend/"
echo "  ✓ 파일 전송 완료"

# 3. NAS에서 Docker 빌드 + 시작
echo "[3/5] Docker 빌드 및 시작..."
ssh $NAS_HOST "cd $DEPLOY_DIR && echo '3945A9595088a!@#' | sudo -S $NAS_COMPOSE build --no-cache lincle-api 2>&1" | tail -5
ssh $NAS_HOST "cd $DEPLOY_DIR && echo '3945A9595088a!@#' | sudo -S $NAS_COMPOSE up -d 2>&1" | tail -5
echo "  ✓ Docker 컨테이너 시작됨"

# 4. 헬스체크
echo "[4/5] 헬스체크..."
sleep 5
HEALTH=$(ssh $NAS_HOST "wget -qO- http://localhost:18090/health 2>/dev/null || echo 'FAIL'")
if echo "$HEALTH" | grep -q '"status":"ok"'; then
    echo "  ✓ API 정상 가동"
else
    echo "  ✗ API 헬스체크 실패: $HEALTH"
    exit 1
fi

# 5. 완료
echo "[5/5] 배포 완료!"
echo ""
echo "  내부: http://$(ssh $NAS_HOST 'hostname -I' | awk '{print $1}'):18090"
echo "  외부: https://giranjt.com (Cloudflare Tunnel 설정 필요)"
echo ""
