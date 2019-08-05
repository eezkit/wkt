package wkt

import (
	"fmt"
	"text/scanner"

	"github.com/IvanZagoskin/wkt/geometry"
	"github.com/IvanZagoskin/wkt/text"
)

func (p *Parser) parsePolygon(ct geometry.CoordinateType) (*geometry.Polygon, error) {
	switch ct {
	case geometry.XY, geometry.XYM, geometry.XYZ, geometry.XYZM:
		// skip first text.OpeningParenthesis, because parseLineString is not waiting it
		if err := p.skipTokenAndCheck(text.OpeningParenthesis); err != nil {
			return nil, fmt.Errorf("skipTokenAndCheck: %w", err)
		}

		polygon := &geometry.Polygon{Type: ct, LineStrings: []*geometry.LineString{}}
		for {
			lineString, err := p.parseLineString(ct)
			if err != nil {
				return nil, fmt.Errorf("parseLineString: %w", err)
			}
			polygon.LineStrings = append(polygon.LineStrings, lineString)

			if p.scanner.Scan() == scanner.EOF {
				return nil, ErrUnexpectedEOF
			}

			switch text.Token(p.scanner.TokenText()) {
			case text.ClosingParenthesis:
				return polygon, nil
			case text.Comma:
				continue
			}
		}

	default:
		return nil, fmt.Errorf("%w: %scanner", ErrUnexpectedToken, p.scanner.TokenText())
	}
}
