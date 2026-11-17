<h1> Contributing </h1>


# AI policy

No AI-generated code. IDE assistants (Copilot etc.) stay off.

The mentor bot may explain concepts, point to docs, review code and screenshots, and ask questions. It never writes solutions or algorithms (at most ~5 lines of generic syntax).

The PR-agent reviews and describes problems, but makes no code suggestions.

In reviews, the apprentice explains his code in his own words. Code he cannot explain is not merged.

# Git & GitHub

The apprentice owns the setup: public repo, branch protection, CI, Dependabot, nightly pre-releases via GitHub Actions.

One feature branch per ticket (AUS-<n>-short-name). main stays clean, with no direct pushes.

Conventional Commits with a rationale ("why") in the body. No "init", "wip" or "fix typo" commits on main.

Squash merge: PR title = conventional commit, PR body = rationale.

Daily push to the feature branch. A watchdog (IN-423) notifies the apprentice and mentor if no push happened Tue–Fri by 16:00.

# Cadence (Scrum lite)

2-week sprints, Tuesday → Monday (Mondays = vocational school)

Planning: Tuesday of week 1. One-sentence sprint goal; tickets move from Backlog to To Do here, not before.

Daily: 09:00 sharp: yesterday / today / blockers.

Review + retro: Friday of week 2. The apprentice demos the running app, not just code.

Rule against silent stalls: blocked for more than 2 hours → tell the mentor. Work-in-progress limit: one ticket.

# Planning in Jira

Releases = milestones with target dates (timeline). Epics have start and end dates. Dates are a plan, not a promise: they may move, but moves are made visible in the review.

Stories are at most one day of work. They describe what and why, never how.
