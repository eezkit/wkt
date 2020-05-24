package parser

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/scanner"

	"github.com/IvanZagoskin/wkt/geometry"
	"github.com/IvanZagoskin/wkt/text"
)

var (
	ErrUnexpectedToken          = errors.New("unexpected token")
	ErrUnexpectedEOF            = errors.New("unexpected EOF")
	ErrUnexpectedGeometryType   = errors.New("unexpected geometry type")
	ErrUnexpectedCoordinateType = errors.New("unexpected coordinate type")
)

// Parser implements parsing wkt
type Parser struct {
	scanner *scanner.Scanner
}

// New returns Parser
func New() *Parser {
	return &Parser{scanner: &scanner.Scanner{}}
}

// ParseWKT detects a geometry object and returns it
func (p *Parser) ParseWKT(r io.Reader) (geometry.Geometry, error) {
	p.scanner.Init(r)

	gt, err := p.detectGeomType()
	if err != nil {
		return nil, fmt.Errorf("detect geometry type: %w", err)
	}

	switch gt {
	case geometry.PointGT:
		ct, err := p.detectCoordType()
		if err != nil {
			return nil, fmt.Errorf("detect coordinate type: %w", err)
		}

		if ct == geometry.Empty {
			return &geometry.Point{Type: geometry.Empty}, nil
		}

		point, err := p.parsePoint(ct)
		if err != nil {
			return nil, fmt.Errorf("parse point: %w", err)
		}

		if err := p.skipTokenAndCheck(text.ClosingParenthesis); err != nil {
			return nil, fmt.Errorf("skip token and check: %w", err)
		}

		return point, nil

	case geometry.MultyPointGT:
		ct, err := p.detectCoordType()
		if err != nil {
			return nil, fmt.Errorf("detect coordinate type: %w", err)
		}

		if ct == geometry.Empty {
			return &geometry.Point{Type: geometry.Empty}, nil
		}

		multiPoint, err := p.parseMultiPoint(ct)
		if err != nil {
			return nil, fmt.Errorf("parse point: %w", err)
		}

		return multiPoint, nil

	case geometry.LineStringGT:
		ct, err := p.detectCoordType()
		if err != nil {
			return nil, fmt.Errorf("detect coordinate type: %w", err)
		}

		if ct == geometry.Empty {
			return &geometry.Point{Type: geometry.Empty}, nil
		}

		lineString, err := p.parseLineString(ct)
		if err != nil {
			return nil, fmt.Errorf("parse linestring: %w", err)
		}

		return lineString, nil

	case geometry.CircularStringGT:
		ct, err := p.detectCoordType()
		if err != nil {
			return nil, fmt.Errorf("detect coordinate type: %w", err)
		}

		if ct == geometry.Empty {
			return &geometry.Point{Type: geometry.Empty}, nil
		}

		circularString, err := p.parseCircularString(ct)
		if err != nil {
			return nil, fmt.Errorf("parse linestring: %w", err)
		}

		return circularString, nil

	case geometry.MultiLineStringGT:
		ct, err := p.detectCoordType()
		if err != nil {
			return nil, fmt.Errorf("detect coordinate type: %w", err)
		}

		if ct == geometry.Empty {
			return &geometry.Point{Type: geometry.Empty}, nil
		}

		multiLineString, err := p.parseMultiLineString(ct)
		if err != nil {
			return nil, fmt.Errorf("parse linestring: %w", err)
		}

		return multiLineString, nil

	case geometry.PolygonGT:
		ct, err := p.detectCoordType()
		if err != nil {
			return nil, fmt.Errorf("detect coordinate type: %w", err)
		}

		if ct == geometry.Empty {
			return &geometry.Point{Type: geometry.Empty}, nil
		}

		polygon, err := p.parsePolygon(ct)
		if err != nil {
			return nil, fmt.Errorf("parse polygon: %w", err)
		}

		return polygon, nil

	default:
		return nil, fmt.Errorf("%w: %s", ErrUnexpectedGeometryType, p.scanner.TokenText())
	}
}

