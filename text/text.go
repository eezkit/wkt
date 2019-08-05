package text

// Token is a type for text tokens
type Token string

const (
	POINT Token = "POINT"

	LINESTRING     Token = "LINESTRING"
	CIRCULARSTRING Token = "CIRCULARSTRING"

	POLYGON Token = "POLYGON"

	OpeningParenthesis Token = "("
	ClosingParenthesis Token = ")"
	Comma              Token = ","
	Minus              Token = "-"

	ZCoordinates  Token = "Z"
	MCoordinates  Token = "M"
	ZMCoordinates Token = "ZM"
	Empty         Token = "EMPTY"
)
