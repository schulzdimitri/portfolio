#!/usr/bin/env bash
set -euo pipefail

TARGET_BRANCH="${1:-}"
CURRENT_BRANCH="$(git rev-parse --abbrev-ref HEAD)"

if [ "$CURRENT_BRANCH" != "develop" ] && [ "$CURRENT_BRANCH" != "main" ]; then
  echo "Already on feature branch: $CURRENT_BRANCH"
  exit 0
fi

if [ -z "$TARGET_BRANCH" ]; then
  echo "Error: Currently on '$CURRENT_BRANCH'. Please provide a new branch name (e.g., feat/my-feature)." >&2
  exit 1
fi

if git show-ref --verify --quiet "refs/heads/$TARGET_BRANCH"; then
  echo "Switching to existing branch '$TARGET_BRANCH'..."
  git checkout "$TARGET_BRANCH"
else
  echo "Creating new branch '$TARGET_BRANCH' from 'develop'..."
  git checkout -b "$TARGET_BRANCH" develop
fi
