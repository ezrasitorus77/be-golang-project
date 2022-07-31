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
	var (
		vendor db.Vendor
		user   db.User
		tx     *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB).Begin()
		method string   = context_.Value.Request.Method
		userId int      = context_.Value.Payload.(*payload.Payload).UserID
	)

	defer func() {
		tx.Rollback()
	}()

	resp.W = context_.Value.Writer

	switch method {
	case "GET":
		if err := tx.Where("id = ?", userId).Find(&user).Error; err != nil {
			resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

			return
		}

		if err := tx.Where("id = ?", user.CompanyID).Find(&vendor).Error; err != nil {
			resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

			return
		}

		resp.SendResponse(http.StatusOK, consts.SuccessRC, consts.SuccessMessage, vendor, nil)

		return

	case "POST":
		if err := context_.ParseRequest(&vendor); err != nil {
			resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

			return
		}

		if err := validation_.Validate(vendor, "register", "vendor", tx); err != nil {
			if strings.Contains(err.Error(), "Duplicate") {
				resp.SendResponse(http.StatusOK, consts.DuplicateEntryRC, consts.DuplicateEntryMessage, err.Error(), err)

				return
			}

			resp.SendResponse(http.StatusOK, consts.InvalidRequestBodyRC, consts.InvalidRequestBodyMessage, nil, err)

			return
		}

		if err := tx.Save(&vendor).Error; err != nil {
			resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

			return
		}

		tx.Commit()

		resp.SendResponse(http.StatusCreated, consts.UpdatedRC, consts.UpdatedMessage, "Successfully update vendor profile", nil)

		return

	default:
		return
	}

}
