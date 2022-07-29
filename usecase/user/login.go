package user

import (
	"be-golang-project/consts"
	"be-golang-project/models/db"
	"be-golang-project/models/handler"
	"be-golang-project/models/interface_"
	"be-golang-project/models/token"
	"be-golang-project/models/validation_"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func Login() *User {
	var existUser User

	return &existUser
}

func (parentCtx *User) Login(context_ *handler.Context) {
	var (
		tokenMaker  interface_.Maker
		newToken    string
		user        db.User
		requestBody db.User
		tx          *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB).Begin()
	)

	defer func() {
		tx.Rollback()
	}()

	resp.W = context_.Value.Writer

	context_.ParseRequest(&requestBody)
	if context_.Value.Error != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, context_.Value.Error)

		return
	}

	if err := tx.Where("username = ?", requestBody.UserName).Find(&user).Error; err != nil {
		context_.Value.Error = err
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, context_.Value.Error)

		return
	}

	if user.ID == 0 {
		resp.SendResponse(http.StatusOK, consts.UserNotFoundRC, consts.UserNotFoundMessage, nil, nil)

		return
	}

	if ok := validation_.VerifyPassword(user.Password, requestBody.Password, parentCtx.Salt); !ok {
		resp.SendResponse(http.StatusForbidden, consts.CredentialNotMatchRC, consts.GeneralForbiddenMessage, consts.CredentialNotMatchMessage, nil)

		return
	}

	tokenMaker, context_.Value.Error = token.NewJWTMaker(consts.JWTSecretKey)
	if context_.Value.Error != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, context_.Value.Error)

		return
	}

	newToken, context_.Value.Error = tokenMaker.CreateToken(user.ID, time.Minute)
	if context_.Value.Error != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, context_.Value.Error)

		return
	}

	resp.SendResponse(http.StatusOK, consts.SuccessRC, consts.SuccessMessage, newToken, nil)

	return
}

func (parentCtx *User) Index(context_ *handler.Context) {
	// var (
	// 	tokenMaker interface_.Maker
	// 	userToken  string = context_.Value.Request.Header.Get("Token")
	// 	payload    *db.Token
	// )

	resp.W = context_.Value.Writer

	resp.SendResponse(http.StatusOK, consts.SuccessRC, consts.SuccessMessage, context_.Value.Payload.(*db.Token), nil)

	return
}
