package infrastructure

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joshuaetim/frontdesk/domain/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func DB() *gorm.DB {
	var dialect gorm.Dialector
	dsn := os.Getenv("DATABASE_URL")

	if os.Getenv("APP_ENV") == "testing" {
		if mem := os.Getenv("SQLITE_MEMORY"); mem != "" {
			dialect = sqlite.Open(mem)
		} else {
			dialect = sqlite.Open(dsn)
		}
	} else {
		dialect = mysql.Open(dsn)
	}

	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database (%v)(%v): %v", dialect.Name(), dsn, err)
		return nil
	}

	db.AutoMigrate(&model.User{}, &model.Staff{}, &model.Visitor{})
	return db
}
