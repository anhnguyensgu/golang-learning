package databaseManagement

import (
	"log"
	"todo-app/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func SetupDatabase() {
	dns := "host=localhost user=postgres password=123456 dbname=todo port=5432"

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database \n", err)
	}

	log.Println("connected")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migration")

	db.AutoMigrate(&entities.Account{})

	DB = DbInstance{
		Db: db,
	}
}
