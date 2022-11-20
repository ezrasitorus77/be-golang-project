package user

// import (
// 	"be-golang-project/consts"
// 	"be-golang-project/models/db"
// 	"be-golang-project/models/handler"
// 	"be-golang-project/models/validation_"
// 	"net/http"
// 	"strings"

// 	"gorm.io/gorm"
// )

// func (parentCtx *User) Register(context_ *handler.Context) {
// 	var (
// 		userRequest db.User
// 		tx          *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB).Begin()
// 	)

// 	defer func() {
// 		tx.Rollback()
// 	}()

// 	resp.W = context_.Value.Writer

// 	if err := context_.ParseRequest(&userRequest); err != nil {
// 		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

// 		return
// 	}

// 	if err := validation_.Validate(userRequest, "register", "user", tx); err != nil {
// 		if strings.Contains(err.Error(), "Duplicate") {
// 			resp.SendResponse(http.StatusOK, consts.DuplicateEntryRC, consts.DuplicateEntryMessage, err.Error(), err)

// 			return
// 		}

// 		resp.SendResponse(http.StatusOK, consts.InvalidRequestBodyRC, consts.InvalidRequestBodyMessage, nil, err)

// 		return
// 	}

// 	userRequest.IsNew = 1
// 	if err := tx.Create(&userRequest).Error; err != nil {
// 		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

// 		return
// 	}

// 	tx.Commit()

// 	resp.SendResponse(http.StatusCreated, consts.CreatedRC, consts.CreatedMessage, "Successfully registered new user", nil)

// 	return
// }
