# 📄 Дизайн-документ  
**Проєкт:** LOLCODE Interpreter  
**Мова реалізації:** Go  
**Лабораторна №4 — Командна робота**

---

## 🎯 Мета

Розробити інтерпретатор мови програмування LOLCODE, який:
- виконує лексичний аналіз (tokenization)
- розпізнає базові синтаксичні конструкції
- інтерпретує інструкції (`VISIBLE`, `HAI`, `KTHXBYE`)
- підтримує тестування та CI
- реалізується у форматі командної розробки з Pull Request'ами та code-review

---

## 🧠 Обґрунтування

- **Мова:** LOLCODE — проста, зрозуміла, з гумористичним синтаксисом, ідеальна для демонстрації основ парсингу та інтерпретації.
- **Go:** статично типізована мова з чудовою підтримкою CLI, тестування й модульності.
- **Командна модель:** дозволяє застосувати практики CI/CD, review, гілкування.

---

## 📚 Основні поняття LOLCODE

| LOLCODE | Опис                   |
|---------|------------------------|
| `HAI`   | Початок програми       |
| `VISIBLE` | Вивід значення         |
| `KTHXBYE` | Завершення програми    |

Приклад:
```
HAI 1.2
VISIBLE "Hello, world!"
KTHXBYE
```

---

## 🏗 Архітектура

```
go-lolcode/
├── cmd/
│   └── main.go              # точка входу
├── internal/
│   ├── lexer/lexer.go       # лексичний аналіз
│   ├── parser/parser.go     # синтаксичний аналіз
│   └── interpreter/interpreter.go  # виконання
├── tests/
│   ├── lexer_test.go
│   ├── parser_test.go
│   └── interpreter_test.go
├── go.mod / go.sum
└── .github/workflows/go.yml  # CI
```

---

## 👥 Командна модель

**Учасники:**
- `SVersj` — основний розробник, CI, документація
- `ghost-coder-co` — код-ревʼю, issue-трекинг, тестування, merge

**Етапи:**
- `feature/project-structure` — базова структура
- `feature/lexer` — токенізація
- `feature/parser` — парсинг `VISIBLE`
- `feature/interpreter` — виконання `VISIBLE`
- `feature/ci` — інтеграція GitHub Actions

---

## 🔁 Взаємодія

- `anirina` відкриває Pull Request
- `SVersj` виконує ревʼю та merge
- Review містять коментарі з прикладами коду
- Створено issue з виявленою проблемою, яке закривається PR

---

## 🧪 Тестування

- `lexer_test.go` перевіряє токени: `HAI`, `VISIBLE`, `KTHXBYE`, `UNKNOWN`
- `parser_test.go` перевіряє структуру `Statement{Command, Args}`
- `interpreter_test.go` перехоплює stdout і перевіряє фактичний вивід

---

## ⚙️ CI (GitHub Actions)

```yaml
name: Go CI

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    - uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    - run: go mod tidy
    - run: go test ./... -v
```

- Автоматично перевіряє тести при кожному PR або push до `main`
- CI-бейдж додано в README

---

## 📈 Перспективи розвитку

- Підтримка математичних операцій (`SUM OF`, `DIFF OF`)
- Введення користувача (`GIMMEH`)
- Підтримка змінних
- Побудова AST
- Web-інтерфейс

---

## 📌 Висновок

Проєкт виконано відповідно до вимог лабораторної роботи.  
Застосовано практики спільної розробки: гілкування, code-review, CI.  
LOLCODE-інтерпретатор реалізує повний цикл: лексер → парсер → інтерпретатор → CI → тести.  
Робота готова до захисту 💼
