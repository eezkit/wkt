package geometry

// Point is wkt point representation
type Point struct {
	X, Y, Z, M float64
	Type       CoordinateType
}

// GetGeometryType returns geometry type
func (p *Point) GetGeometryType() Type {
	return PointGT
}
