package parser

import (
	"fmt"
	"text/scanner"

	"github.com/IvanZagoskin/wkt/geometry"
	"github.com/IvanZagoskin/wkt/text"
)

func (p *Parser) parseMultiPoint(ct geometry.CoordinateType) (*geometry.MultiPoint, error) {
	switch ct {
	case geometry.XY, geometry.XYM, geometry.XYZ, geometry.XYZM:
		multiPoint := &geometry.MultiPoint{Type: ct}
		for {
			point, err := p.parsePoint(ct)
			if err != nil {
				return nil, fmt.Errorf("parseLineString: %w", err)
			}
			multiPoint.Points = append(multiPoint.Points, point)

			if p.scanner.Scan() == scanner.EOF {
				return nil, ErrUnexpectedEOF
			}

			switch text.Token(p.scanner.TokenText()) {
			case text.ClosingParenthesis:
				return multiPoint, nil
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
