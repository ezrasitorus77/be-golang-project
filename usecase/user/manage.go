package user

// import (
// 	"be-golang-project/consts"
// 	"be-golang-project/models/db"
// 	"be-golang-project/models/handler"
// 	"be-golang-project/models/payload"
// 	"be-golang-project/models/request"
// 	"be-golang-project/models/validation_"
// 	"net/http"
// 	"strings"

// 	"gorm.io/gorm"
// )

// func (parentCtx *User) Manage(context_ *handler.Context) {
// 	var method string = context_.Value.Request.Method

// 	switch method {
// 	case "GET":
// 		getUsers(context_)
// 	case "POST":
// 		addUser(context_, parentCtx.Salt)
// 	case "PUT":
// 		editUser(context_)
// 	case "DELETE":
// 		deleteUsers(context_)
// 	}
// }

// func getUsers(context_ *handler.Context) {
// 	var (
// 		DB        *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB)
// 		users     []db.User
// 		companyID db.User
// 	)

// 	resp.W = context_.Value.Writer

// 	if err := DB.Where("id =?", context_.Value.Payload.(*payload.Payload).UserID).Find(&companyID).Error; err != nil {
// 		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

// 		return
// 	}

// 	if err := DB.Where("company_id = ? AND user_role != 0", companyID.ID).Find(&users).Error; err != nil {
// 		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

// 		return
// 	}

// 	resp.SendResponse(http.StatusOK, consts.SuccessRC, consts.SuccessMessage, users, nil)

// 	return
// }

// func addUser(context_ *handler.Context, salt []byte) {
// 	var (
// 		childUserRequest db.User
// 		DB               *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB)
// 		companyID        db.User
// 	)

// 	resp.W = context_.Value.Writer

// 	if err := context_.ParseRequest(&childUserRequest); err != nil {
// 		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

// 		return
// 	}

// 	if err := validation_.Validate(childUserRequest, "register", "user", DB); err != nil {
// 		if strings.Contains(err.Error(), "Duplicate") {
// 			resp.SendResponse(http.StatusOK, consts.DuplicateEntryRC, consts.DuplicateEntryMessage, err.Error(), err)

// 			return
// 		}

// 		resp.SendResponse(http.StatusOK, consts.InvalidRequestBodyRC, consts.InvalidRequestBodyMessage, nil, err)

// 		return
// 	}

// 	if err := DB.Where("id =?", context_.Value.Payload.(*payload.Payload).UserID).Find(&companyID).Error; err != nil {
// 		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

// 		return
// 	}

// 	if err := DB.Create(&db.User{
// 		UserName:    childUserRequest.UserName,
// 		Name:        childUserRequest.Name,
// 		Password:    validation_.HashPassword(childUserRequest.Password, salt),
// 		IDNumber:    childUserRequest.IDNumber,
// 		UserPhone:   childUserRequest.UserPhone,
// 		UserAddress: childUserRequest.UserAddress,
// 		CompanyID:   companyID.ID,
// 		UserRole:    1,
// 		IsNew:       1,
// 	}).Error; err != nil {
// 		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

// 		return
// 	}

// 	resp.SendResponse(http.StatusCreated, consts.CreatedRC, consts.CreatedMessage, "Successfully registered new user", nil)

// 	return
// }

// func editUser(context_ *handler.Context) {
// 	var (
// 		childUserRequest db.User
// 		DB               *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB)
// 	)

// 	resp.W = context_.Value.Writer

// 	if err := context_.ParseRequest(&childUserRequest); err != nil {
// 		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

// 		return
// 	}

// 	if err := validation_.Validate(childUserRequest, "register", "user", DB); err != nil {
// 		if strings.Contains(err.Error(), "Duplicate") {
// 			resp.SendResponse(http.StatusOK, consts.DuplicateEntryRC, consts.DuplicateEntryMessage, err.Error(), err)

// 			return
// 		}

// 		resp.SendResponse(http.StatusOK, consts.InvalidRequestBodyRC, consts.InvalidRequestBodyMessage, nil, err)

// 		return
// 	}

// 	if err := DB.Save(&childUserRequest).Error; err != nil {
// 		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

// 		return
// 	}

// 	resp.SendResponse(http.StatusCreated, consts.UpdatedRC, consts.UpdatedMessage, "Successfully update user profile", nil)

// 	return

// }

// func deleteUsers(context_ *handler.Context) {
// 	var (
// 		tx             *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB).Begin()
// 		usersIDRequest request.DeleteIDsRequest
// 	)

// 	defer func() {
// 		tx.Rollback()
// 	}()

// 	resp.W = context_.Value.Writer

// 	if err := context_.ParseRequest(&usersIDRequest); err != nil {
// 		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

// 		return
// 	}

// 	for _, i := range usersIDRequest.ID {
// 		if err := tx.Where("id = ?", i).Delete(&db.User{}).Error; err != nil {
// 			resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

// 			return
// 		}
// 	}

// 	tx.Commit()

// 	resp.SendResponse(http.StatusCreated, consts.DeletedRC, consts.DeletedMessage, "Successfully deleted user(s)", nil)

// 	return

// }
