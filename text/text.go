package text

// Token is a type for text tokens
type Token string

const (
	POINT      Token = "POINT"
	MULTIPOINT Token = "MULTIPOINT"

	LINESTRING      Token = "LINESTRING"
	CIRCULARSTRING  Token = "CIRCULARSTRING"
	MULTILINESTRING Token = "MULTILINESTRING"

	POLYGON      Token = "POLYGON"
	MULTIPOLYGON Token = "MULTIPOLYGON"

	OpeningParenthesis Token = "("
	ClosingParenthesis Token = ")"
	Comma              Token = ","
	Minus              Token = "-"

	ZCoordinates  Token = "Z"
	MCoordinates  Token = "M"
	ZMCoordinates Token = "ZM"
	Empty         Token = "EMPTY"
)
