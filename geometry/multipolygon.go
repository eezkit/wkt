package geometry

type MultiPolygon struct {
	Polygons []*Polygon
	Type     CoordinateType
}

func (m *MultiPolygon) GetGeometryType() Type {
	return MultiPolygonGT
}
