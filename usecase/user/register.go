package user

import (
	"be-golang-project/consts"
	"be-golang-project/models/db"
	"be-golang-project/models/handler"
	"be-golang-project/models/validation_"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

// type User struct {
// 	handler.ParentContext
// }

// func New() *handler.ParentContext {
// 	var newUser handler.ParentContext

// 	return &newUser
// }

func (parentCtx *User) Register(context_ *handler.Context) {
	var (
		user db.User
		tx   *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB).Begin()
	)

	fmt.Println("tx: ", &tx)

	defer func() {
		tx.Rollback()
	}()

	resp.W = context_.Value.Writer

	context_.ParseRequest(&user)
	if context_.Value.Error != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, context_.Value.Error)

		return
	}

	if context_.Value.Error = validation_.Validate(user, "register"); context_.Value.Error != nil {
		resp.SendResponse(http.StatusOK, consts.InvalidRequestBodyRC, consts.InvalidRequestBodyMessage, nil, context_.Value.Error)

		return
	}

	if err := tx.Create(&db.User{
		Name:     user.Name,
		Password: validation_.HashPassword(user.Password, parentCtx.Salt),
		UserName: user.UserName,
		Email:    user.Email,
	}).Error; err != nil {
		context_.Value.Error = err
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, context_.Value.Error)

		return
	}

	tx.Commit()

	resp.SendResponse(http.StatusCreated, consts.CreatedRC, consts.CreatedMessage, "Successfully registered new user", nil)

	return
}
