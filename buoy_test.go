package gobuoy

import (
	"testing"
)

func TestICanParseFullUrl(t *testing.T) {
	p := New(NewKeyMatcher("{}"))

	if !p.Match("/pouet/?roi=dadidou", "/{pouet}/") ||
		p.GetPathMatches()["pouet"] != "pouet" ||
		p.GetQueryStringMatches().Get("roi") != "dadidou" {
		t.Fail()
	}
}

func TestICanParseUrlWithoutQueryString(t *testing.T) {
	p := NewCurlyMatcher()

	if !p.Match("/thatimfeelin/", "/{isthislove}/") ||
		p.GetPathMatches()["isthislove"] != "thatimfeelin" {
		t.Fail()
	}
}
