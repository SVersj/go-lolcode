# 🤖 LOLCODE Interpreter (Go)  
*Lab 4 — Командна робота | Методології та технології розробки ПЗ*

![Go CI](https://github.com/SVersj/go-lolcode/actions/workflows/go.yml/badge.svg)

---

## 🔍 Опис проєкту

Цей репозиторій містить простий інтерпретатор мови програмування [LOLCODE](https://esolangs.org/wiki/LOLCODE), реалізований на Go.  
Метою проєкту є:

- відпрацювання навичок командної розробки через GitHub
- реалізація базового парсингу та виконання LOLCODE-інструкцій
- CI-тестування
- Pull Request + Code Review

---

## 🧠 Функціонал

- 📥 Парсинг LOLCODE-програм (`HAI`, `VISIBLE`, `BTW`, `KTHXBYE`)
- 🧾 Виведення значень на екран
- ✅ Підтримка простих виразів
- 🚫 Валідація синтаксису
- 🧪 Тестування парсера
- 🧰 CI через GitHub Actions

---

## 📁 Приклад LOLCODE-програми

```
HAI 1.2
VISIBLE "Hello, human!"
KTHXBYE
```

---

## 🚀 Запуск

### 🔧 Вимоги
- Go >= 1.20
- Git

### ⚙️ Команди

```bash
git clone https://github.com/SVersj/go-lolcode.git
cd go-lolcode

# запуск основної програми
go run cmd/main.go

# запуск тестів
go test ./...
```

---

## 📦 Структура репозиторію

```
go-lolcode/
├── cmd/main.go              # точка входу
├── internal/
│   ├── lexer/lexer.go
│   ├── parser/parser.go
│   └── interpreter/interpreter.go
├── tests/
│   ├── lexer_test.go
│   ├── parser_test.go
│   └── interpreter_test.go
├── go.mod / go.sum
└── .github/workflows/go.yml
```

---

## 🧪 CI (GitHub Actions)

Файл: `.github/workflows/go.yml`

- запускає `go test ./...` при push/PR до `main`
- гарантує стабільність та правильність коду

---

## 🤝 Командна робота

### Учасники:
- **SVersj** — імплементація парсера, інтерпретатора, CI, README
- **anirina** — code review, merge PR, issue-доповіді

### Pull Requests:

| № | Назва                      | Статус | Review by |
|---|----------------------------|--------|-----------|
| 1 | feature/project-structure  | ✅ merged | ghost-coder-co |
| 2 | feature/lexer              | ✅ merged | ghost-coder-co |
| 3 | feature/parser             | ✅ merged | ghost-coder-co |
| 4 | feature/interpreter        | ✅ merged | ghost-coder-co |
| 5 | feature/ci                 | ✅ merged | ghost-coder-co |

---

## 📎 Посилання

- 📄 [Design-документ](./Design.md)
- 🧠 [Що таке LOLCODE?](https://esolangs.org/wiki/LOLCODE)
- 📘 [Go Tour](https://go.dev/tour)

---

## 📌 Примітка

Робота виконана у рамках лабораторної №4 з дисципліни  
**«Методології та технології розробки ПЗ»**
