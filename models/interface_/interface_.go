package interface_

import (
	"be-golang-project/models/db"
	"be-golang-project/models/handler"
	"time"
)

type (
	Maker interface {
		CreateToken(userID int, duration time.Duration) (string, error)
		VerifyToken(token string) (*db.Token, error)
	}

	User interface {
		Register(*handler.Context)
		Login(*handler.Context)
		Index(*handler.Context)
	}
)
