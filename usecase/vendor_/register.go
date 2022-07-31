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

func (parentCtx *Vendor) Register(context_ *handler.Context) {
	var (
		vendorRequest db.Vendor
		vendorID      db.Vendor
		tx            *gorm.DB = context_.ChildCtx.Value("DB").(*gorm.DB).Begin()
	)

	defer func() {
		tx.Rollback()
	}()

	resp.W = context_.Value.Writer

	if err := context_.ParseRequest(&vendorRequest); err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	if err := validation_.Validate(vendorRequest, "register", "vendor", tx); err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			resp.SendResponse(http.StatusOK, consts.DuplicateEntryRC, consts.DuplicateEntryMessage, err.Error(), err)

			return
		}

		resp.SendResponse(http.StatusOK, consts.InvalidRequestBodyRC, consts.InvalidRequestBodyMessage, nil, err)

		return
	}

	if err := tx.Create(&db.Vendor{
		VendorName:    vendorRequest.VendorName,
		VendorField:   vendorRequest.VendorField,
		VendorType:    vendorRequest.VendorType,
		VendorAddress: vendorRequest.VendorAddress,
		VendorPhone:   vendorRequest.VendorPhone,
		VendorWebsite: vendorRequest.VendorWebsite,
		Email:         vendorRequest.Email,
		NPWP:          vendorRequest.NPWP,
		SocialMedia:   vendorRequest.SocialMedia,
		Province:      vendorRequest.Province,
		City:          vendorRequest.City,
		District:      vendorRequest.District,
		Avatar:        vendorRequest.Avatar,
		IsNew:         1,
	}).Error; err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	if err := tx.Where("vendor_name = ?", vendorRequest.VendorName).Find(&vendorID).Error; err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	if err := tx.Model(&db.User{}).Where("id = ?", context_.Value.Payload.(*payload.Payload).UserID).Update("company_id", vendorID.ID).Error; err != nil {
		resp.SendResponse(http.StatusInternalServerError, consts.GeneralInternalServerErrorRC, consts.GeneralInternalServerErrorMessage, nil, err)

		return
	}

	tx.Commit()

	resp.SendResponse(http.StatusCreated, consts.CreatedRC, consts.CreatedMessage, "Successfully registered new vendor", nil)

	return
}
