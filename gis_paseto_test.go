package gis_paseto

import (
	"fmt"
	"testing"
)

func TestGCHandlerFunc(t *testing.T) {
	data := GCHandlerFunc("string", "GIS", "geogis")

	fmt.Printf("%+v", data)
}

