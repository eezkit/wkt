# WKT parser

WKT parser is a library for parsing wkt geometry into simple structures, which is designed to be converted to representations of other types of geometry (for example [S2](https://github.com/golang/geo)).

## Install

```bash
go get -u github.com/IvanZagoskin/wkt
```

## Example
```go
package main

import (
	"bytes"
	"fmt"

	"github.com/IvanZagoskin/wkt/parser"
	"github.com/IvanZagoskin/wkt/geometry"
)

func main() {
	p := parser.New()
	g, _ := p.ParseWKT(bytes.NewReader([]byte("POINT (30 20)")))
	switch geom := g.(type) {
	case *geometry.Point:
		fmt.Printf("%+v", geom)
	}
}

```
You can see more usage examples in tests.

## Supported geometry

Added support for basic geometry types:
* POINT
* LINESTRING
* POLYGON
* MULTIPOINT
* MULTILINESTRING
* MULTIPOLYGON
* CIRCULARSTRING

