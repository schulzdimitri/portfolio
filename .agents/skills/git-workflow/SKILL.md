---
name: git-workflow
description: >-
  CRITICAL: Always use this skill FIRST before writing or modifying any code for
  a new feature, bugfix, or refactoring task, and whenever the user asks to
  commit changes, manage git branches, or prepare a Pull Request.
---

# Git Branching, Atomic Commits & PR Workflow Skill

Follow this workflow **before starting any code implementation** and whenever committing code or preparing a Pull Request.

## 1. Pre-Implementation Branch Creation (From `develop`)

**MANDATORY STEP BEFORE EDITING ANY PROJECT CODE:**
Whenever the user asks to implement a new feature, fix a bug, or refactor code, your **very first action** (before using `write_file` or `edit_file` on project code) MUST be to ensure you are on a dedicated, unmerged feature branch created from the latest `develop` branch:

1. **Run [`ensure_feature_branch.sh`](./scripts/ensure_feature_branch.sh)** with a descriptive branch name (`feat/<short-name>`, `fix/<short-name>`, `refactor/<short-name>`, `docs/<short-name>`):
   ```bash
   bash .agents/skills/git-workflow/scripts/ensure_feature_branch.sh <type>/<branch-name>
   ```
   - The script fetches `origin/develop`, fast-forwards local `develop`, and checks if the current branch is `develop`, `main`, or **already merged into `develop`** (e.g., from a previous merged PR).
   - If you are on `develop`, `main`, or a previously merged branch, it automatically creates and switches to `<type>/<branch-name>` starting from `develop` (`git checkout -b <branch-name> develop`).

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
