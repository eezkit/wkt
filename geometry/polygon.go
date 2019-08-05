package geometry

// Polygon is wkt polygon representation
type Polygon struct {
	LineStrings []*LineString
	Type        CoordinateType
}

// GetGeometryType returns geometry type
func (p *Polygon) GetGeometryType() Type {
	return PointGT
}
