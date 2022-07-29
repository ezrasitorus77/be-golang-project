package orm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DSN string = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	err error
	DB  *gorm.DB
)

type Database struct {
	User     string
	Password string
	DBName   string
	DBHost   string
	DBPort   string
}

func (db *Database) Init() error {
	dsn := fmt.Sprintf(DSN, db.User, db.Password, db.DBHost, db.DBPort, db.DBName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func CreateDBSession() *gorm.DB {
	return DB.Debug()
}
