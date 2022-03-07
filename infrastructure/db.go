package infrastructure

import (
	"log"
	"os"

	"github.com/joshuaetim/frontdesk/domain/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil
	}
	db.AutoMigrate(&model.User{}, &model.Staff{}, &model.Visitor{})
	return db
}
