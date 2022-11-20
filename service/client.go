package service

import (
	"be-golang-project/domain"
	"be-golang-project/domain/database"
	"be-golang-project/repository"
)

type clientService struct {
	repo domain.ClientRepository
}

var ClientService domain.ClientService

func init() {
	ClientService = &clientService{
		repo: repository.ClientRepo,
	}
}

func (obj *clientService) Get(userID string) (e error, client database.Client) {
	return obj.repo.Get(userID)
}
