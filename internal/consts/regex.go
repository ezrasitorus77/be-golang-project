package consts

const (
	AlphaSpaceRegex        = "^[a-zA-Z\\.]?[a-zA-Z\\.\\s]+$"
	AlphaUpperRegex        = "[[:upper:]]"
	AlphaLowerRegex        = "[[:lower:]]"
	ComplexCharactersRegex = "^[a-zA-Z]?[a-zA-Z0-9\\-\\.\\,\\(\\)\\/\\s\\+]+$"
	DigitOnlyRegex         = "[[:digit:]]"
	ExceptSpaceRegex       = "^[^\\s]+$"
	SymbolsOnlyRegex       = "[~!@#$%^&*()_\\-+=|\\}\\]{\\[\"\\':;?/>.<,]"
)
