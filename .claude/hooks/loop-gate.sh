#!/bin/bash
# harness-loop: 자율 실행 루프 Stop 훅
# 자동 생성됨 — /harness loop init 재실행으로 재생성 가능

INPUT=$(cat)
STOP_HOOK_ACTIVE=$(echo "$INPUT" | jq -r '.stop_hook_active')

# 무한루프 방지: 이미 재시도 중이면 멈춤 허용
if [ "$STOP_HOOK_ACTIVE" = "true" ]; then
  exit 0
fi

ROOT=$(git rev-parse --show-toplevel 2>/dev/null || echo ".")
TASKS_FILE="$ROOT/tasks.json"

# tasks.json 존재 확인
if [ ! -f "$TASKS_FILE" ]; then
  exit 0
fi

# pending 태스크 수 확인
PENDING=$(jq '[.tasks[] | select(.status == "pending")] | length' "$TASKS_FILE")

if [ "$PENDING" -gt 0 ]; then
  echo "{\"decision\":\"block\",\"reason\":\"tasks.json에 미완료 태스크 ${PENDING}개. 다음 pending 태스크를 처리하세요.\"}"
else
  exit 0
fi
