package gobuoy

import (
	"testing"
)

func TestKeyMatcherCanMatchSimpleString(t *testing.T) {
	trial := "/image/{id}"
	url := "/image/04713a0d49f030a254bdd2d61742bf1ee1546360281"
	p := NewKeyMatcher("{}")

	if !p.Match([]byte(url), []byte(trial)) {
		t.Fail()
	}
}
func TestKeyMatcherCanMatch(t *testing.T) {
	trial := "/test/{wesh}/3/{salut}/{nada}//{}}{/{}/}{/}}}}}{/}}}{{{{{{"
	url := "/test/alors/3/ccool//2/3/4"
	p := NewKeyMatcher("{}")

	p.Match([]byte(url), []byte(trial))
	m := p.GetMatches()

	if len(m) != 3 {
		t.Fail()
	}
	if v, ok := m["wesh"]; !ok || v != "alors" {
		t.Fail()
	}
	if v, ok := m["salut"]; !ok || v != "ccool" {
		t.Fail()
	}
	if v, ok := m["nada"]; !ok || v != "" {
		t.Fail()
	}
}

func TestKeyMatcherFailOnNonMatchingString(t *testing.T) {
	trial := "/b/a/c/d/"
	url := "/a/b/c/d/"
	p := NewKeyMatcher("{}")

	if p.Match([]byte(trial), []byte(url)) {
		t.Fail()
	}
}

func TestResetMatchesOnConsecutivesMatches(t *testing.T) {
	trial1 := []byte("/test/{oh}")
	url1 := []byte("/test/no")
	trial2 := []byte("/test/{ah}")
	url2 := []byte("/test/ok")

	p := NewKeyMatcher("{}")
	p.Match(url1, trial1)
	p.Match(url2, trial2)

	if p.GetMatches()["oh"] == "no" || p.GetMatches()["ah"] != "ok" {
		t.Fail()
	}
	
}
