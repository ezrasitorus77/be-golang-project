package validation_

// import (
// 	"be-golang-project/internal/consts"
// 	"be-golang-project/models/db"
// 	"errors"
// 	"fmt"
// 	"regexp"

// 	validation "github.com/go-ozzo/ozzo-validation"
// 	"github.com/go-ozzo/ozzo-validation/is"
// 	"gorm.io/gorm"
// )

// var (
// 	AlphaSpace        = validation.Match(regexp.MustCompile(consts.AlphaSpaceRegex)).Error(consts.AlphaSpaceValidationMessage)
// 	AlphaLower        = validation.Match(regexp.MustCompile(consts.AlphaLowerRegex)).Error(consts.AlphaLowerValidationMessage)
// 	AlphaUpper        = validation.Match(regexp.MustCompile(consts.AlphaUpperRegex)).Error(consts.AlphaUpperValidationMessage)
// 	ComplexCharacters = validation.Match(regexp.MustCompile(consts.ComplexCharactersRegex)).Error(consts.ComplexCharactersValidationMessage)
// 	Digits            = validation.Match(regexp.MustCompile(consts.DigitOnlyRegex)).Error(consts.DigitOnlyValidationMessage)
// 	EmailRule         = is.Email.Error(consts.EmailValidationMessage)
// 	ExceptSpace       = validation.Match(regexp.MustCompile(consts.ExceptSpaceRegex))
// 	IsRequired        = validation.Required.Error(consts.RequiredValidationMessage)
// 	Length16          = validation.Length(16, 16).Error(consts.Length16ValidationMessage)
// 	LengthMin8Max20   = validation.Length(8, 20).Error(consts.LengthMin8Max20ValidationMessage)
// 	LengthMin15Max16  = validation.Length(15, 16).Error(consts.LengthMin15Max16ValidationMessage)
// 	LengthMax10       = validation.Length(1, 10).Error(consts.LengthMax10ValidationMessage)
// 	LengthMax20       = validation.Length(1, 20).Error(consts.LengthMax20ValidationMessage)
// 	LengthMax50       = validation.Length(1, 50).Error(consts.LengthMax50ValidationMessage)
// 	LengthMin10Max100 = validation.Length(10, 100).Error(consts.LengthMin10Max100ValidationMessage)
// 	LengthMin20Max500 = validation.Length(20, 500).Error(consts.LengthMin20Max500ValidationMessage)
// 	Symbols           = validation.Match(regexp.MustCompile(consts.SymbolsOnlyRegex)).Error(consts.SymbolsOnlyValidationMessage)
// )

// func Validate(form interface{}, interest, entity string, tx *gorm.DB) error {
// 	var (
// 		uniqueField []string
// 		uniqueForm  []string
// 	)

// 	switch interest {
// 	case "register":
// 		switch entity {
// 		case "user":
// 			model := form.(db.User)
// 			fieldCheck := db.User{}
// 			uniqueField = []string{"user_name", "id_number", "user_phone"}
// 			uniqueForm = []string{model.UserName, model.IDNumber, model.UserPhone}

// 			if err := validation.ValidateStruct(
// 				&model,
// 				validation.Field(&model.UserName, IsRequired, LengthMax20, ExceptSpace),
// 				validation.Field(&model.Name, IsRequired, AlphaSpace, LengthMax50),
// 				validation.Field(&model.Password, IsRequired, LengthMin8Max20, AlphaLower, AlphaUpper, Digits),
// 				validation.Field(&model.IDNumber, IsRequired, Digits, Length16),
// 				validation.Field(&model.UserPhone, IsRequired, Digits, LengthMin8Max20),
// 				validation.Field(&model.UserAddress, IsRequired, ComplexCharacters, LengthMin10Max100),
// 			); err != nil {
// 				return err
// 			}

// 			for i, f := range uniqueField {
// 				if err := tx.Where(fmt.Sprintf("%s = ?", f), uniqueForm[i]).Find(&fieldCheck).Error; err != nil {
// 					return err
// 				}

// 				if model.ID != 0 {
// 					if model.ID != fieldCheck.ID {
// 						return errors.New(fmt.Sprintf("Duplicate entry for %s field", f))
// 					}
// 				}

// 				if fieldCheck.ID != 0 {
// 					return errors.New(fmt.Sprintf("Duplicate entry for %s field", f))
// 				}
// 			}

// 			return nil

