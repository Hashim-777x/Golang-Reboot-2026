# Golang Reboot 2026 

> Learning Go the right way — from syntax to systems, no shortcuts.  
> **Status:** Active | Hyderabad, India  
> **Started:** January 2026 | **Restarted seriously:** May 2026

---

## Why this repo exists

I'm a self-taught backend engineer from Hyderabad.  
No CS degree yet. No MacBook. Built my current rig from scrap parts — i3 4th Gen, wall-mounted in cardboard with a custom cooling setup.

This repo is my Go learning journal. Every file I push here is code I wrote myself — not copy-pasted, not AI-generated. If I don't understand a line, it doesn't go in.

The goal isn't to finish a course. It's to understand Go deeply enough to build systems from scratch — TCP servers, databases, parsers — without reaching for a framework.

---

## What this repo covers

**Phase 1 — Go Basics** (Jan 2026)  
Syntax fundamentals: variables, types, loops, functions, pointers, arrays, slices, maps, closures, structs, interfaces, goroutines, channels.  
Done via [A Tour of Go](https://go.dev/tour) and Go language spec.  
Files: `phase-1/`

> Honest note: Jan–Mar 2026 commits were inconsistent — internship, exams, and a rough patch killed the streak. Some early files were rushed. Restarting properly from May 2026.

**Phase 2 — Go Basics, done right** (May 2026 → )  
Revisiting the concepts I rushed. Writing every line from scratch, questioning every behaviour, taking notebook notes.  
Current focus: pointers, slices internals, interfaces.  
Files: `phase-2/`

---

## Daily Log

| Day | Date | File | What I actually did |
|-----|------|------|---------------------|
| 01 | Jan 01 | — | Research + roadmap planning |
| 02 | Jan 02 | README | Folder structure + roadmap in README |
| 03 | Jan 03 | first.go | Variables, types, functions, returns |
| 04 | Jan 04 | second.go | For loops (while-style), if with short statements, switch |
| 05 | Jan 05 | third.go | Pointers, fixed-size arrays |
| 06 | Jan 06 | forth.go | Slices, len vs cap, make function |
| 07 | Jan 08 | fifth.go | Maps, closures |
| 08 | Jan 09 | sixth.go | Struct receivers, Stringer interface |
| 09 | Jan 10 | seventh.go | Implicit interface implementation, custom error types |
| 10 | Jan 11 | eighth.go | Goroutines and channels |
| 11 | Jan 12 | effective_go.go | Allocation, literals, defer |
| 12 | Jan 13 | language_spec.go | Slice internals, memory layout, 24-byte SliceHeader |
| 13 | Jan 14 | language_spec.go | Updated slice spec notes |
| 14 | Jan 15 | tcp_client.go | First TCP connection in Go |
| 15 | Jan 17 | tcp_client.go | Refactored with timeouts and buffered writes |
| 16 | Jan 18 | tcp_client.go | Separated I/O functions, error context |
| — | Jan 20–Feb 14 | — | Semester exams + lab internals/externals. No commits. |
| 17 | Feb 15 | tcp_client.go | Rewrote: concurrent peer handling |
| 18 | Feb 16 | tcp_client.go | Efficient I/O reading in peer readLoop |
| 19 | Feb 17 | tcp_client.go | Thread-safe message passing via channels |
| — | Feb 18–May 12 | — | Internship (multiple projects), semester exams, final year project, depression, no golang learning. Honest gap.|
| 20 | May 16 | — | Restarted. Cleaned README. New plan. |

---

## What's next

The 6 systems projects are in a separate repo: [go-systems-from-scratch](https://github.com/Hashim-777x/go-systems-from-scratch)

Those projects are where I apply everything from this repo at a lower level — TCP, HTTP parsing, databases, concurrency — built from scratch without frameworks.

---

## Connect

- **X (Twitter):** [@KillerPand34973](https://x.com/KillerPand34973)  
- **LinkedIn:** [syed-hashim721](https://linkedin.com/in/syed-hashim721)
