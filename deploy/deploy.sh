#!/bin/bash
# 기란장터 NAS 무중단 배포 스크립트
# 롤링 업데이트: 이미지 빌드 → 서비스별 순차 재시작 → 헬스체크
# Caddy가 upstream 재시도로 다운타임 최소화
#
# Usage: ./deploy.sh

set -e

NAS_HOST="jym-nas"
DEPLOY_DIR="/volume1/docker/lincle-deploy"
NAS_COMPOSE="sudo PATH=/usr/local/bin:\$PATH /usr/local/bin/docker-compose"
PROJECT_ROOT="$(cd "$(dirname "$0")/.." && pwd)"

echo "=== 기란장터 무중단 배포 시작 ==="

# 1. 파일 전송
echo "[1/5] 파일 전송..."
ssh $NAS_HOST "mkdir -p $DEPLOY_DIR/web $DEPLOY_DIR/backend $DEPLOY_DIR/shared"

scp -O "$PROJECT_ROOT/docker-compose.yml" "$NAS_HOST:$DEPLOY_DIR/"
scp -O "$PROJECT_ROOT/Caddyfile" "$NAS_HOST:$DEPLOY_DIR/"
scp -O "$PROJECT_ROOT/shared/design-tokens.json" "$NAS_HOST:$DEPLOY_DIR/shared/"

echo "  → Backend..."
tar czf - -C "$PROJECT_ROOT/backend" \
  --exclude='data' --exclude='uploads' --exclude='*.db*' . \
  | ssh $NAS_HOST "cd $DEPLOY_DIR/backend && tar xzf -"

echo "  → Web..."
tar czf - -C "$PROJECT_ROOT/web" \
  --exclude='node_modules' --exclude='.next' --exclude='test-results' . \
  | ssh $NAS_HOST "cd $DEPLOY_DIR/web && tar xzf -"

echo "  ✓ 파일 전송 완료"

# 2. 이미지 빌드 (서비스 유지한 채로 빌드만)
echo "[2/5] Docker 이미지 빌드 (서비스 유지)..."
ssh $NAS_HOST "cd $DEPLOY_DIR && $NAS_COMPOSE build --no-cache lincle-api lincle-web 2>&1" | tail -5
echo "  ✓ 이미지 빌드 완료"

# 3. Caddy 먼저 업데이트 (Caddyfile 변경 반영 — 재시도 설정)
echo "[3/5] Caddy 업데이트..."
ssh $NAS_HOST "cd $DEPLOY_DIR && $NAS_COMPOSE up -d --no-deps caddy 2>&1" | tail -3
sleep 2
echo "  ✓ Caddy 업데이트 완료"

# 4. 롤링 재시작: API → 대기 → Web
echo "[4/5] 롤링 재시작..."

# 4a. API 재시작 (Caddy가 재시도하므로 짧은 다운타임만)
echo "  → API 재시작..."
ssh $NAS_HOST "cd $DEPLOY_DIR && $NAS_COMPOSE up -d --no-deps lincle-api 2>&1" | tail -3

# API 헬스 대기
echo "  → API 헬스체크 대기..."
for i in $(seq 1 20); do
  HEALTH=$(ssh $NAS_HOST "wget -qO- http://localhost:8080/health 2>/dev/null || echo 'FAIL'")
  if echo "$HEALTH" | grep -q '"status":"ok"'; then
    echo "  ✓ API healthy (${i}0초)"
    break
  fi
  if [ "$i" = "20" ]; then
    echo "  ✗ API 헬스체크 실패"
    ssh $NAS_HOST "cd $DEPLOY_DIR && $NAS_COMPOSE logs --tail=10 lincle-api 2>&1"
    exit 1
  fi
  sleep 10
done

# 4b. Web 재시작 (Caddy가 재시도하므로 짧은 다운타임만)
echo "  → Web 재시작..."
ssh $NAS_HOST "cd $DEPLOY_DIR && $NAS_COMPOSE up -d --no-deps lincle-web 2>&1" | tail -3

# Web 헬스 대기
echo "  → Web 헬스체크 대기..."
for i in $(seq 1 15); do
  WEB=$(ssh $NAS_HOST "wget -qO- http://localhost:18090/ 2>/dev/null | head -c 20 || echo 'FAIL'")
  if echo "$WEB" | grep -q "DOCTYPE\|html"; then
    echo "  ✓ Web healthy (${i}0초)"
    break
  fi
  if [ "$i" = "15" ]; then
    echo "  ⚠ Web 응답 지연 — 로그 확인 필요"
  fi
  sleep 10
done

echo "  ✓ 롤링 재시작 완료"

# 5. 최종 확인
echo "[5/5] 최종 확인..."
ssh $NAS_HOST "cd $DEPLOY_DIR && $NAS_COMPOSE ps 2>&1"
echo ""
NAS_IP=$(ssh $NAS_HOST 'hostname -I' | awk '{print $1}')
echo "  내부: http://${NAS_IP}:18090"
echo "  외부: https://giranjt.com"
echo ""
echo "=== 무중단 배포 완료 ==="
