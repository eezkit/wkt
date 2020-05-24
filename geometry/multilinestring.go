package geometry

// MULTILINESTRING
type MultiLineString struct {
	Lines []*LineString
	Type  CoordinateType
}

func (m MultiLineString) GetGeometryType() Type {
	return MultiLineStringGT
}
