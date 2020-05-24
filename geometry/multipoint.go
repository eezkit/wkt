package geometry

// MultiPoint is wkt multiPoint representation
type MultiPoint struct {
	Points []*Point
	Type   CoordinateType
}

// GetGeometryType returns geometry type
func (m *MultiPoint) GetGeometryType() Type {
	return MultyPointGT
}
