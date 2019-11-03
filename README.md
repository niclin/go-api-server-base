# Golang api server base

- Gin
- Gorm
- 封裝 logger
- 封裝 router

## Godotenv

```
# MySQL 連接位置
MYSQL_DSN="db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local"
GIN_MODE="debug"
```

## Go Mod

用 go mod 管理套件

```
go run main.go
```

## Basic API

`/api/v1/ping`
