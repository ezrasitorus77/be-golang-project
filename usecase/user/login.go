package user

// import (
// 	"be-golang-project/consts"
// 	"be-golang-project/models/db"
// 	"be-golang-project/models/handler"
// 	"be-golang-project/models/interface_"
// 	"be-golang-project/models/jwt"
// 	"be-golang-project/models/validation_"
// 	"net/http"
// 	"time"

// 	"gorm.io/gorm"
// )

// func (parentCtx *User) Login(context_ *handler.Context) {
// 	var (
// 		tokenMaker  interface_.Maker
// 		newToken    string
// 		user        db.User
// 		requestBody db.User
// 		DB          *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB)
// 		err         error
// 	)

// 	resp.W = context_.Value.Writer

// 	if err = context_.ParseRequest(&requestBody); err != nil {
// 		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

// 		return
// 	}

// 	if err = DB.Where("user_name = ?", requestBody.UserName).Find(&user).Error; err != nil {
// 		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

// 		return
// 	}

// 	if user.ID == 0 {
// 		resp.SendResponse(http.StatusOK, consts.UserNotFoundRC, consts.UserNotFoundMessage, nil, nil)

// 		return
// 	}

// 	if ok := validation_.VerifyPassword(user.Password, requestBody.Password, parentCtx.Salt); !ok {
// 		resp.SendResponse(http.StatusForbidden, consts.CredentialNotMatchRC, consts.GeneralForbiddenMessage, consts.CredentialNotMatchMessage, nil)

// 		return
// 	}

// 	tokenMaker, err = jwt.NewJWTMaker(consts.JWTSecretKey)
// 	if err != nil {
// 		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

// 		return
// 	}

// 	newToken, err = tokenMaker.CreateToken(user.ID, time.Minute)
// 	if err != nil {
// 		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

// 		return
// 	}

// 	// if user.IsNew == 1 {
// 	// 	http.Redirect(context_.Value.Writer, context_.Value.Request, "/user/profile", http.StatusSeeOther)

// 	// 	return
// 	// }

// 	resp.SendResponse(http.StatusOK, consts.SuccessRC, consts.SuccessMessage, newToken, nil)

// 	return
// }
