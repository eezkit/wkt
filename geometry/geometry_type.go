package geometry

// Type is geometry type such as Point, LineString, Polygon etc.
type Type uint8

const (
	UndefinedGT Type = 0 + iota
	PointGT

	LineStringGT
	CircularStringGT

	PolygonGT
)
