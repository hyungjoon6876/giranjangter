#!/bin/bash
# Ralph Dual-Loop preflight check
# Usage: bash scripts/ralph-preflight.sh
# Exit 0 = OK to start. Non-zero = fix the issue first.

set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$REPO_ROOT"

fail() { echo "FAIL: $1" >&2; exit 1; }
ok()   { echo "OK:   $1"; }

# 1. Git: clean working tree + on main
if [ -n "$(git status --porcelain)" ]; then
  fail "working tree not clean. commit/stash first."
fi
ok "git clean"

CURRENT_BRANCH="$(git rev-parse --abbrev-ref HEAD)"
if [ "$CURRENT_BRANCH" != "main" ]; then
  fail "not on main (current: $CURRENT_BRANCH). switch with: git checkout main"
fi
ok "on main"

# 2. State files exist
for f in \
  tasks/bug-state.json \
  tasks/improve-state.json \
  tasks/bug-found.md \
  tasks/bug-fixed.md \
  tasks/needs-human.md \
  tasks/improvements.md \
  tasks/ralph-progress.md \
  tasks/ralph-dual.md; do
  if [ ! -f "$f" ]; then
    fail "missing state file: $f"
  fi
done
ok "state files present"

# 3. State files are valid JSON
for j in tasks/bug-state.json tasks/improve-state.json; do
  if ! python3 -c "import json; json.load(open('$j'))" 2>/dev/null; then
    fail "invalid JSON: $j"
  fi
done
ok "state JSON valid"

# 4. Web dev server reachable on :3000
if ! curl -sf -m 3 http://localhost:3000 >/dev/null 2>&1; then
  fail "web dev server not running on :3000. start in another terminal: cd web && npm run dev"
fi
ok "web dev server up"

# 5. Production reachable
if ! curl -sf -m 5 https://giranjt.com/ >/dev/null 2>&1; then
  fail "https://giranjt.com/ not reachable"
fi
ok "production reachable"

# 6. Playwright present in web/
if ! (cd web && [ -f node_modules/.bin/playwright ]); then
  fail "playwright not installed in web/. run: cd web && npm install"
fi
ok "playwright installed"

echo "=== Preflight passed ==="
