package validation_

import (
	"be-golang-project/consts"
	"be-golang-project/models/db"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"hash"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

var (
	AlphaSpace        = validation.Match(regexp.MustCompile(consts.AlphaSpaceRegex)).Error(consts.AlphaSpaceValidationMessage)
	AlphaLower        = validation.Match(regexp.MustCompile(consts.AlphaLowerRegex)).Error(consts.AlphaLowerValidationMessage)
	AlphaUpper        = validation.Match(regexp.MustCompile(consts.AlphaUpperRegex)).Error(consts.AlphaUpperValidationMessage)
	ComplexCharacters = validation.Match(regexp.MustCompile(consts.ComplexCharactersRegex)).Error(consts.ComplexCharactersValidationMessage)
	Digits            = validation.Match(regexp.MustCompile(consts.DigitOnlyRegex)).Error(consts.DigitOnlyValidationMessage)
	EmailRule         = is.Email.Error(consts.EmailValidationMessage)
	ExceptSpace       = validation.Match(regexp.MustCompile(consts.ExceptSpaceRegex))
	IsRequired        = validation.Required.Error(consts.RequiredValidationMessage)
	Length16          = validation.Length(16, 16).Error(consts.Length16ValidationMessage)
	LengthMin8Max20   = validation.Length(8, 20).Error(consts.LengthMin8Max20ValidationMessage)
	LengthMin15Max16  = validation.Length(15, 16).Error(consts.LengthMin15Max16ValidationMessage)
	LengthMax10       = validation.Length(1, 10).Error(consts.LengthMax10ValidationMessage)
	LengthMax20       = validation.Length(1, 20).Error(consts.LengthMax20ValidationMessage)
	LengthMax50       = validation.Length(1, 50).Error(consts.LengthMax50ValidationMessage)
	LengthMin10Max100 = validation.Length(10, 100).Error(consts.LengthMin10Max100ValidationMessage)
	Symbols           = validation.Match(regexp.MustCompile(consts.SymbolsOnlyRegex)).Error(consts.SymbolsOnlyValidationMessage)
)

func Validate(form interface{}, interest string) error {
	switch interest {
	case "profile":
		model := form.(db.Vendor)
		return validation.ValidateStruct(
			&model,
			validation.Field(&model.UserName, IsRequired, LengthMax20, ExceptSpace),
			validation.Field(&model.Password, IsRequired, LengthMin8Max20, AlphaLower, AlphaUpper, Digits),
			validation.Field(&model.Email, IsRequired, EmailRule),
			validation.Field(&model.IDNumber, IsRequired, Digits, Length16),
			validation.Field(&model.NPWP, IsRequired, Digits, LengthMin15Max16),
			validation.Field(&model.CompanyName, IsRequired, AlphaSpace, LengthMax50),
			validation.Field(&model.CompanyField, IsRequired, AlphaSpace, LengthMax20),
			validation.Field(&model.CompanyType, IsRequired, AlphaSpace, LengthMax10),
			validation.Field(&model.CompanyAddress, IsRequired, ComplexCharacters, LengthMin10Max100),
			validation.Field(&model.CompanyPhone, IsRequired, Digits, LengthMin8Max20),
			validation.Field(&model.Province, IsRequired, AlphaSpace, LengthMax20),
			validation.Field(&model.City, IsRequired, AlphaSpace, LengthMax20),
			validation.Field(&model.District, IsRequired, AlphaSpace, LengthMax20),
		)
	}

	return nil
}

// Generate 16 bytes randomly and securely using the
// Cryptographically secure pseudorandom number generator (CSPRNG)
// in the crypto.rand package
func GenerateRandomSalt(saltSize int) ([]byte, error) {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		return nil, err
	}

	return salt, nil
}

// Combine password and salt then hash them using the SHA-512
// hashing algorithm and then return the hashed password
// as a base64 encoded string
func HashPassword(password string, salt []byte) string {
	var (
		passwordBytes             []byte    = append([]byte(password), salt...)
		sha512Hasher              hash.Hash = sha512.New()
		hashedPasswordBytes       []byte
		base64EncodedPasswordHash string
	)

	// Write password bytes to the hasher
	sha512Hasher.Write(passwordBytes)

	// Get the SHA-512 hashed password
	hashedPasswordBytes = sha512Hasher.Sum(nil)

	// Convert the hashed password to a base64 encoded string
	base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)

	return base64EncodedPasswordHash
}

// Check if two passwords match
func VerifyPassword(hashedPassword, currPassword string, salt []byte) bool {
	return hashedPassword == HashPassword(currPassword, salt)
}
