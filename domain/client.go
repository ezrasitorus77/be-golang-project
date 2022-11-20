package domain

import (
	"be-golang-project/domain/database"
)

type (
	ClientRepository interface {
		Get(userID string) (e error, client database.Client)
	}

	ClientService interface {
		Get(userID string) (e error, client database.Client)
	}
)
