# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this repository is

This is **course material**, not an application. It is the "Linux Fundamentals" training delivered by cloudnativetrainings. There is nothing to build, deploy, or unit-test — the deliverable is the set of lab instructions that students follow.

Each top-level `NN_topic/` directory (`01_hello-linux` … `11_zsh`) is one self-contained lab. The lab's `README.md` is the student-facing script: prose explaining a concept, interleaved with fenced ```bash``` code blocks of commands the student types. A few labs ship companion source files (e.g. `03_executables/my-executable.go`) that the lab text references.

## How the environment works

- Students open the repo as a **GitHub Codespace / devcontainer** (`.devcontainer/devcontainer.json`), which bind-mounts the repo at **`/training/`** inside the container.
- From there they SSH into a separate training **VM** and run the lab commands. The container itself is mostly a launchpad + editor.
- Because the mount point is fixed, **all file paths in lab material must be absolute and start with `/training`** (e.g. `/training/.secrets/ssh-config`). Relative paths that depend on the student's CWD are a bug.
- The devcontainer image is pinned in `devcontainer.json` and disables VS Code AI/telemetry features on purpose — keep those settings.

## Editing conventions

- Code blocks are the primary content. They frequently include commands that are *expected to fail* as a teaching device (e.g. running a script before `chmod +x`); the surrounding prose explains why. Do not "fix" these.
- **Placeholders are intentional**: `TODO`, `XXXXX`, `TODO-STUDENT-EMAIL@...` etc. Never substitute real values for them. If you find a genuine stray `TODO`, surface it rather than silently changing it.
- Lab prose may contain casual phrasing/typos; the linter skills deliberately ignore prose grammar and focus on the **code inside the blocks** and on path correctness.

## Off-limits directories

Never read, open, grep, or write under:
- `.secrets/` — SSH private keys, configs, service-account material (also git-ignored).
- `.99_todos/` — internal authoring notes, excluded from the student view.

## Tooling (Claude Code skills)

This repo defines custom skills in `.claude/skills/` instead of build/test commands. Invoke them by the trigger phrases below:

- **`code-linter`** ("lint" / "lint code") — checks the bash/yaml code *inside* lab READMEs and code files for errors, and verifies `/training`-absolute paths. Respects the scoping rules: a bare lab number (`07`) lints only that lab; otherwise unstaged changes only; otherwise every `*/README.md`.
- **`md-linter`** ("lint md" / "lint markdown") — Markdown/typo checks on a lab README.
- **`secrets-remover`** ("remove secrets") — scans for sensitive info that must not be pushed to GitHub.

When asked to "lint" without qualification, both `code-linter` and `md-linter` match — pick based on whether the concern is code-in-blocks or Markdown prose, or ask.
