package repository

import (
	"be-golang-project/config"
	"be-golang-project/domain"
	"be-golang-project/domain/database"

	"gorm.io/gorm"
)

type clientRepo struct {
	db *gorm.DB
}

var ClientRepo domain.ClientRepository

func init() {
	ClientRepo = &clientRepo{
		db: config.DB,
	}
}

func (obj *clientRepo) Get(userID string) (e error, client database.Client) {
	var user database.User

	if e = obj.db.Where("id = ?", userID).First(&user).Error; e != nil {
		return
	}

	if e = obj.db.Where("id = ?", user.CompanyID).First(&client).Error; e != nil {
		return
	}

	return
}
