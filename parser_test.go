package wkt_test

import (
	"bytes"
	"errors"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/IvanZagoskin/wkt"
	"github.com/IvanZagoskin/wkt/geometry"
)

func TestWktParser_Point(t *testing.T) {
	testCases := []struct {
		Name     string
		Wkt      []byte
		Expected *geometry.Point
		Error    error
	}{
		{
			Name:     "Simple point",
			Wkt:      []byte("POINT (30 20)"),
			Expected: &geometry.Point{X: 30, Y: 20, Type: geometry.XY},
		},
		{
			Name:     "Point with float coordinates",
			Wkt:      []byte("POINT (30.2 20.7)"),
			Expected: &geometry.Point{X: 30.2, Y: 20.7, Type: geometry.XY},
		},
		{
			Name:     "Point Z",
			Wkt:      []byte("POINT Z (30.2 20.7 34.777)"),
			Expected: &geometry.Point{X: 30.2, Y: 20.7, Z: 34.777, Type: geometry.XYZ},
		},
		{
			Name:     "Point M",
			Wkt:      []byte("POINT M (30.2 20.7 34.777)"),
			Expected: &geometry.Point{X: 30.2, Y: 20.7, M: 34.777, Type: geometry.XYM},
		},
		{
			Name:     "Point ZM",
			Wkt:      []byte("POINT ZM (30.2 20.7 34.777 63.23)"),
			Expected: &geometry.Point{X: 30.2, Y: 20.7, Z: 34.777, M: 63.23, Type: geometry.XYZM},
		},
		{
			Name:  "Bad point",
			Wkt:   []byte("POINT (30.2 20.7 34.777 63.23 63.23)"),
			Error: wkt.ErrUnexpectedToken,
		},
		{
			Name:  "Bad wkt(1)",
			Wkt:   []byte("POIN (30.2 20.7 34.777 63.23 63.23)"),
			Error: wkt.ErrUnexpectedGeometryType,
		},
		{
			Name:  "Bad wkt(2)",
			Wkt:   []byte("POINT 30.2 20.7)"),
			Error: wkt.ErrUnexpectedCoordinateType,
		},
		{
			Name:  "Bad wkt(3)",
			Wkt:   []byte("POINt (30.2 20.7 34.777 63.23 63.23)"),
			Error: wkt.ErrUnexpectedGeometryType,
		},
		{
			Name:  "Bad wkt(4)",
			Wkt:   []byte("POINT ()"),
			Error: strconv.ErrSyntax,
		},
		{
			Name:  "Bad wkt(6)",
			Wkt:   []byte("POINT (23)"),
			Error: strconv.ErrSyntax,
		},
	}

	wktParser := wkt.NewParser()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			geom, err := wktParser.ParseWKT(bytes.NewReader(tc.Wkt))
			if tc.Error != nil {
				if !errors.Is(err, tc.Error) {
					t.Fatalf("\ngot: %v\nexpected error: %s\n", err, tc.Error)
				}
				return
			}

			if err != nil {
				t.Fatalf("\nunexpected error:%v\n\n", err)
				return
			}

			point := geom.(*geometry.Point)
			if diff := cmp.Diff(point, tc.Expected); diff != "" {
				t.Fatal("\n-want +got\n", diff)
			}
		})
	}
}

