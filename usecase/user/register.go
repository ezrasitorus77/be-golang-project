package user

import (
	"be-golang-project/consts"
	"be-golang-project/models/db"
	"be-golang-project/models/handler"
	"be-golang-project/models/validation_"
	"net/http"

	"gorm.io/gorm"
)

func (parentCtx *User) Register(context_ *handler.Context) {
	var (
		user db.User
		tx   *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB).Begin()
	)

	defer func() {
		tx.Rollback()
	}()

	resp.W = context_.Value.Writer

	if err := context_.ParseRequest(&user); err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	if err := validation_.Validate(user, "register", "user"); err != nil {
		resp.SendResponse(http.StatusOK, consts.InvalidRequestBodyRC, consts.InvalidRequestBodyMessage, nil, err)

		return
	}

	if err := tx.Create(&db.User{
		UserName:    user.UserName,
		Name:        user.Name,
		Password:    validation_.HashPassword(user.Password, parentCtx.Salt),
		IDNumber:    user.IDNumber,
		UserPhone:   user.UserPhone,
		UserAddress: user.UserAddress,
		UserRole:    user.UserRole,
	}).Error; err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	tx.Commit()

	resp.SendResponse(http.StatusCreated, consts.CreatedRC, consts.CreatedMessage, "Successfully registered new user", nil)

	return
}
