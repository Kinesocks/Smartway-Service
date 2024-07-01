# Тестовое задание smartway
## How to build
Для сборки бинраного файла надо задать перемнные окружения для компилятора GO
для Windows :
```
$Env:GOOS = "linux"; $Env:CGO_ENABLED=0 
```
для Linux:
```
CGO_ENABLED=0 GOOS=linux
```
Далее компиляция:
```
go build -ldflags="-s -w" -tags=containers .
```
Сборка Docker контейнера:
```
docker build -t smartway_service .
```
Далее поднимаем все контейнеры про помощи Docker Compose:
```
docker compose up -d
```
