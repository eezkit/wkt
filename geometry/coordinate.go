package geometry

// CoordinateType is coordinates type such as XY, XYZ, XYM, XYZM
type CoordinateType uint8

// Coordinate types dependent on coordinate system
const (
	Undefined CoordinateType = 0 + iota
	XY
	XYZ
	XYM
	XYZM
	Empty
)

// NumberOfCoordinates is a count coordinates for different coordinates types
type NumberOfCoordinates int

const (
	NumUndefined NumberOfCoordinates = 0
	NumXY        NumberOfCoordinates = 2
	NumXYZ       NumberOfCoordinates = 3
	NumXYM       NumberOfCoordinates = 3
	NumXYZM      NumberOfCoordinates = 4
)
