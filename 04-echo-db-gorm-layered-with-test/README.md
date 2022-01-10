# Howto


## Run the program using environment file

```bash
source template.env
go run main.go
```

## Run the program by using manually set environment

```bash
export DB_CONNECTION_STRING='root:Alch3mist@tcp(127.0.0.1:3306)/db_sirclo_api_gorm?charset=utf8&parseTime=True&loc=Local'
go run main.go
```