func (p *Parser) detectGeomType() (geometry.Type, error) {
	if p.scanner.Scan() == scanner.EOF {
		return geometry.UndefinedGT, ErrUnexpectedEOF
	}

	switch text.Token(p.scanner.TokenText()) {
	case text.POINT:
		return geometry.PointGT, nil

	case text.LINESTRING:
		return geometry.LineStringGT, nil

	case text.CIRCULARSTRING:
		return geometry.CircularStringGT, nil

	case text.MULTILINESTRING:
		return geometry.MultiLineStringGT, nil

	case text.POLYGON:
		return geometry.PolygonGT, nil

	case text.MULTIPOINT:
		return geometry.MultyPointGT, nil

	default:
		return geometry.UndefinedGT, fmt.Errorf("%w: %s", ErrUnexpectedGeometryType, p.scanner.TokenText())
	}
}

func (p *Parser) detectCoordType() (geometry.CoordinateType, error) {
	if p.scanner.Scan() == scanner.EOF {
		return geometry.Undefined, ErrUnexpectedEOF
	}

	switch text.Token(p.scanner.TokenText()) {
	case text.ZCoordinates:
		if err := p.skipTokenAndCheck(text.OpeningParenthesis); err != nil {
			return geometry.Undefined, fmt.Errorf("skip token and check: %w", err)
		}
		return geometry.XYZ, nil

	case text.MCoordinates:
		if err := p.skipTokenAndCheck(text.OpeningParenthesis); err != nil {
			return geometry.Undefined, fmt.Errorf("skip token and check: %w", err)
		}
		return geometry.XYM, nil

	case text.ZMCoordinates:
		if err := p.skipTokenAndCheck(text.OpeningParenthesis); err != nil {
			return geometry.Undefined, fmt.Errorf("skip token and check: %w", err)
		}
		return geometry.XYZM, nil

	case text.OpeningParenthesis:
		return geometry.XY, nil

	case text.Empty:
		return geometry.Empty, nil

	default:
		return geometry.Undefined, fmt.Errorf("%w: %s", ErrUnexpectedCoordinateType, p.scanner.TokenText())
	}
}

// skipTokenAndCheck skips next token and checks that skipped token equal specified token
func (p *Parser) skipTokenAndCheck(token text.Token) error {
	if p.scanner.Scan() == scanner.EOF {
		return ErrUnexpectedEOF
	}

	if text.Token(p.scanner.TokenText()) != token {
		return fmt.Errorf("%w: %s", ErrUnexpectedToken, p.scanner.TokenText())
	}
	return nil
}

// countCoordinatesBy returns count expected coordinates for current coordinate type
func countCoordinatesBy(ct geometry.CoordinateType) geometry.NumberOfCoordinates {
	switch ct {
	case geometry.XY:
		return geometry.NumXY
	case geometry.XYM:
		return geometry.NumXYM
	case geometry.XYZ:
		return geometry.NumXYZ
	case geometry.XYZM:
		return geometry.NumXYZM
	default:
		return geometry.NumUndefined
	}
}

// parsePointCoords parse coordinates and returns slice with them.
//
// Must be called only if you sure that next tokens are coordinates.
func parsePointCoords(s *scanner.Scanner, ct geometry.CoordinateType) ([]float64, error) {
	countCoordinates := countCoordinatesBy(ct)
	coordinates := make([]float64, 0, countCoordinates)
	for tok := s.Scan(); ; tok = s.Scan() {
		if tok == scanner.EOF {
			return nil, ErrUnexpectedEOF
		}

		isNegative := false
		if strings.EqualFold(s.TokenText(), string(text.Minus)) {
			isNegative = true
			if s.Scan() == scanner.EOF {
				return nil, ErrUnexpectedEOF
			}
		}

		c, err := strconv.ParseFloat(s.TokenText(), 64)
		if err != nil {
			return nil, err
		}

		if isNegative {
			c = -c
		}

		coordinates = append(coordinates, c)
		if len(coordinates) == int(countCoordinates) {
			return coordinates, nil
		}
	}
}
