package consts

const (
	// GENERAL
	RCNotFound            string = "404"
	RCBadRequest          string = "400"
	RCUnauthorized        string = "403"
	RCForbidden           string = "403"
	RCMethodNotAllowed    string = "405"
	RCInternalServerError string = "500"

	// SPECIFIC
	// status 200
	RCSuccess      string = "000"
	RCUpdated      string = "002"
	RCUserNotFound string = "008"

	// status 201
	RCCreated string = "001"
	RCDeleted string = "003"

	// status 400
	RCInvalidRequestBody string = "005"
	RCDuplicateEntry     string = "006"

	// status 401
	RCCredentialNotMatch string = "007"
	RCInvalidToken       string = "010"
	RCExpiredToken       string = "011"
	RCBadParam           string = "099"
)
