package payload

import (
	"be-golang-project/consts"
	"errors"
	"time"
)

type Payload struct {
	UserID    int       `gorm:"column:user_id" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	ExpiredAt time.Time `gorm:"column:expired_at" json:"expired_at"`
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return errors.New(consts.ErrExpiredToken)
	}

	return nil
}

func NewPayload(userID int, duration time.Duration) (*Payload, error) {
	payload := &Payload{
		UserID:    userID,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}
