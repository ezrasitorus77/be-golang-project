package consts

const (
	AlphaSpaceRegex             = "^[a-zA-Z\\.]?[a-zA-Z\\.\\s]+$"
	AlphaSpaceValidationMessage = "Field must only contains alphabets"

	AlphaUpperRegex             = "[[:upper:]]"
	AlphaUpperValidationMessage = "Field must contain uppercase letter"

	AlphaLowerRegex             = "[[:lower:]]"
	AlphaLowerValidationMessage = "Field must contain lowercase letter"

	ComplexCharactersRegex             = "^[a-zA-Z]?[a-zA-Z0-9\\-\\.\\,\\(\\)\\/\\s\\+]+$"
	ComplexCharactersValidationMessage = "Field can only contain alpahnumeric and (-.,()/\\+) symbols"

	DigitOnlyRegex             = "[[:digit:]]"
	DigitOnlyValidationMessage = "Field must contain digit"

	SymbolsOnlyRegex             = "[~!@#$%^&*()_\\-+=|\\}\\]{\\[\"\\':;?/>.<,]"
	SymbolsOnlyValidationMessage = "Field must contain symbol"

	ExceptSpaceRegex                  = "^[^\\s]+$"
	ExceptSpaceRegexValidationMessage = "Field must not contain space"

	SuccessRC      = "000"
	SuccessMessage = "OK"

	InvalidRequestBodyRC      = "005"
	InvalidRequestBodyMessage = "Invalid request body"

	GeneralForbiddenRC      = "403"
	GeneralForbiddenMessage = "User is not authorized"
	InvalidTokenRC          = "010"
	ExpiredTokenRC          = "011"

	TokenExpiredForbiddenRC = "112"

	CreatedRC      = "001"
	CreatedMessage = "Data successfully created"

	GeneralInternalServerErrorRC      = "099"
	GeneralInternalServerErrorMessage = "Internal server error"

	MethodNotAllowedRC      = "405"
	MethodNotAllowedMessage = "Method not allowed"

	ContextKey = "Y9tn1zr9gV"

	RequiredValidationMessage = "Field is required"
	EmailValidationMessage    = "Email format does not match"

	Length16ValidationMessage          = "Field must contain exactly 16 characters"
	LengthMin15Max16ValidationMessage  = "Field must contain 15 - 16 characters"
	LengthMin8Max20ValidationMessage   = "Must contain minimal 8 and maximum 20 characters"
	LengthMax50ValidationMessage       = "Maximum characters is 50"
	LengthMax20ValidationMessage       = "Maximum characters is 20"
	LengthMax10ValidationMessage       = "Maximum characters is 10"
	LengthMin10Max100ValidationMessage = "Field must contain 10 - 100 characters"

	MinSecretKeySize = 32.

	ErrInvalidToken = "Token is invalid"
	ErrExpiredToken = "Token has expired"

	JWTSecretKey = "6V$SY=HR:}1\\y{bU2/<VG)q*R<\"&:{a?"

	SaltSize = 16

	UserNotFoundRC      = "002"
	UserNotFoundMessage = "User is not found"

	CredentialNotMatchRC      = "003"
	CredentialNotMatchMessage = "Username and password do not match"
)
