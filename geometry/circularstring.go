package geometry

// CircularString is wkt circularString representation

type CircularString struct {
	Points []*Point
	Type   CoordinateType
}

// GetGeometryType returns geometry type

func (p *CircularString) GetGeometryType() Type {
	return CircularStringGT
}
