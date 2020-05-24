package parser

import (
	"fmt"
	"text/scanner"

	"github.com/IvanZagoskin/wkt/geometry"
	"github.com/IvanZagoskin/wkt/text"
)

func (p *Parser) parseMultiPolygon(ct geometry.CoordinateType) (*geometry.MultiPolygon, error) {
	switch ct {
	case geometry.XY, geometry.XYM, geometry.XYZ, geometry.XYZM:
		multyPolygon := &geometry.MultiPolygon{Type: ct}
		for {
			// skip first text.OpeningParenthesis, because parseLineString is not waiting it
			if err := p.skipTokenAndCheck(text.OpeningParenthesis); err != nil {
				return nil, fmt.Errorf("skipTokenAndCheck: %w", err)
			}

			polygon, err := p.parsePolygon(ct)
			if err != nil {
				return nil, fmt.Errorf("parseLineString: %w", err)
			}
			multyPolygon.Polygons = append(multyPolygon.Polygons, polygon)

			if p.scanner.Scan() == scanner.EOF {
				return nil, ErrUnexpectedEOF
			}

			switch text.Token(p.scanner.TokenText()) {
			case text.ClosingParenthesis:
				return multyPolygon, nil
			case text.Comma:
				continue
			default:
				return nil, fmt.Errorf("%w: %s", ErrUnexpectedToken, p.scanner.TokenText())
			}
		}

	default:
		return nil, fmt.Errorf("%w: %d", ErrUnexpectedCoordinateType, ct)
	}
}
