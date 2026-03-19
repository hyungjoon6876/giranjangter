#!/bin/bash
# harness-verify: Tier 1 린트/포맷 검증
# 자동 생성됨 — /harness verify 재실행으로 재생성 가능

set -euo pipefail

# stdin에서 JSON 읽기 — PostToolUse 훅은 tool_input을 stdin으로 받는다
INPUT=$(cat)
FILE=$(echo "$INPUT" | jq -r '.tool_input.file_path // ""' 2>/dev/null)

# 파일이 없으면 건너뜀
[ -z "$FILE" ] && exit 0

ROOT=$(git rev-parse --show-toplevel)

# === 확장자 필터 ===
case "$FILE" in
  *.ts|*.tsx|*.js|*.jsx|*.mjs) STACK="web" ;;
  *.go)                        STACK="go" ;;
  *.dart)                      STACK="dart" ;;
  *)                           exit 0 ;;
esac

EXIT_CODE=0

# === Tier 1: 린트/포맷 ===
case "$STACK" in
  web)
    # ESLint — web/ 파일만
    case "$FILE" in
      */web/*)
        cd "$ROOT/web"
        npx eslint --no-warn-ignored "$FILE" 2>&1 | tail -20 || EXIT_CODE=$?
        ;;
    esac
    ;;
  go)
    cd "$ROOT/backend"
    golangci-lint run "$FILE" 2>&1 | tail -20 || EXIT_CODE=$?
    ;;
  dart)
    cd "$ROOT/frontend"
    flutter analyze "$FILE" 2>&1 | tail -20 || EXIT_CODE=$?
    ;;
esac

exit $EXIT_CODE
