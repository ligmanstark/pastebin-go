# Управление БД

## Создание в докере
```
docker run --name pastebin-postgres -e POSTGRES_USER=pastebin-db-user -e POSTGRES_PASSWORD=zhbb46 -e POSTGRES_DB=pastebin-db -p 5432:5432 -d postgres:15
```

## Вход в БД
```
psql -h localhost -U pastebin-db-user -d pastebin-db 
```

## Команды
INSERT INTO [NAME_TABLE] ([ROW]) VALUES ([VALUE]) - добавь строку в таблицу в таблицу NAME_TABLE, устанавливая значение столбца ROW равным VALUE

\i migration/005_create_users.sql - миграция