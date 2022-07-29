package db

import (
	"be-golang-project/consts"
	"errors"
	"time"

	"github.com/google/uuid"
)

type (
	User struct {
		ID        int       `gorm:"column:id" json:"user_id"`
		Name      string    `gorm:"column:name" json:"name"`
		UserName  string    `gorm:"column:username" json:"username"`
		Password  string    `gorm:"column:password" json:"password"`
		CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
		Email     string    `gorm:"column:email" json:"user_email"`
	}

	Token struct {
		ID        int       `gorm:"column:id" json:"token_id"`
		UserID    int       `gorm:"column:user_id" json:"user_id"`
		UUID      uuid.UUID `gorm:"column:uuid" json:"uuid"`
		TokenHash string    `gorm:"column:token_hash" json:"token_hash"`
		CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
		ExpiredAt time.Time `gorm:"column:expired_at" json:"expired_at"`
	}
)

func (tbl *User) TableName() string {
	return "user"
}

func (payload *Token) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return errors.New(consts.ErrExpiredToken)
	}

	return nil
}
