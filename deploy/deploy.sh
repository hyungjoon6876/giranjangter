#!/bin/bash
# 기란장터 NAS 배포 스크립트
# Usage: ./deploy.sh

set -e

NAS_HOST="jym-nas"
DEPLOY_DIR="/volume1/docker/lincle-deploy"
PROJECT_ROOT="$(cd "$(dirname "$0")/.." && pwd)"

echo "=== 기란장터 배포 시작 ==="

# 1. 파일 전송
echo "[1/4] 파일 전송..."
ssh $NAS_HOST "mkdir -p $DEPLOY_DIR/web $DEPLOY_DIR/backend $DEPLOY_DIR/shared"

# 배포 설정 파일
scp -O "$PROJECT_ROOT/docker-compose.yml" "$NAS_HOST:$DEPLOY_DIR/"
scp -O "$PROJECT_ROOT/Caddyfile" "$NAS_HOST:$DEPLOY_DIR/"

# Backend 소스 — tar로 묶어서 전송 (rsync 호환성 문제 회피)
echo "  → Backend 전송 중..."
tar czf - -C "$PROJECT_ROOT/backend" \
  --exclude='data' --exclude='uploads' --exclude='*.db*' . \
  | ssh $NAS_HOST "cd $DEPLOY_DIR/backend && tar xzf -"

# Web 소스 — tar로 묶어서 전송
echo "  → Web 전송 중..."
tar czf - -C "$PROJECT_ROOT/web" \
  --exclude='node_modules' --exclude='.next' . \
  | ssh $NAS_HOST "cd $DEPLOY_DIR/web && tar xzf -"

# Shared 디자인 토큰
echo "  → Shared 전송 중..."
scp -O "$PROJECT_ROOT/shared/design-tokens.json" "$NAS_HOST:$DEPLOY_DIR/shared/"

echo "  ✓ 파일 전송 완료"

# 2. NAS에서 Docker 빌드 + 시작
echo "[2/4] Docker 빌드 및 시작..."
NAS_COMPOSE="/usr/local/bin/docker-compose"
ssh $NAS_HOST "cd $DEPLOY_DIR && $NAS_COMPOSE build --no-cache lincle-api lincle-web 2>&1" | tail -10
ssh $NAS_HOST "cd $DEPLOY_DIR && $NAS_COMPOSE up -d 2>&1" | tail -5
echo "  ✓ Docker 컨테이너 시작됨"

# 3. 헬스체크
echo "[3/4] 헬스체크..."
sleep 15

# API 헬스체크
HEALTH=$(ssh $NAS_HOST "wget -qO- http://localhost:80/health 2>/dev/null || echo 'FAIL'")
if echo "$HEALTH" | grep -q '"status":"ok"'; then
    echo "  ✓ API 정상 가동"
else
    echo "  ✗ API 헬스체크 실패: $HEALTH"
    echo "  로그: ssh $NAS_HOST 'cd $DEPLOY_DIR && $NAS_COMPOSE logs --tail=20'"
    exit 1
fi

# Web 헬스체크
WEB_CHECK=$(ssh $NAS_HOST "wget --spider -q http://localhost:80/ 2>&1 && echo 'OK' || echo 'FAIL'")
if [ "$WEB_CHECK" = "OK" ]; then
    echo "  ✓ Web 정상 가동"
else
    echo "  ⚠ Web 응답 확인 필요"
    ssh $NAS_HOST "cd $DEPLOY_DIR && $NAS_COMPOSE logs --tail=10 lincle-web" 2>&1 | tail -5
fi

# 4. 완료
echo "[4/4] 배포 완료!"
echo ""
NAS_IP=$(ssh $NAS_HOST 'hostname -I' | awk '{print $1}')
echo "  내부: http://${NAS_IP}"
echo "  외부: https://giranjt.com (Cloudflare Tunnel 설정 필요)"
echo ""
echo "  상태: ssh $NAS_HOST 'cd $DEPLOY_DIR && /usr/local/bin/docker-compose ps'"
echo "  로그: ssh $NAS_HOST 'cd $DEPLOY_DIR && /usr/local/bin/docker-compose logs -f'"
