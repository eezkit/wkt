package geometry

// LineString is wkt lineString representation
type LineString struct {
	Points []*Point
	Type   CoordinateType
}

// GetGeometryType returns geometry type
func (p *LineString) GetGeometryType() Type {
	return LineStringGT
}
