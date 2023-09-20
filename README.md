Бибилотека для работы с файлами конфигурации
go get -u github.com/spf13/viper

Docker
docker pull postgres
docker run --name=todo-db -e POSTGRES_PASSWORD='root' -p 5436:5432 -d --rm postgres
docker ps

Миграции
migrate create -ext sql -dir ./schema -seq init
migrate -path ./schema -database 'postgres://postgres:root@localhost:5436/postgres?sslmode=disable' up

Подключение к БД
docker exec -it 155b0128516d /bin/bash
root@155b0128516d:/# psql -U postgres
psql