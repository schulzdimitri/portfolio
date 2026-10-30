#!/usr/bin/env bash
set -euo pipefail

VENV_BIN="./.venv/bin"

if [ ! -d "$VENV_BIN" ]; then
  echo "Virtual environment not found at $VENV_BIN. Using system binaries..."
  VENV_BIN=""
else
  VENV_BIN="$VENV_BIN/"
fi

echo "==> Running pytest..."
${VENV_BIN}pytest -s -v

echo "==> Running flake8..."
${VENV_BIN}flake8 .

echo "==> Running pylint..."
${VENV_BIN}pylint config portfolio manage.py

echo "==> All tests and quality checks passed!"
