package main

import (
	"log"

	"gitlab.com/raihanlh/messenger-api/config"
	"gitlab.com/raihanlh/messenger-api/internal/model"
	"gitlab.com/raihanlh/messenger-api/pkg/postgres"
)

func main() {
	conf := config.New()

	db := postgres.New(conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPassword, conf.DBName, conf.DBTimezone)
	err := db.AutoMigrate(model.Models...)

	// migrator := db.Migrator()
	// if migrator.HasTable(&model.User{}) {
	// 	if err := migrator.DropColumn(&model.User{}, "birth_date"); err != nil {
	// 		log.Println("Column birth_date doesn't exist")
	// 	}
	// }

	if err != nil {
		panic(err)
	}
	log.Println("Migration success")
}
