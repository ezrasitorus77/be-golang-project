package vendor

import (
	"be-golang-project/consts"
	"be-golang-project/models/db"
	"be-golang-project/models/handler"
	"be-golang-project/models/payload"
	"be-golang-project/models/request"
	"be-golang-project/models/validation_"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

func (parentCtx *Vendor) Procurement(context_ *handler.Context) {
	var method string = context_.Value.Request.Method

	switch method {
	case "GET":
		getProcurements(context_)
	case "POST":
		addProcurement(context_)
	case "PUT":
		editProcurement(context_)
	case "DELETE":
		deleteProcurement(context_)
	}
}

func getProcurements(context_ *handler.Context) {
	var (
		procurements []db.Procurement
		user         db.User
		DB           *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB)
		userId       int      = context_.Value.Payload.(*payload.Payload).UserID
	)

	resp.W = context_.Value.Writer

	if err := DB.Where("id = ?", userId).Find(&user).Error; err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	if err := DB.Where("company_id = ?", user.CompanyID).Find(&procurements).Error; err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	resp.SendResponse(http.StatusOK, consts.SuccessRC, consts.SuccessMessage, procurements, nil)

	return
}

func addProcurement(context_ *handler.Context) {
	var (
		procurement db.Procurement
		user        db.User
		DB          *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB)
		userId      int      = context_.Value.Payload.(*payload.Payload).UserID
	)

	resp.W = context_.Value.Writer

	if err := context_.ParseRequest(&procurement); err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	if err := validation_.Validate(procurement, "procurement", "", DB); err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			resp.SendResponse(http.StatusOK, consts.DuplicateEntryRC, consts.DuplicateEntryMessage, err.Error(), err)

			return
		}

		resp.SendResponse(http.StatusOK, consts.InvalidRequestBodyRC, consts.InvalidRequestBodyMessage, nil, err)

		return
	}

	if err := DB.Where("id = ?", userId).Find(&user).Error; err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	procurement.VendorID = user.CompanyID
	if user.UserRole == 0 {
		procurement.EditedBySuper = 1
	}

	if err := DB.Create(&procurement).Error; err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	resp.SendResponse(http.StatusCreated, consts.CreatedRC, consts.CreatedMessage, "Successfully create new procurement", nil)

	return

}

func editProcurement(context_ *handler.Context) {
	var (
		procurement db.Procurement
		user        db.User
		DB          *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB)
		userId      int      = context_.Value.Payload.(*payload.Payload).UserID
	)

	resp.W = context_.Value.Writer

	if err := context_.ParseRequest(&procurement); err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	if err := validation_.Validate(procurement, "procurement", "", DB); err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			resp.SendResponse(http.StatusOK, consts.DuplicateEntryRC, consts.DuplicateEntryMessage, err.Error(), err)

			return
		}

		resp.SendResponse(http.StatusOK, consts.InvalidRequestBodyRC, consts.InvalidRequestBodyMessage, nil, err)

		return
	}

	if err := DB.Where("id = ?", userId).Find(&user).Error; err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	if user.UserRole == 0 {
		procurement.EditedBySuper = 1
	}

	if err := DB.Save(&procurement).Error; err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	resp.SendResponse(http.StatusCreated, consts.UpdatedRC, consts.UpdatedMessage, "Successfully update procurement", nil)

	return

}

func deleteProcurement(context_ *handler.Context) {
	var (
		tx                    *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB).Begin()
		procurementsIDRequest request.DeleteIDsRequest
	)

	defer func() {
		tx.Rollback()
	}()

	resp.W = context_.Value.Writer

	if err := context_.ParseRequest(&procurementsIDRequest); err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	for _, i := range procurementsIDRequest.ID {
		if err := tx.Where("id = ?", i).Delete(&db.Procurement{}).Error; err != nil {
			resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

			return
		}
	}

	tx.Commit()

	resp.SendResponse(http.StatusCreated, consts.DeletedRC, consts.DeletedMessage, "Successfully deleted procurement(s)", nil)

	return

}