// 		case "vendor":
// 			model := form.(db.Vendor)
// 			fieldCheck := db.Vendor{}
// 			uniqueField = []string{"vendor_name", "vendor_phone", "vendor_website", "vendor_email", "npwp", "account_name", "account_number"}
// 			uniqueForm = []string{model.VendorName, model.VendorPhone, model.VendorWebsite, model.Email, model.NPWP, model.AccountName, model.AccountNumber}

// 			if err := validation.ValidateStruct(
// 				&model,
// 				validation.Field(&model.VendorName, IsRequired, LengthMax20, ComplexCharacters),
// 				validation.Field(&model.VendorField, IsRequired, Digits),
// 				validation.Field(&model.VendorType, IsRequired, Digits),
// 				validation.Field(&model.VendorAddress, IsRequired, ComplexCharacters, LengthMin10Max100),
// 				validation.Field(&model.VendorPhone, IsRequired, Digits, LengthMin8Max20),
// 				validation.Field(&model.Email, IsRequired, EmailRule),
// 				validation.Field(&model.NPWP, IsRequired, Digits, LengthMin15Max16),
// 				validation.Field(&model.AccountName, IsRequired, AlphaSpace, LengthMax50),
// 				validation.Field(&model.AccountNumber, IsRequired, Digits, LengthMax20),
// 				validation.Field(&model.Province, IsRequired, Digits),
// 				validation.Field(&model.City, IsRequired, Digits),
// 				validation.Field(&model.District, IsRequired, Digits),
// 			); err != nil {
// 				return err
// 			}

// 			for i, f := range uniqueField {
// 				if err := tx.Where(fmt.Sprintf("%s = ?", f), uniqueForm[i]).Find(&fieldCheck).Error; err != nil {
// 					return err
// 				}

// 				if model.ID != 0 {
// 					if model.ID != fieldCheck.ID {
// 						return errors.New(fmt.Sprintf("Duplicate entry for %s field", f))
// 					}
// 				}

// 				if fieldCheck.ID != 0 {
// 					return errors.New(fmt.Sprintf("Duplicate entry for %s field", f))
// 				}
// 			}

// 			return nil

// 		case "client":
// 			model := form.(db.Client)
// 			fieldCheck := db.Client{}
// 			uniqueField = []string{"client_name", "client_phone", "client_website", "client_email"}
// 			uniqueForm = []string{model.ClientName, model.ClientPhone, model.ClientWebsite, model.Email}

// 			if err := validation.ValidateStruct(
// 				&model,
// 				validation.Field(&model.ClientName, IsRequired, LengthMax20, ExceptSpace),
// 				validation.Field(&model.ClientParent, IsRequired, Digits),
// 				validation.Field(&model.ClientField, IsRequired, Digits),
// 				validation.Field(&model.ClientAddress, IsRequired, ComplexCharacters, LengthMin10Max100),
// 				validation.Field(&model.ClientPhone, IsRequired, Digits, LengthMin8Max20),
// 				validation.Field(&model.Email, IsRequired, EmailRule),
// 				validation.Field(&model.Province, IsRequired, Digits),
// 				validation.Field(&model.City, IsRequired, Digits),
// 				validation.Field(&model.District, IsRequired, Digits),
// 			); err != nil {
// 				return err
// 			}

// 			for i, f := range uniqueField {
// 				if err := tx.Where(fmt.Sprintf("%s = ?", f), uniqueForm[i]).Find(&fieldCheck).Error; err != nil {
// 					return err
// 				}

// 				if model.ID != 0 {
// 					if model.ID != fieldCheck.ID {
// 						return errors.New(fmt.Sprintf("Duplicate entry for %s field", f))
// 					}
// 				}

// 				if fieldCheck.ID != 0 {
// 					return errors.New(fmt.Sprintf("Duplicate entry for %s field", f))
// 				}
// 			}

// 			return nil
// 		}
// 	case "procurement":
// 		model := form.(db.Procurement)

// 		if err := validation.ValidateStruct(
// 			&model,
// 			validation.Field(&model.Category, IsRequired),
// 			validation.Field(&model.Title, IsRequired, LengthMax20),
// 			validation.Field(&model.Body, IsRequired, LengthMin20Max500),
// 			validation.Field(&model.PriceStart, IsRequired, Digits),
// 			validation.Field(&model.PriceEnd, IsRequired, Digits),
// 			validation.Field(&model.PaymentMethod, IsRequired, Digits),
// 		); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
