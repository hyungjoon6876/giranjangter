#!/bin/bash
# harness-verify: Tier 2+3 Pre-Commit 검증
# 자동 생성됨 — /harness verify 재실행으로 재생성 가능

set -uo pipefail

# jq 필수
command -v jq >/dev/null 2>&1 || { echo "jq not found, skipping verification"; exit 0; }

# stdin에서 JSON 읽기
INPUT=$(cat)
COMMAND=$(echo "$INPUT" | jq -r '.tool_input.command // ""')

# git commit이 아니면 통과
case "$COMMAND" in
  git\ commit*|git\ -C\ *commit*) ;;
  *) exit 0 ;;
esac

ROOT=$(git rev-parse --show-toplevel)
EXIT_CODE=0

# === Tier 2: 타입체크/컴파일 ===

# Go vet
echo "[Tier 2] go vet..."
(cd "$ROOT/backend" && go vet ./... 2>&1 | tail -20) || EXIT_CODE=$?

if [ $EXIT_CODE -ne 0 ]; then
  echo '{"hookSpecificOutput":{"hookEventName":"PreToolUse","permissionDecision":"deny","permissionDecisionReason":"Tier 2 실패: go vet"}}'
  exit 0
fi

# TypeScript
echo "[Tier 2] tsc..."
(cd "$ROOT/web" && npx tsc --noEmit 2>&1 | tail -20) || EXIT_CODE=$?

if [ $EXIT_CODE -ne 0 ]; then
  echo '{"hookSpecificOutput":{"hookEventName":"PreToolUse","permissionDecision":"deny","permissionDecisionReason":"Tier 2 실패: tsc"}}'
  exit 0
fi

# === Tier 3: 테스트 ===

# Go test
echo "[Tier 3] go test..."
(cd "$ROOT/backend" && go test ./... 2>&1 | tail -20) || EXIT_CODE=$?

if [ $EXIT_CODE -ne 0 ]; then
  echo '{"hookSpecificOutput":{"hookEventName":"PreToolUse","permissionDecision":"deny","permissionDecisionReason":"Tier 3 실패: go test"}}'
  exit 0
fi

# Vitest
echo "[Tier 3] vitest..."
(cd "$ROOT/web" && npx vitest run 2>&1 | tail -20) || EXIT_CODE=$?

if [ $EXIT_CODE -ne 0 ]; then
  echo '{"hookSpecificOutput":{"hookEventName":"PreToolUse","permissionDecision":"deny","permissionDecisionReason":"Tier 3 실패: vitest"}}'
  exit 0
fi

exit 0
