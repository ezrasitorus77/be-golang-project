package vendor

import (
	"be-golang-project/consts"
	"be-golang-project/models/db"
	"be-golang-project/models/handler"
	"be-golang-project/models/validation_"
	"net/http"

	"gorm.io/gorm"
)

func (parentCtx *Vendor) Register(context_ *handler.Context) {
	var (
		vendor db.Vendor
		tx     *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB).Begin()
	)

	defer func() {
		tx.Rollback()
	}()

	resp.W = context_.Value.Writer

	if err := context_.ParseRequest(&vendor); err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	if err := validation_.Validate(vendor, "profile"); err != nil {
		resp.SendResponse(http.StatusOK, consts.InvalidRequestBodyRC, consts.InvalidRequestBodyMessage, nil, err)

		return
	}

	if err := tx.Create(&db.Vendor{
		UserName: vendor.UserName,
		Password: validation_.HashPassword(vendor.Password, parentCtx.Salt),
		Email:    vendor.Email,
	}).Error; err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	tx.Commit()

	resp.SendResponse(http.StatusCreated, consts.CreatedRC, consts.CreatedMessage, "Successfully registered new user", nil)

	return
}
