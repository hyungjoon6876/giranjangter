#!/bin/bash
# 기란장터 NAS 배포 스크립트
# Usage: ./deploy.sh

set -e

NAS_HOST="jym-nas"
NAS_DOCKER="/usr/local/bin/docker"
NAS_COMPOSE="/usr/local/bin/docker-compose"
DEPLOY_DIR="/volume1/docker/lincle-deploy"
PROJECT_ROOT="$(cd "$(dirname "$0")/.." && pwd)"

echo "=== 기란장터 배포 시작 ==="

# 1. 파일 전송
echo "[1/4] 파일 전송..."
ssh $NAS_HOST "mkdir -p $DEPLOY_DIR"

# 배포 설정 파일
scp "$PROJECT_ROOT/docker-compose.yml" "$NAS_HOST:$DEPLOY_DIR/"
scp "$PROJECT_ROOT/Caddyfile" "$NAS_HOST:$DEPLOY_DIR/"
scp "$PROJECT_ROOT/deploy/.env" "$NAS_HOST:$DEPLOY_DIR/" 2>/dev/null || echo "  (no .env file, using defaults)"

# Backend 소스 (Docker build on NAS)
rsync -az --delete \
  --exclude='data/' --exclude='uploads/' --exclude='*.db*' \
  "$PROJECT_ROOT/backend/" "$NAS_HOST:$DEPLOY_DIR/backend/"

# Web 소스 (Docker build on NAS)
rsync -az --delete \
  --exclude='node_modules/' --exclude='.next/' \
  "$PROJECT_ROOT/web/" "$NAS_HOST:$DEPLOY_DIR/web/"

echo "  ✓ 파일 전송 완료"

# 2. NAS에서 Docker 빌드 + 시작
echo "[2/4] Docker 빌드 및 시작..."
ssh $NAS_HOST "cd $DEPLOY_DIR && $NAS_COMPOSE build --no-cache lincle-api lincle-web 2>&1" | tail -5
ssh $NAS_HOST "cd $DEPLOY_DIR && $NAS_COMPOSE up -d 2>&1" | tail -5
echo "  ✓ Docker 컨테이너 시작됨"

# 3. 헬스체크
echo "[3/4] 헬스체크..."
sleep 10
HEALTH=$(ssh $NAS_HOST "wget -qO- http://localhost:80/health 2>/dev/null || echo 'FAIL'")
if echo "$HEALTH" | grep -q '"status":"ok"'; then
    echo "  ✓ API 정상 가동"
else
    echo "  ✗ API 헬스체크 실패: $HEALTH"
    echo "  로그 확인: ssh $NAS_HOST 'cd $DEPLOY_DIR && $NAS_COMPOSE logs --tail=20'"
    exit 1
fi

# Web 헬스체크
WEB_STATUS=$(ssh $NAS_HOST "wget -qO- -S http://localhost:80/ 2>&1 | head -1 || echo 'FAIL'")
if echo "$WEB_STATUS" | grep -q "200"; then
    echo "  ✓ Web 정상 가동"
else
    echo "  ⚠ Web 응답 확인 필요: $WEB_STATUS"
fi

# 4. 완료
echo "[4/4] 배포 완료!"
echo ""
echo "  내부: http://$(ssh $NAS_HOST 'hostname -I' | awk '{print $1}')"
echo "  외부: https://giranjt.com (Cloudflare Tunnel 설정 필요)"
echo ""
