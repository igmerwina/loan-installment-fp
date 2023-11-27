package config

import (
	"fmt"
	"goloan/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Connect() {
	if err != nil {
		panic(err)
	}

	var (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "golang"
	)

	psqlInfo := fmt.Sprintf(`
	host=%s 
	port=%d 
	user=%s 	
	password=%s 
	dbname=%s 
	sslmode=disable`,
		host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Sakses konek tu DB\n", psqlInfo)

	db.AutoMigrate(&model.Installment{}, &model.Loan{}, &model.MaxLoan{})
}

func GetDB() *gorm.DB {
	return db
}
