package consts

const (
	// status 200
	SuccessMessage      string = "OK"
	UserNotFoundMessage string = "User is not found"

	// status 201
	CreatedMessage string = "Data successfully created"
	DeletedMessage string = "Data successfully deleted"
	UpdatedMessage string = "Data successfully updated"

	// status 400
	AlphaSpaceValidationMessage        string = "Field must only contains alphabets"
	AlphaUpperValidationMessage        string = "Field must contain uppercase letter"
	AlphaLowerValidationMessage        string = "Field must contain lowercase letter"
	BadRequestMessage                  string = "Bad request"
	ComplexCharactersValidationMessage string = "Field can only contain alpahnumeric and (-.,()/\\+) symbols"
	DigitOnlyValidationMessage         string = "Field must contain digit"
	DuplicateEntryMessage              string = "Duplicate entry found"
	EmailValidationMessage             string = "Email format does not match"
	ExceptSpaceRegexValidationMessage  string = "Field must not contain space"
	InvalidRequestBodyParamMessage     string = "Invalid request body / param"
	Length16ValidationMessage          string = "Field must contain exactly 16 characters"
	LengthMin15Max16ValidationMessage  string = "Field must contain 15 - 16 characters"
	LengthMin8Max20ValidationMessage   string = "Must contain minimal 8 and maximum 20 characters"
	LengthMax50ValidationMessage       string = "Maximum characters is 50"
	LengthMax20ValidationMessage       string = "Maximum characters is 20"
	LengthMax10ValidationMessage       string = "Maximum characters is 10"
	LengthMin10Max100ValidationMessage string = "Field must contain 10 - 100 characters"
	LengthMin20Max500ValidationMessage string = "Field must contain 20 - 500 characters"

	RequiredValidationMessage    string = "Field is required"
	SymbolsOnlyValidationMessage string = "Field must contain symbol"

	// status 401
	CredentialNotMatchMessage string = "Username and password do not match"
	InvalidToken              string = "Token is invalid"
	ExpiredToken              string = "Token has expired"
	UserUnauthorizedMessage   string = "User is not authorized"

	// status 404
	NotFoundMessage     string = "Not found"
	PageNotFoundMessage string = "Page you are looking for is not found"

	// status 405
	NotAllowedMessage       string = "Method not allowed"
	MethodNotAllowedMessage string = "Requested path doesn't accept the method"

	// status 500
	InternalServerErrorMessage string = "Internal server error"
	ServerErrorMessage         string = "Error occurred in server, please try again later"
)
