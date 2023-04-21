run Migrate

$ migrate create -ext=mysql -dir=sql/migrations -seq init

## Generate sqlc

sqlc generate

## Verify if migration has been created

docker-compose exec mysql bash
mysql -root -p chat_test;
Enter Pass: root
show tables;

## Build tiktokenlib

docker-compose exec chatservice bash

cd tiktoken-cffi/

cargo build --release

## References

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

## Issues

https://github.com/rubenv/sql-migrate/issues/213
