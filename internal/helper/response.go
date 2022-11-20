package helper

import (
	"be-golang-project/internal/consts"
	"net/http"

	"github.com/ezrasitorus77/http-handler/domain/delivery"
	log "github.com/ezrasitorus77/http-handler/helper"
)

func InternalServerError(w http.ResponseWriter) {
	log.Send(w, http.StatusInternalServerError, consts.RCInternalServerError, delivery.ResponseData{
		Description:  consts.InternalServerErrorMessage,
		ErrorMessage: consts.ServerErrorMessage,
	})
}
