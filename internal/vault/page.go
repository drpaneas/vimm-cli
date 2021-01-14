package vault

import (
	"fmt"
	"log"
	"net/url"
)

type page struct {
	link *url.URL
}

func ExampleURL_EscapedPath() {
	u, err := url.Parse("http://example.com/x/y%2Fz")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Path:", u.Path)
	fmt.Println("RawPath:", u.RawPath)
	fmt.Println("EscapedPath:", u.EscapedPath())
	// Output:
	// Path: /x/y/z
	// RawPath: /x/y%2Fz
	// EscapedPath: /x/y%2Fz
}
