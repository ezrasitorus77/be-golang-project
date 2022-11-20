package config

import (
	"be-golang-project/internal/consts"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	var (
		dsn        string
		e          error
		dbUser     string
		dbPassword string
		dbHost     string
		dbPort     string
		dbName     string
	)

	if e = godotenv.Load(); e != nil {
		log.Fatal(e)
	}

	dbUser = os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatal(consts.ErrMandatoryEnv)
	}

	dbPassword = os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatal(consts.ErrMandatoryEnv)
	}

	dbHost = os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Fatal(consts.ErrMandatoryEnv)
	}

	dbPort = os.Getenv("DB_PORT")
	if dbPort == "" {
		log.Fatal(consts.ErrMandatoryEnv)
	}

	dbName = os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal(consts.ErrMandatoryEnv)
	}

	dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	DB, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if e != nil {
		log.Fatal(e)
	} else {
		DB.Debug()
	}
}
