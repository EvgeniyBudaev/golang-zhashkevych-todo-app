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

Для работы с БД
go get -u github.com/jmoiron/sqlx

Go postgres driver for Go's database/sql package
go get -u github.com/lib/pq

Для получения паролей в приложении используются переменные окружения
go get -u github.com/joho/godotenv

Бибилиотека для логирования
go get -u github.com/sirupsen/logrus

JWT
go get -u github.com/dgrijalva/jwt-go