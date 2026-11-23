---
name: go-system-design-lesson
description: Author or extend Ukrainian-language lessons for this Go and System Design course. Use when the user asks to write, add, or rework a lesson, topic, module, example, or exercise in 01-go-review, 02-system-design, or 03-practice-projects.
---

# Go & System Design Lesson Author

Write lessons for this course repo (`C:\Go System Design`) in the user's established style.

## Language & audience
- All prose in Ukrainian.
- Audience already knows JS/TS, is relearning Go, forgets syntax easily.
- Never invent "when Go vs when Node.js" trade-off sections. Node/JS appears ONLY as a syntax analogy when a Go construct might be unclear (e.g. `make` ~ `new Array`, `<-chan` ~ nothing in JS, goroutine ~ lightweight async task).

## Lesson structure (Markdown)
1. Short title + one-line goal.
2. Concept explained simply, with a real-life analogy.
3. Go code block with inline comments explaining each non-obvious line.
4. Syntax analogies to JS/TS only where helpful (fenced `js` block beside the Go).
5. Common mistakes / gotchas (if any).
6. 3-5 practice exercises at the end.

## Style rules
- Comments in Go code: Ukrainian, explain WHY not WHAT, only on non-obvious lines.
- Keep examples runnable and self-contained (`package main`, full imports).
- Prefer tables for comparisons (concept | Go | JS-analogy).
- No emoji unless the user asks.
- One lesson per file, numbered like `NN-topic.md` inside the right module folder.
- Do not create files until the user names the topic; build one lesson per session to save tokens.

## Where lessons live
- `01-go-review/` — Go language review.
- `02-system-design/` — system design concepts.
- `03-practice-projects/` — project walkthroughs.
- After writing, remind the user to `git add`/commit only if they ask; never auto-commit.