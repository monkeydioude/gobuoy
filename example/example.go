package main

import (
	"github.com/monkeydioude/gobuoy"
	"fmt"
)

func yeahhhhhBuoyyyy() {
	// Short for gobuoy.NewKeyMatcher("{}")
	yb := gobuoy.NewCurlyMatcher()

	// Let's say this URL path hits your API...
	urlPath := "/url-example/buoyyyy"

	// ... and you, an intellect, want to make a route that matches this kind of URL path.
	// This pattern defines the route and identify important parts of the path
	// you want to work with by giving them a name.
	pattern := "/url-example/{yeahhhhhh}"

	// Match returns a bool, should be verified here.
	// Do not reproduce this at home.
	yb.Match(urlPath, pattern)

	// If Match() is successful, matches will contain
	// every matched named parts of the path.
	matches := yb.GetPathMatches()

	// > "buoyyyy"
	fmt.Println(matches["yeahhhhhh"])
}

func main() {
	yeahhhhhBuoyyyy()
}
