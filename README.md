# FinBot
## Бот для учета финансов 

### Запуск
```bash
go build -o finbot cmd/main.go
echo "MONGODB_URI=''\nTOKEN=''" > .env
./finbot
```

### Docker
```bash
docker-compose --env-file .env up
```

### Использование
В запросе должны присутсвовать: 
* Стоимость - 1 или 2 место
* Категория - 1 или 2 место
* Комментарий - в конце *опционально*