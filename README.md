# try_parse_youtrack

CLI для запроса задач из YouTrack через REST API. Сейчас утилита получает
список issues по поисковому запросу и печатает результат в JSON.

## Требования

- Go 1.24 или новее.
- Доступ к YouTrack.
- Permanent token YouTrack с правами на чтение задач.

## Быстрый старт

Запуск без установки:

```bash
YOUTRACK_BASE_URL="https://youtrack.example.com" \
YOUTRACK_TOKEN="perm:..." \
go run . -query "#Unresolved" -top 10
```

Сборка бинарника:

```bash
go build -o try_parse_youtrack .
```

После сборки в корне проекта появится исполняемый файл:

```text
./try_parse_youtrack
```

Запуск собранного бинарника:

```bash
YOUTRACK_BASE_URL="https://youtrack.example.com" \
YOUTRACK_TOKEN="perm:..." \
./try_parse_youtrack -query "project: ABC #Unresolved" -top 20
```

## Настройка

CLI читает настройки в таком порядке приоритета:

1. Флаги командной строки.
2. Переменные окружения.
3. JSON-конфиг.

Если одно и то же значение задано в нескольких местах, победит источник с
более высоким приоритетом.

### Переменные окружения

```bash
export YOUTRACK_BASE_URL="https://youtrack.example.com"
export YOUTRACK_TOKEN="perm:..."
```

### Конфиг

По умолчанию конфиг читается из пользовательской директории настроек:

```text
~/.config/try_parse_youtrack/config.json
```

На других ОС путь определяется через `os.UserConfigDir()`.

Пример конфига:

```json
{
  "base_url": "https://youtrack.example.com",
  "token": "perm:..."
}
```

Можно указать другой файл:

```bash
go run . -config ./config.json
```

## Флаги

```text
-base-url string
    YouTrack base URL, например https://youtrack.example.com

-token string
    Permanent token YouTrack

-config string
    Путь к JSON-конфигу

-query string
    Поисковый запрос YouTrack (по умолчанию "#Unresolved")

-top string
    Максимальное количество задач для загрузки (по умолчанию "10")
```

Справка по флагам:

```bash
go run . -h
```

Или через собранный бинарник:

```bash
./try_parse_youtrack -h
```

## Примеры

Получить 10 нерешенных задач:

```bash
go run .
```

Получить 50 задач по проекту:

```bash
go run . -query "project: ABC" -top 50
```

Получить нерешенные задачи, назначенные на конкретного пользователя:

```bash
go run . -query "assignee: me #Unresolved" -top 25
```

Использовать настройки только через флаги:

```bash
./try_parse_youtrack \
  -base-url "https://youtrack.example.com" \
  -token "perm:..." \
  -query "project: ABC #Unresolved" \
  -top 20
```

## Вывод

Утилита печатает JSON-массив задач в `stdout`. Для каждой задачи запрашиваются
поля:

- `id`
- `idReadable`
- `summary`
- `updated`

Пример:

```json
[
  {
    "id": "2-123",
    "idReadable": "ABC-123",
    "summary": "Fix login error",
    "updated": 1730000000000
  }
]
```

Чтобы сохранить результат в файл:

```bash
go run . -query "project: ABC" -top 100 > issues.json
```

## Частые ошибки

`missing base URL`

Укажите адрес YouTrack через `-base-url`, `YOUTRACK_BASE_URL` или `base_url` в
конфиге.

`missing token`

Укажите permanent token через `-token`, `YOUTRACK_TOKEN` или `token` в конфиге.

`base url must include scheme and host`

Адрес должен содержать схему и host, например:

```text
https://youtrack.example.com
```

`unexpected status 401`

Токен не указан, неверный или не имеет нужных прав.

`unexpected status 403`

У токена нет доступа к запрошенным задачам или проекту.

## Разработка

Проверить сборку и тесты:

```bash
go test ./...
go build ./...
```

Отформатировать код:

```bash
gofmt -w .
```
