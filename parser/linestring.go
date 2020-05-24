package parser

import (
	"fmt"
	"text/scanner"

	"github.com/IvanZagoskin/wkt/geometry"
	"github.com/IvanZagoskin/wkt/text"
)

func (p *Parser) parseLineString(ct geometry.CoordinateType) (*geometry.LineString, error) {
	switch ct {
	case geometry.XY, geometry.XYM, geometry.XYZ, geometry.XYZM:
		lineString := &geometry.LineString{Type: ct}
		for {
			point, err := p.parsePoint(ct)
			if err != nil {
				return nil, fmt.Errorf("parsePointCoords: %w", err)
			}
			lineString.Points = append(lineString.Points, point)

			if p.scanner.Scan() == scanner.EOF {
				return nil, ErrUnexpectedEOF
			}

			switch text.Token(p.scanner.TokenText()) {
			case text.ClosingParenthesis:
				return lineString, nil
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
