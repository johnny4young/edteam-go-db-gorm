package main

import (
	"github.com/johnny4young/edteam-go-db-gorm/model"
	"github.com/johnny4young/edteam-go-db-gorm/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	storage.DB().AutoMigrate(
		&model.Product{},
		&model.InvoiceHeader{},
		&model.InvoiceItem{},
	)

	// Create
	product1 := model.Product{
		Name:  "Go's Course",
		Price: 120,
	}

	obs := "testing with GO"
	product2 := model.Product{
		Name:         "Go's Testing",
		Price:        80,
		Observations: &obs,
	}
	storage.DB().Create(&product1)
	storage.DB().Create(&product2)
}
