package vendor

import (
	"be-golang-project/consts"
	"be-golang-project/models/db"
	"be-golang-project/models/handler"
	"be-golang-project/models/payload"
	"be-golang-project/models/validation_"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

func (parentCtx *Vendor) Profile(context_ *handler.Context) {
	var method string = context_.Value.Request.Method

	switch method {
	case "GET":
		getProfile(context_)
	case "PUT":
		editProfile(context_)
	}
}

func getProfile(context_ *handler.Context) {
	var (
		vendor db.Vendor
		user   db.User
		DB     *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB)
		userId int      = context_.Value.Payload.(*payload.Payload).UserID
	)

	resp.W = context_.Value.Writer

	if err := DB.Where("id = ?", userId).Find(&user).Error; err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	if err := DB.Where("id = ?", user.CompanyID).Find(&vendor).Error; err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	resp.SendResponse(http.StatusOK, consts.SuccessRC, consts.SuccessMessage, vendor, nil)

	return
}

func editProfile(context_ *handler.Context) {
	var (
		vendor db.Vendor
		DB     *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB)
	)

	if err := context_.ParseRequest(&vendor); err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	if err := validation_.Validate(vendor, "register", "vendor", DB); err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			resp.SendResponse(http.StatusOK, consts.DuplicateEntryRC, consts.DuplicateEntryMessage, err.Error(), err)

			return
		}

		resp.SendResponse(http.StatusOK, consts.InvalidRequestBodyRC, consts.InvalidRequestBodyMessage, nil, err)

		return
	}

	if err := DB.Save(&vendor).Error; err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	resp.SendResponse(http.StatusCreated, consts.UpdatedRC, consts.UpdatedMessage, "Successfully update vendor profile", nil)

	return
}
