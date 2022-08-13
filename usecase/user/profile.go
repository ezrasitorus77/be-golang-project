package user

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

func (parentCtx *User) Profile(context_ *handler.Context) {
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
		user   db.User
		DB     *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB)
		userID int      = context_.Value.Payload.(*payload.Payload).UserID
	)

	resp.W = context_.Value.Writer

	if err := DB.Where("id = ?", userID).Find(&user).Error; err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	resp.SendResponse(http.StatusOK, consts.SuccessRC, consts.SuccessMessage, user, nil)

	return
}

func editProfile(context_ *handler.Context) {
	var (
		user db.User
		DB   *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB)
	)

	if err := context_.ParseRequest(&user); err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	if err := validation_.Validate(user, "register", "user", DB); err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			resp.SendResponse(http.StatusOK, consts.DuplicateEntryRC, consts.DuplicateEntryMessage, err.Error(), err)

			return
		}

		resp.SendResponse(http.StatusOK, consts.InvalidRequestBodyRC, consts.InvalidRequestBodyMessage, nil, err)

		return
	}

	if err := DB.Save(&user).Error; err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	resp.SendResponse(http.StatusCreated, consts.UpdatedRC, consts.UpdatedMessage, "Successfully update user profile", nil)

	return
}
