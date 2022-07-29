package infrastructure

import (
	"be-golang-project/repository/orm"
)

var (
	DB orm.Database
)

func ConfigInit() error {
	DB.User = "root"
	DB.Password = "10089677esp"
	DB.DBName = "umkm"
	DB.DBHost = "127.0.0.1"
	DB.DBPort = "3306"

	if err := DB.Init(); err != nil {
		return err
	}

	return nil
}
