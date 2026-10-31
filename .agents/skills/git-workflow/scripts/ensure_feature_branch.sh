#!/usr/bin/env bash
set -euo pipefail

TARGET_BRANCH="${1:-}"

git fetch origin develop >/dev/null 2>&1 || true
if git show-ref --verify --quiet "refs/remotes/origin/develop"; then
  git branch -f develop origin/develop >/dev/null 2>&1 || true
fi

CURRENT_BRANCH="$(git rev-parse --abbrev-ref HEAD)"
IS_MERGED_INTO_DEVELOP=0

if [ "$CURRENT_BRANCH" != "develop" ] && [ "$CURRENT_BRANCH" != "main" ]; then
  if git merge-base --is-ancestor HEAD develop 2>/dev/null; then
    IS_MERGED_INTO_DEVELOP=1
  fi
fi

if [ "$CURRENT_BRANCH" != "develop" ] && [ "$CURRENT_BRANCH" != "main" ] && [ "$IS_MERGED_INTO_DEVELOP" -eq 0 ]; then
  echo "Already on active unmerged feature branch: $CURRENT_BRANCH"
  exit 0
fi

if [ -z "$TARGET_BRANCH" ]; then
  echo "Error: Current branch '$CURRENT_BRANCH' is '$CURRENT_BRANCH' (or already merged into 'develop'). Provide a new branch name (e.g., feat/my-feature)." >&2
  exit 1
fi

if [ "$TARGET_BRANCH" = "$CURRENT_BRANCH" ] && [ "$IS_MERGED_INTO_DEVELOP" -eq 1 ]; then
  echo "Error: Branch '$TARGET_BRANCH' was already merged into 'develop'. Provide a new branch name for the new work." >&2
  exit 1
fi

if git show-ref --verify --quiet "refs/heads/$TARGET_BRANCH"; then
  echo "Switching to existing branch '$TARGET_BRANCH'..."
  git checkout "$TARGET_BRANCH"
else
  echo "Creating new branch '$TARGET_BRANCH' from 'develop'..."
  git checkout -b "$TARGET_BRANCH" develop
fi
