package postgres

import (
	"fmt"
	"log"

	"gitlab.com/raihanlh/messenger-api/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(host string, port string, username string, password string, dbname string, timezone string) *gorm.DB {
	dbParams := []any{host, port, username, password, dbname, timezone}
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		dbParams...)

	gormDB, err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{
		Logger: logger.GetLoggerGorm(),
	})

	if err != nil {
		panic(err)
	}

	log.Printf("Success Connecting to DB : %s\n", dbname)
	return gormDB
}