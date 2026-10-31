---
name: git-workflow
description: >-
  Always use this skill whenever starting new feature/fix work without an active
  feature branch, or when the user asks to commit changes, manage git branches,
  or prepare a Pull Request.
---

# Git Branching, Atomic Commits & PR Workflow Skill

Follow this workflow whenever committing code, managing branches, or preparing changes in this repository.

## 1. Branch Verification & Creation (From `develop`)

Before staging or committing changes (or at the start of a new task), verify that you are NOT working directly on `develop` or `main`.

1. **Check the active branch**:
   Run [`ensure_feature_branch.sh`](./scripts/ensure_feature_branch.sh) with the desired branch name (e.g., `feat/short-description`, `fix/issue-name`, `refactor/module-name`, `docs/topic`):
   ```bash
   bash .agents/skills/git-workflow/scripts/ensure_feature_branch.sh <type>/<branch-name>
   ```
   - If you are already on a dedicated feature branch (not `develop` or `main`), keep using the current branch.
   - If you are on `develop` or `main`, the script automatically creates and switches to the new branch based on `develop` (`git checkout -b <branch-name> develop`) while preserving uncommitted working-tree changes.

## 2. Granular, Atomic Commits (Conventional Commits)

Never bundle unrelated changes into a single monolithic commit. When the user requests to commit changes, inspect `git status -uall` and `git diff` and break the work into **small, logical commits** ordered by responsibility:

1. **Recommended Commit Order**:
   - **Dependencies & Configuration** (`chore: ` or `build: `) — e.g., `pyproject.toml`, `requirements.txt`, `config/settings.py`, `vercel.json`.
   - **Models & Migrations** (`feat: ` or `refactor: `) — e.g., `portfolio/models.py`, `portfolio/migrations/`.
   - **Serializers, Views & Routes** (`feat: `, `fix: `, or `refactor: `) — e.g., `portfolio/serializers.py`, `portfolio/views.py`, `portfolio/urls.py`, `portfolio/admin.py`.
   - **Frontend Templates & Static Assets** (`feat: ` or `style: `) — e.g., `templates/`, `static/`.
   - **Automated Tests** (`test: ` or `feat: `) — e.g., `config/tests/`, `portfolio/tests/`.
   - **Documentation & Agent Customizations** (`docs: ` or `chore: `) — e.g., `docs/`, `README.md`, `.agents/`.

2. **Commit Message Rules**:
   - Always write commit messages in **English (EN-US)**.
   - Strictly follow **Conventional Commits**: `feat: `, `fix: `, `refactor: `, `docs: `, `test: `, `chore: `.
   - Explain **what** changed and **why** the change occurred.
   - Stage only the specific files or hunks belonging to each logical unit (`git add <paths>` or `git add -p`) before running `git commit -m "<type>: <descriptive message>"`.

## 3. Automatic Pull Request Title & Description Generation

Immediately after completing the atomic commits, ALWAYS generate and display a ready-to-copy **Pull Request Title** and **PR Description** in your response to the user using the template below:

```markdown
### PR Title
`<type>: <concise summary of the branch changes>`

### PR Description
## Summary
<1-2 sentences explaining why these changes were made and the problem they solve.>

## Commits & Key Changes
- **<Component / Layer>:** <Brief description of the changes>
- **<Component / Layer>:** <Brief description of the changes>

## Verification & Testing
- [x] `pytest -s -v`
- [x] `flake8 .`
- [x] `pylint config portfolio manage.py`
```
