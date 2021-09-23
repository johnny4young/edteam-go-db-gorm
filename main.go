package main

import "github.com/johnny4young/edteam-go-db-gorm/storage"

func main() {
	driver := storage.Postgres
	storage.New(driver)
}
