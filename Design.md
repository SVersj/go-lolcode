# Design Document

What is a design document? How should we write it?  
https://www.industrialempathy.com/posts/design-docs-at-google/

---

## Context

This project is implemented as part of Lab #4: Team Software Development.  
It aims to develop a minimal interpreter for the LOLCODE programming language using team practices such as GitHub Flow (branching, pull requests, code reviews), modular design, unit testing, and CI.

**Team members:**
- **SVersj** — lead developer, CI configuration, documentation
- **ghost-coder-co** — code reviewer, issue tracking, test authoring

---

## Goals and non-goals

### Goals

- Implement a lexical analyzer (lexer) for LOLCODE
- Implement a parser for a minimal subset of the language: `HAI`, `VISIBLE`, `KTHXBYE`
- Create an interpreter that executes these statements
- Add unit tests for each subsystem (lexer, parser, interpreter)
- Integrate GitHub Actions CI pipeline
- Apply collaborative workflow: pull requests, code reviews, issues

**Specification:** Internal course guideline for Lab #4 — Team Development Project

### Non-goals

- Full support for the LOLCODE language (e.g., variables, math expressions, control flow)
- Input handling (`GIMMEH`) or arithmetic (`SUM OF`, `DIFF OF`)
- Web or GUI interface
- Persistent data storage or compiled output
- Authentication, authorization, or user session management

---

## Subsystems

### Project structure

```plaintext
go-lolcode/
├── cmd/main.go                   // CLI entrypoint
├── internal/
│   ├── lexer/lexer.go            // Lexical analysis (tokenization)
│   ├── parser/parser.go          // Parsing into statement structures
│   └── interpreter/interpreter.go // Execution of parsed instructions
├── tests/
│   ├── lexer_test.go
│   ├── parser_test.go
│   └── interpreter_test.go
├── go.mod / go.sum               // Go modules
└── .github/workflows/go.yml      // CI pipeline
```

### Module interaction

```
main.go → lexer → parser → interpreter
```

### Dependency injection

Manual dependency wiring is used; no external DI frameworks are applied.

---

## Authorization subsystem

Not applicable.

This project is a local CLI utility and does not require authentication, user management, or session tokens.

---

## Data storage subsystem

Not applicable.

The interpreter does not persist any data.  
Programs are read from standard input or file at runtime and executed in memory.  
There is no need for on-disk formats, databases, or session state.

---

## Business logic

The core logic is responsible for:

- Splitting input code into tokens (`lexer`)
- Parsing these tokens into statements (`parser`)
- Executing statements sequentially (`interpreter`)

### Example program

```lolcode
HAI 1.2
VISIBLE "Hello, world!"
KTHXBYE
```

### Execution flow

1. Program starts at `HAI`
2. `VISIBLE` prints the string to stdout
3. `KTHXBYE` ends the program

The interpreter currently supports only sequential instruction execution.

---

## HTTP frontend

Not applicable.  
There is no HTTP layer; this is a CLI-based interpreter.

---

## Alternatives considered

| Option                          | Reason for Rejection                                      |
|---------------------------------|------------------------------------------------------------|
| Using Python or Java            | Slower tooling and less native CLI/CI integration         |
| Implementing full LOLCODE spec | Too broad for an educational lab project                  |
| Adding variables and expressions| Would significantly expand scope and complexity           |
| Using parser/lexer libraries    | Educational goal was to build these components manually   |

The final design was chosen to match the scope and learning goals of a limited-time team project focused on fundamentals of parsing and modular development.
