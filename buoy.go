package gobuoy

import (
	"net/url"
	"strings"
)

type Purl struct {
	PathMatcher       Matcher
	QueryStringValues url.Values
}

func New(path Matcher) *Purl {
	return &Purl{
		PathMatcher: path,
	}
}

type Matcher interface {
	Match([]byte, []byte) bool
	GetMatches() map[string]string
}

func (p *Purl) MatchPath(url, match string) bool {
	index := strings.IndexByte(url, '?')
	if index == -1 {
		index = len(url)
	}
	path := url[0:index]
	return p.PathMatcher.Match([]byte(path), []byte(match))
}

func (p *Purl) GetPathMatches() map[string]string {
	return p.PathMatcher.GetMatches()
}

func (p *Purl) MatchQueryString(u string) bool {
	iindex := strings.IndexByte(u, '?')
	if iindex == -1 {
		return true
	}
	v, err := url.ParseQuery(u[iindex+1:])

	if err != nil {
		return false
	}

	p.QueryStringValues = v
	return true
}

func (p *Purl) GetQueryStringMatches() url.Values {
	return p.QueryStringValues
}

func (p *Purl) Match(url, pattern string) bool {
	return p.MatchPath(url, pattern) && p.MatchQueryString(url)
}

func NewCurlyMatcher() *Purl {
	return New(
		NewKeyMatcher("{}"),
	)
}
