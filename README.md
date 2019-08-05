# WKT parser

WKT parser is a library for parsing wkt geometry into simple structures, which is designed to be converted to representations of other types of geometry (for example [S2](https://github.com/golang/geo)).

The basis for parsing is taken bnf, which is found here: http://svn.osgeo.org/postgis/trunk/doc/bnf-wkt.txt

## Install

```bash
go get -u github.com/IvanZagoskin/wkt-parser
```

## Example
You can see examples of use in tests.