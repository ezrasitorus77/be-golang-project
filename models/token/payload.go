package token

import (
	"be-golang-project/models/db"
	"time"

	"github.com/google/uuid"
)

func NewPayload(userID int, duration time.Duration) (*db.Token, error) {
	userUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &db.Token{
		UserID:    userID,
		UUID:      userUUID,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}
