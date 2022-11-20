package controller

import (
	"be-golang-project/domain"
	"be-golang-project/domain/database"
	"be-golang-project/internal/consts"
	"be-golang-project/internal/helper"
	"be-golang-project/service"
	"net/http"

	"github.com/ezrasitorus77/http-handler/domain/delivery"
	log "github.com/ezrasitorus77/http-handler/helper"
	"gorm.io/gorm"
)

type clientController struct {
	service domain.ClientService
}

var ClientController clientController

func init() {
	ClientController = clientController{
		service: service.ClientService,
	}
}

func (obj *clientController) Get(w http.ResponseWriter, r *http.Request) {
	var (
		e      error
		params map[string]string = r.Context().Value("params").(map[string]string)
		userID string
		client database.Client
	)

	userID = params["id"]
	if userID == "" {
		log.Send(w, http.StatusBadRequest, consts.RCInvalidRequestBody, delivery.ResponseData{
			Description:  consts.InvalidRequestBodyParamMessage,
			ErrorMessage: "User ID param is empty",
		})

		return
	}

	e, client = obj.service.Get(userID)
	if e != nil {
		if e == gorm.ErrRecordNotFound {
			log.Send(w, http.StatusOK, consts.RCUserNotFound, delivery.ResponseData{
				Description:  consts.UserNotFoundMessage,
				ErrorMessage: "Client is not found",
			})

			return
		} else {
			helper.InternalServerError(w)

			return
		}
	}

	log.Send(w, http.StatusOK, consts.RCSuccess, delivery.ResponseData{
		Data: client,
	})

	return
}
