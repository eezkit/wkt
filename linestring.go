package wkt

import (
	"fmt"
	"text/scanner"

	"github.com/IvanZagoskin/wkt/geometry"
	"github.com/IvanZagoskin/wkt/text"
)

func (p *Parser) parseLineString(ct geometry.CoordinateType) (*geometry.LineString, error) {
	switch ct {
	case geometry.XY:
		lineString := &geometry.LineString{Type: ct, Points: []*geometry.Point{}}

		for {
			coords, err := parsePointCoords(p.scanner, geometry.XY)
			if err != nil {
				return nil, fmt.Errorf("parsePointCoords: %w", err)
			}
			lineString.Points = append(lineString.Points, &geometry.Point{
				Type: geometry.XY, X: coords[0], Y: coords[1],
			})

			if p.scanner.Scan() == scanner.EOF {
				return nil, ErrUnexpectedEOF
			}

			switch text.Token(p.scanner.TokenText()) {
			case text.ClosingParenthesis:
				return lineString, nil
			case text.Comma:
				continue
			}
		}

	case geometry.XYM:
		lineString := &geometry.LineString{Type: ct, Points: []*geometry.Point{}}

		for {
			coords, err := parsePointCoords(p.scanner, geometry.XYM)
			if err != nil {
				return nil, fmt.Errorf("parsePointCoords: %w", err)
			}
			lineString.Points = append(lineString.Points, &geometry.Point{
				Type: geometry.XYM, X: coords[0], Y: coords[1], M: coords[2],
			})

			if p.scanner.Scan() == scanner.EOF {
				return nil, ErrUnexpectedEOF
			}

			switch text.Token(p.scanner.TokenText()) {
			case text.ClosingParenthesis:
				return lineString, nil
			case text.Comma:
				continue
			}
		}

	case geometry.XYZ:
		lineString := &geometry.LineString{Type: ct, Points: []*geometry.Point{}}

		for {
			coords, err := parsePointCoords(p.scanner, geometry.XYZ)
			if err != nil {
				return nil, fmt.Errorf("parsePointCoords: %w", err)
			}
			lineString.Points = append(lineString.Points, &geometry.Point{
				Type: geometry.XYZ, X: coords[0], Y: coords[1], Z: coords[2],
			})

			if p.scanner.Scan() == scanner.EOF {
				return nil, ErrUnexpectedEOF
			}

			switch text.Token(p.scanner.TokenText()) {
			case text.ClosingParenthesis:
				return lineString, nil
			case text.Comma:
				continue
			}
		}

	case geometry.XYZM:
		lineString := &geometry.LineString{Type: ct, Points: []*geometry.Point{}}

		for {
			coords, err := parsePointCoords(p.scanner, geometry.XYZM)
			if err != nil {
				return nil, fmt.Errorf("parsePointCoords: %w", err)
			}
			lineString.Points = append(lineString.Points, &geometry.Point{
				Type: geometry.XYZM, X: coords[0], Y: coords[1], Z: coords[2], M: coords[3],
			})

			if p.scanner.Scan() == scanner.EOF {
				return nil, ErrUnexpectedEOF
			}

			switch text.Token(p.scanner.TokenText()) {
			case text.ClosingParenthesis:
				return lineString, nil
			case text.Comma:
				continue
			}
		}

	default:
		return nil, fmt.Errorf("%w: %scanner", ErrUnexpectedToken, p.scanner.TokenText())
	}
}
