# Copilot Instructions

## Build, test, and lint commands

This repository is organized as independent Go lesson modules (`01hello` ... `06Pointers`), each with its own `go.mod`. Run commands from a specific lesson directory unless otherwise noted.

- **Run an example:** `go run .` (or `go run main.go`)
- **Build a lesson:** `go build .`
- **Run tests for a lesson:** `go test ./...`
- **Run a single test in a lesson:** `go test -run TestName ./...`
- **Run tests across all lessons from repo root:**
  `for d in */; do [ -f "$d/go.mod" ] && (cd "$d" && go test ./...); done`

There is no repo-level `go.mod`, so repo-wide Go commands must be run per module (or via a loop).

## High-level architecture

- The repo is a set of **topic-based, standalone executables**. Each numbered directory represents one concept (hello world, variables, input, conversions, time, pointers).
- Every lesson is a single `package main` program in `main.go`, intended to be run directly and inspected in isolation.
- There are **no shared internal packages** and no cross-directory imports between lessons; examples are intentionally self-contained.

## Key conventions in this codebase

- Keep lesson scope local: changes for one concept should stay inside that lesson folder and not introduce shared abstractions unless explicitly requested.
- Prefer standard-library-first examples (`fmt`, `bufio`, `os`, `strconv`, `strings`, `time`) and straightforward CLI output/input flows.
- Preserve the teaching-oriented style: simple `main()`-driven demos with inline comments explaining concepts near the code being demonstrated.
