# Log Linter

Статический анализатор для Go, который проверяет лог-сообщения на соответствие базовым правилам качества.

Линтер анализирует вызовы логгеров и валидирует текст сообщений.

Поддерживается запуск как отдельной утилиты и интеграция с `golangci-lint`.

---

# Возможности

Проверяет лог-сообщения по следующим правилам:

- сообщение должно начинаться со строчной буквы  
- допускается только английский язык (кириллица и другие языки запрещены)  
- запрещены спецсимволы и эмодзи (`!`, `?`, `🚀`, `…`)  
- запрещено логировать чувствительные данные (пароли, токены, ключи)

---

# Поддерживаемые логгеры

Линтер ищет вызовы следующих пакетов:

- `log` — стандартная библиотека Go  
- `log/slog`  
- `go.uber.org/zap`

---

# Совместимость

Линтер может использоваться:

- как отдельная CLI-утилита
- как анализатор через `analysis.Analyzer`
- внутри `golangci-lint`

Поддерживается конфигурация через JSON.

---

# Установка и запуск

## 1. Клонировать репозиторий
```bash
git clone github.com/PhosFactum/loglinter
cd test-selectel
```

## 2. Скачать зависимости
```bash
go mod tidy
```

## 3. Собрать бинарник 
```bash
make build
```

После сборки появится бинарник
```
./bin/loglinter
```

## 4. Проверить запуск
```bash
./bin/loglinter -h
```
Команда должна вывести справку и доступные флаги.

# Тестирование

Запуск всех тестов:
```bash
make test
```

Запуск только интеграционных тестов:
```bash
go test ./internal/analyzer -v -run TestAnalyzer
```

Ручная проверка на тестовых данных:
```bash
make run-err
make run-ok
```
Так должны появится ошибки в первом случае и без них во втором.

# Использование

Проверить весь проект:
```bash
./bin/loglinter ./...
```

Проверить конкретный пакет:
```bash
./bin/loglinter ./internal/...
```

Запуск с кастомной конфигурацией:
```bash
./bin/loglinter -config .loglinter.json ./...
```

Вывод в JSON (удобно для CI):
```bash
./bin/loglinter -json ./... > report.json
```

---

# Конфигурация

Пример конфигурационного файла .loglinter.json:

```json
{
  "rules": {
    "lowercase": { "enabled": true },
    "language": { "enabled": true },
    "symbols": { "enabled": true },
    "sensitive": { "enabled": true }
  },
  "sensitive": {
    "keywords": ["password", "token", "secret", "api_key", "custom_key"]
  },
  "symbols": {
    "forbidden": ["!", "?", "…"]
  }
}
```

Запуск линтера с конфигурацией:
```bash
./bin/loglinter -config .loglinter.json ./...
```

---

# Очистка сборки

Удалить артефакты сборки:
```bash
make clean
```

---

# Отладка

Если возникают проблемы со сборкой:
```bash
go mod tidy
go build -v ./...
```

Если падают тесты:
```bash
go test ./... -v
```

Проверить версию Go:
```bash
go version
```

Для проекта требуется Go 1.22 или новее.

---

# Быстрый старт
```bash
git clone github.com/PhosFactum/loglinter && cd loglinter
go mod tidy
make build
./bin/loglinter ./...
```
