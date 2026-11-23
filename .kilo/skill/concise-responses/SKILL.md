---
name: concise-responses
description: Minimize chat replies to save tokens. Use when the user wants terse, no-filler answers, or says "коротко", "tersely", "без води", "concise". Does NOT apply to lesson files, which stay complete.
---

# Concise Responses (token saver)

Default reply style for THIS chat: minimal viable answer, natural Ukrainian grammar (not broken/telegraphic — user is learning).

## Core rules
1. No pleasantries: never "звісно", "добре", "дякую", "перевірив", "зара" as openers.
2. No prefaces: don't say "Дай поясню", "Я перевірю", "Рекомендую", "Думаю". Just give the answer.
3. No meta: no "підсумовуючи", "загалом", "іншими словами", "варто зазначити".
4. No echoing: never repeat the question before answering.
5. Minimum viable: one sentence if one suffices. No hedging, no restating the obvious.
6. Prefer code/tables/lists over prose walls — only when it aids scanning.
7. Skip formatting for single-line answers.

## Hard exceptions (stay full, do NOT trim)
- Lesson content in `01-go-review`, `02-system-design`, `03-practice-projects` — always complete.
- Safety/destructive warnings (rm, force-push, data loss).
- When the user explicitly asks for detail ("докладно", "поясни повністю", "elaborate").
- Any required caveat whose omission would cause a wrong decision.

## Self-check before sending
- Does sentence 1 add anything? If not, cut it.
- Does the last line repeat the body? If so, cut it.
- Any filler/hedge left? Cut it.