func TestWktParser_LineString(t *testing.T) {
	testCases := []struct {
		Name     string
		Wkt      []byte
		Expected *geometry.LineString
		Error    error
	}{
		{
			Name: "Simple lineString",
			Wkt:  []byte("LINESTRING (30 10, 10 30, 40 40)"),
			Expected: &geometry.LineString{
				Points: []*geometry.Point{
					{X: 30, Y: 10, Type: geometry.XY},
					{X: 10, Y: 30, Type: geometry.XY},
					{X: 40, Y: 40, Type: geometry.XY},
				},
				Type: geometry.XY,
			},
		},
		{
			Name: "LineString with float",
			Wkt:  []byte("LINESTRING (30.123 10.15, 10.66 30.23, 40.23 40.66)"),
			Expected: &geometry.LineString{
				Points: []*geometry.Point{
					{X: 30.123, Y: 10.15, Type: geometry.XY},
					{X: 10.66, Y: 30.23, Type: geometry.XY},
					{X: 40.23, Y: 40.66, Type: geometry.XY},
				},
				Type: geometry.XY,
			},
		},
		{
			Name: "LineString Z",
			Wkt:  []byte("LINESTRING Z(30.123 10.15 11.22, 10.66 30.23 22.33, 40.23 40.66 44.44)"),
			Expected: &geometry.LineString{
				Points: []*geometry.Point{
					{X: 30.123, Y: 10.15, Z: 11.22, Type: geometry.XYZ},
					{X: 10.66, Y: 30.23, Z: 22.33, Type: geometry.XYZ},
					{X: 40.23, Y: 40.66, Z: 44.44, Type: geometry.XYZ},
				},
				Type: geometry.XYZ,
			},
		},
		{
			Name: "LineString M",
			Wkt:  []byte("LINESTRING M(30.123 10.15 11.22, 10.66 30.23 22.33, 40.23 40.66 44.44)"),
			Expected: &geometry.LineString{
				Points: []*geometry.Point{
					{X: 30.123, Y: 10.15, M: 11.22, Type: geometry.XYM},
					{X: 10.66, Y: 30.23, M: 22.33, Type: geometry.XYM},
					{X: 40.23, Y: 40.66, M: 44.44, Type: geometry.XYM},
				},
				Type: geometry.XYM,
			},
		},
		{
			Name: "LineString ZM",
			Wkt:  []byte("LINESTRING ZM(30.123 10.15 11.22 55.55, 10.66 30.23 22.33 66.66, 40.23 40.66 44.44 77.77)"),
			Expected: &geometry.LineString{
				Points: []*geometry.Point{
					{X: 30.123, Y: 10.15, Z: 11.22, M: 55.55, Type: geometry.XYZM},
					{X: 10.66, Y: 30.23, Z: 22.33, M: 66.66, Type: geometry.XYZM},
					{X: 40.23, Y: 40.66, Z: 44.44, M: 77.77, Type: geometry.XYZM},
				},
				Type: geometry.XYZM,
			},
		},
	}

	wktParser := wkt.NewParser()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			geom, err := wktParser.ParseWKT(bytes.NewReader(tc.Wkt))
			if tc.Error != nil {
				if !errors.Is(err, tc.Error) {
					t.Fatalf("\ngot: %v\nexpected error: %s\n", err, tc.Error)
				}
				return
			}

			if err != nil {
				t.Fatalf("\nunexpected error:%v\n\n", err)
				return
			}

			point := geom.(*geometry.LineString)
			if diff := cmp.Diff(point, tc.Expected); diff != "" {
				t.Fatal("\n-want +got\n", diff)
			}
		})
	}
}

func TestWktParser_CircularString(t *testing.T) {
	testCases := []struct {
		Name     string
		Wkt      []byte
		Expected *geometry.CircularString
		Error    error
	}{
		{
			Name: "Simple CircularString",
			Wkt:  []byte("CIRCULARSTRING (30 10, 10 30, 40 40)"),
			Expected: &geometry.CircularString{
				Points: []*geometry.Point{
					{X: 30, Y: 10, Type: geometry.XY},
					{X: 10, Y: 30, Type: geometry.XY},
					{X: 40, Y: 40, Type: geometry.XY},
				},
				Type: geometry.XY,
			},
		},
		{
			Name: "CircularString with float",
			Wkt:  []byte("CIRCULARSTRING (30.123 10.15, 10.66 30.23, 40.23 40.66)"),
			Expected: &geometry.CircularString{
				Points: []*geometry.Point{
					{X: 30.123, Y: 10.15, Type: geometry.XY},
					{X: 10.66, Y: 30.23, Type: geometry.XY},
					{X: 40.23, Y: 40.66, Type: geometry.XY},
				},
				Type: geometry.XY,
			},
		},
		{
			Name: "CircularString Z",
			Wkt:  []byte("CIRCULARSTRING Z(30.123 10.15 11.22, 10.66 30.23 22.33, 40.23 40.66 44.44)"),
			Expected: &geometry.CircularString{
				Points: []*geometry.Point{
					{X: 30.123, Y: 10.15, Z: 11.22, Type: geometry.XYZ},
					{X: 10.66, Y: 30.23, Z: 22.33, Type: geometry.XYZ},
					{X: 40.23, Y: 40.66, Z: 44.44, Type: geometry.XYZ},
				},
				Type: geometry.XYZ,
			},
		},
		{
			Name: "CircularString M",
			Wkt:  []byte("CIRCULARSTRING M(30.123 10.15 11.22, 10.66 30.23 22.33, 40.23 40.66 44.44)"),
			Expected: &geometry.CircularString{
				Points: []*geometry.Point{
					{X: 30.123, Y: 10.15, M: 11.22, Type: geometry.XYM},
					{X: 10.66, Y: 30.23, M: 22.33, Type: geometry.XYM},
					{X: 40.23, Y: 40.66, M: 44.44, Type: geometry.XYM},
				},
				Type: geometry.XYM,
			},
		},
		{
			Name: "CircularString ZM",
			Wkt:  []byte("CIRCULARSTRING ZM(30.123 10.15 11.22 55.55, 10.66 30.23 22.33 66.66, 40.23 40.66 44.44 77.77)"),
			Expected: &geometry.CircularString{
				Points: []*geometry.Point{
					{X: 30.123, Y: 10.15, Z: 11.22, M: 55.55, Type: geometry.XYZM},
					{X: 10.66, Y: 30.23, Z: 22.33, M: 66.66, Type: geometry.XYZM},
					{X: 40.23, Y: 40.66, Z: 44.44, M: 77.77, Type: geometry.XYZM},
				},
				Type: geometry.XYZM,
			},
		},
	}

	wktParser := wkt.NewParser()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			geom, err := wktParser.ParseWKT(bytes.NewReader(tc.Wkt))
			if tc.Error != nil {
				if !errors.Is(err, tc.Error) {
					t.Fatalf("\ngot: %v\nexpected error: %s\n", err, tc.Error)
				}
				return
			}

			if err != nil {
				t.Fatalf("\nunexpected error:%v\n\n", err)
				return
			}

			point := geom.(*geometry.CircularString)
			if diff := cmp.Diff(point, tc.Expected); diff != "" {
				t.Fatal("\n-want +got\n", diff)
			}
		})
	}
}

