package parser

import (
	"fmt"
	"text/scanner"

	"github.com/IvanZagoskin/wkt/geometry"
	"github.com/IvanZagoskin/wkt/text"
)

func (p *Parser) parseCircularString(ct geometry.CoordinateType) (*geometry.CircularString, error) {
	switch ct {
	case geometry.XY, geometry.XYZ, geometry.XYM, geometry.XYZM:
		circularString := &geometry.CircularString{Type: ct}
		for {
			point, err := p.parsePoint(ct)
			if err != nil {
				return nil, fmt.Errorf("parsePointCoords: %w", err)
			}
			circularString.Points = append(circularString.Points, point)

			if p.scanner.Scan() == scanner.EOF {
				return nil, ErrUnexpectedEOF
			}

			switch text.Token(p.scanner.TokenText()) {
			case text.ClosingParenthesis:
				return circularString, nil
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
