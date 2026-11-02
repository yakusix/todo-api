# todo-api

REST API для списка задач на Go. Пет-проект для практики.

## Возможности

| Метод | Путь | Описание |
|-------|------|----------|
| GET | `/tasks` | список всех задач |
| POST | `/tasks` | создать задачу |
| GET | `/tasks/{id}` | получить задачу по ID |
| DELETE | `/tasks/{id}` | удалить задачу |
| PATCH | `/tasks/{id}/done` | отметить задачу выполненной |


Данные хранятся в памяти и сбрасываются при перезапуске сервера

## Структура проекта

```
main.go  точка входа, роуты, запуск сервера на порту 8080
handler/ HTTP-слой : разбор запроса, статусы, json
usecase/ бизнес логика
repository/ хранилище задач 
```
Запрос проходит путь `handler - usecase - repository` каждый слой знает о соседнем через интерфейс.

## Запуск

Нужен Go 1.22 или новее.

```bash
git clone https://github.com/yakusix/todo-api.git
cd todo-api
go run .
```

Сервер стартует на `http://localhost:8080`.

## Примеры запросов

Создать задачу:
```bash
curl -X POST localhost:8080/tasks -d '{"title":"купить продукты","description":"молоко, хлеб"}'
```

Список задач:
```bash
curl localhost:8080/tasks
```

Отметить выполненной:
```bash
curl -X PATCH localhost:8080/tasks/1/done
```

Удалить:
```bash
curl -X DELETE localhost:8080/tasks/1
```

## Тесты

```bash
go test -race ./...
```

Есть тесты бизнес-логики и тест на конкурентное создание задач (проверка `-race`).

## Планы

- [ ] PostgreSQL вместо хранения в памяти
- [ ] GitHub Actions: тесты на каждый push
- [ ] Тесты handler'ов (`httptest`)
- [ ] Docker