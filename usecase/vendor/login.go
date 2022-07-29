package vendor

import (
	"be-golang-project/consts"
	"be-golang-project/models/db"
	"be-golang-project/models/handler"
	"be-golang-project/models/interface_"
	"be-golang-project/models/jwt"
	"be-golang-project/models/payload"
	"be-golang-project/models/validation_"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func (parentCtx *Vendor) Login(context_ *handler.Context) {
	var (
		tokenMaker  interface_.Maker
		newToken    string
		vendor      db.Vendor
		requestBody db.Vendor
		db          *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB)
		err         error
	)

	resp.W = context_.Value.Writer

	if err = context_.ParseRequest(&requestBody); err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	if err = db.Where("username = ?", requestBody.UserName).Find(&vendor).Error; err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	if vendor.ID == 0 {
		resp.SendResponse(http.StatusOK, consts.UserNotFoundRC, consts.UserNotFoundMessage, nil, nil)

		return
	}

	if ok := validation_.VerifyPassword(vendor.Password, requestBody.Password, parentCtx.Salt); !ok {
		resp.SendResponse(http.StatusForbidden, consts.CredentialNotMatchRC, consts.GeneralForbiddenMessage, consts.CredentialNotMatchMessage, nil)

		return
	}

	tokenMaker, err = jwt.NewJWTMaker(consts.JWTSecretKey)
	if err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	newToken, err = tokenMaker.CreateToken(vendor.ID, time.Minute)
	if err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	resp.SendResponse(http.StatusOK, consts.SuccessRC, consts.SuccessMessage, newToken, nil)

	return
}

func (parentCtx *Vendor) Index(context_ *handler.Context) {
	// var (
	// 	user db.Vendor
	// 	db          *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB)
	// )

	resp.W = context_.Value.Writer

	// if

	resp.SendResponse(http.StatusOK, consts.SuccessRC, consts.SuccessMessage, context_.Value.Payload.(*payload.Payload), nil)

	return
}
