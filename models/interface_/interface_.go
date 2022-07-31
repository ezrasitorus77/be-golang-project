package interface_

import (
	"be-golang-project/models/handler"
	"be-golang-project/models/payload"
	"time"
)

type (
	Maker interface {
		CreateToken(userID int, duration time.Duration) (string, error)
		VerifyToken(token string) (*payload.Payload, error)
	}

	User interface {
		Register(*handler.Context)
		Login(*handler.Context)
		Index(*handler.Context)
		Profile(*handler.Context)
	}

	Vendor interface {
		Register(*handler.Context)
		// Profile(*handler.Context)
	}
)
