package client

// import (
// 	"be-golang-project/consts"
// 	"be-golang-project/models/db"
// 	"be-golang-project/models/handler"
// 	"be-golang-project/models/payload"
// 	"be-golang-project/models/validation_"
// 	"net/http"
// 	"strings"

// 	"gorm.io/gorm"
// )

// func (parentCtx *Client) Register(context_ *handler.Context) {
// 	var (
// 		clientRequest db.Client
// 		clientID      db.Client
// 		tx            *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB).Begin()
// 	)

// 	defer func() {
// 		tx.Rollback()
// 	}()

// 	resp.W = context_.Value.Writer

// 	if err := context_.ParseRequest(&clientRequest); err != nil {
// 		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

// 		return
// 	}

// 	if err := validation_.Validate(clientRequest, "register", "client", tx); err != nil {
// 		if strings.Contains(err.Error(), "Duplicate") {
// 			resp.SendResponse(http.StatusOK, consts.DuplicateEntryRC, consts.DuplicateEntryMessage, err.Error(), err)

// 			return
// 		}

// 		resp.SendResponse(http.StatusOK, consts.InvalidRequestBodyRC, consts.InvalidRequestBodyMessage, nil, err)

// 		return
// 	}

// 	clientRequest.IsNew = 1
// 	if err := tx.Create(&clientRequest).Error; err != nil {
// 		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

// 		return
// 	}

// 	if err := tx.Where("client_name = ?", clientRequest.ClientName).Find(&clientID).Error; err != nil {
// 		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

// 		return
// 	}

// 	if err := tx.Model(&db.User{}).Where("id = ?", context_.Value.Payload.(*payload.Payload).UserID).Update("company_id", clientID.ID).Error; err != nil {
// 		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

// 		return
// 	}

// 	tx.Commit()

// 	resp.SendResponse(http.StatusCreated, consts.CreatedRC, consts.CreatedMessage, "Successfully registered new client", nil)

// 	return
// }
