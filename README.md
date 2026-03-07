Задания к курсу Backend разработка на GoLang 2026

## Задание 1 — GitHub Repository Info CLI

CLI-инструмент на Go для получения информации о репозитории GitHub через публичный API.

### Что делает

Запрашивает у пользователя репозиторий, отправляет HTTP-запрос к GitHub API и выводит:

- Название репозитория
- Описание
- Количество звёзд
- Количество форков
- Дату создания
- Ссылку на репозиторий

### Запуск

```bash
cd task1
go run main.go
```

После запуска введите репозиторий в одном из форматов:

```
owner/repo
https://github.com/owner/repo
```

**Пример:**

```
Enter GitHub repo (owner/repo or full URL): torvalds/linux
Name: linux
Description: Linux kernel source tree
Stars: 190000
Forks: 57000
Created at: April 8, 2011
Link: https://github.com/torvalds/linux
```

### Сборка

```bash
cd task1
go build -o ghinfo
./ghinfo
```