func TestWktParser_Polygon(t *testing.T) {
	testCases := []struct {
		Name     string
		Wkt      []byte
		Expected *geometry.Polygon
		Error    error
	}{
		{
			Name: "Simple polygon",
			Wkt:  []byte("POLYGON ((30 10, 40 40, 20 40, 10 20, 30 10))"),
			Expected: &geometry.Polygon{
				Type: geometry.XY,
				LineStrings: []*geometry.LineString{
					{
						Type: geometry.XY,
						Points: []*geometry.Point{
							{X: 30, Y: 10, Type: geometry.XY},
							{X: 40, Y: 40, Type: geometry.XY},
							{X: 20, Y: 40, Type: geometry.XY},
							{X: 10, Y: 20, Type: geometry.XY},
							{X: 30, Y: 10, Type: geometry.XY},
						},
					},
				},
			},
		},
		{
			Name: "Polygon Z",
			Wkt:  []byte("POLYGON Z((30 10 10, 40 40 20, 20 40 30, 10 20 40, 30 10 50))"),
			Expected: &geometry.Polygon{
				Type: geometry.XYZ,
				LineStrings: []*geometry.LineString{
					{
						Type: geometry.XYZ,
						Points: []*geometry.Point{
							{X: 30, Y: 10, Z: 10, Type: geometry.XYZ},
							{X: 40, Y: 40, Z: 20, Type: geometry.XYZ},
							{X: 20, Y: 40, Z: 30, Type: geometry.XYZ},
							{X: 10, Y: 20, Z: 40, Type: geometry.XYZ},
							{X: 30, Y: 10, Z: 50, Type: geometry.XYZ},
						},
					},
				},
			},
		},
		{
			Name: "Polygon M",
			Wkt:  []byte("POLYGON M((30 10 10, 40 40 20, 20 40 30, 10 20 40, 30 10 50))"),
			Expected: &geometry.Polygon{
				Type: geometry.XYM,
				LineStrings: []*geometry.LineString{
					{
						Type: geometry.XYM,
						Points: []*geometry.Point{
							{X: 30, Y: 10, M: 10, Type: geometry.XYM},
							{X: 40, Y: 40, M: 20, Type: geometry.XYM},
							{X: 20, Y: 40, M: 30, Type: geometry.XYM},
							{X: 10, Y: 20, M: 40, Type: geometry.XYM},
							{X: 30, Y: 10, M: 50, Type: geometry.XYM},
						},
					},
				},
			},
		},
		{
			Name: "Polygon ZM",
			Wkt:  []byte("POLYGON ZM((30 10 10 -10.10, 40 40 20 -20.20, 20 40 30 -30.30, 10 20 40 -40.40, 30 10 50 -50.50))"),
			Expected: &geometry.Polygon{
				Type: geometry.XYZM,
				LineStrings: []*geometry.LineString{
					{
						Type: geometry.XYZM,
						Points: []*geometry.Point{
							{X: 30, Y: 10, Z: 10, M: -10.10, Type: geometry.XYZM},
							{X: 40, Y: 40, Z: 20, M: -20.20, Type: geometry.XYZM},
							{X: 20, Y: 40, Z: 30, M: -30.30, Type: geometry.XYZM},
							{X: 10, Y: 20, Z: 40, M: -40.40, Type: geometry.XYZM},
							{X: 30, Y: 10, Z: 50, M: -50.50, Type: geometry.XYZM},
						},
					},
				},
			},
		},
	}

	wktParser := wkt.NewParser()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			geom, err := wktParser.ParseWKT(bytes.NewReader(tc.Wkt))
			if tc.Error != nil {
				if !errors.Is(err, tc.Error) {
					t.Fatalf("\ngot: %v\nexpected error: %s\n", err, tc.Error)
				}
				return
			}

			if err != nil {
				t.Fatalf("\nunexpected error:%v\n\n", err)
				return
			}

			point := geom.(*geometry.Polygon)
			if diff := cmp.Diff(point, tc.Expected); diff != "" {
				t.Fatal("\n-want +got\n", diff)
			}
		})
	}
}
