package match_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/jm96441n/gotools/linecount/match"
)

func TestMatcherReturnsLinesContainingMatchingString(t *testing.T) {
	fakeInput := bytes.NewBuffer([]byte{})
	m, err := match.NewMatcher(match.WithInput(io.Reader(fakeInput)))
	if err != nil {
		t.Error(err)
	}
	fakeInput.Write([]byte("hello world\nthis should not match\nthis one will match hello\n"))

	got := m.MatchLines("hello")
	want := []string{"hello world", "this one will match hello"}
	if !eqSlices(got, want) {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func eqSlices[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
