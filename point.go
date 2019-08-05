package wkt

import (
	"fmt"

	"github.com/IvanZagoskin/wkt/geometry"
)

func (p *Parser) parsePoint(ct geometry.CoordinateType) (*geometry.Point, error) {
	switch ct {
	case geometry.XY:
		coords, err := parsePointCoords(p.scanner, geometry.XY)
		if err != nil {
			return nil, fmt.Errorf("parsePointCoords: %w", err)
		}

		return &geometry.Point{Type: geometry.XY, X: coords[0], Y: coords[1]}, nil

	case geometry.XYM:
		coords, err := parsePointCoords(p.scanner, geometry.XYM)
		if err != nil {
			return nil, fmt.Errorf("parsePointCoords: %w", err)
		}

		return &geometry.Point{Type: geometry.XYM, X: coords[0], Y: coords[1], M: coords[2]}, nil

	case geometry.XYZ:
		coords, err := parsePointCoords(p.scanner, geometry.XYZ)
		if err != nil {
			return nil, fmt.Errorf("parsePointCoords: %w", err)
		}

		return &geometry.Point{Type: geometry.XYZ, X: coords[0], Y: coords[1], Z: coords[2]}, nil

	case geometry.XYZM:
		coords, err := parsePointCoords(p.scanner, geometry.XYZM)
		if err != nil {
			return nil, fmt.Errorf("parsePointCoords: %w", err)
		}

		return &geometry.Point{Type: geometry.XYZM, X: coords[0], Y: coords[1], Z: coords[2], M: coords[3]}, nil

	default:
		return nil, fmt.Errorf("%w: %scanner", ErrUnexpectedToken, p.scanner.TokenText())
	}
}
