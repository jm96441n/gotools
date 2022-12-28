package match

import (
	"bufio"
	"errors"
	"io"
	"os"
)

type Matcher struct {
	input io.Reader
}

type option func(*Matcher) error

func NewMatcher(opts ...option) (Matcher, error) {
	l := Matcher{
		input: os.Stdin,
	}
	for _, opt := range opts {
		err := opt(&l)
		if err != nil {
			return Matcher{}, err
		}
	}
	return l, nil
}

func WithInput(input io.Reader) option {
	return func(m *Matcher) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		m.input = input
		return nil
	}
}

func (m Matcher) MatchLines(needle string) []string {
	scanner := bufio.NewScanner(m.input)
	matchingLines := make([]string, 0)
	for scanner.Scan() {
		haystack := scanner.Text()
		if Search(haystack, needle) {
			matchingLines = append(matchingLines, haystack)
		}
	}
	return matchingLines
}

func Search(haystack, needle string) bool {
	lps := make([]int, len(needle))
	prefix := 0
	for i := 1; i < len(needle); i++ {
		for prefix > 0 && needle[i] != needle[prefix] {
			prefix = lps[prefix-1]
		}
		if needle[i] == needle[prefix] {
			prefix += 1
			lps[i] = prefix
		}
	}
	prefixIdx := 0
	for idx := range haystack {
		for prefixIdx > 0 && haystack[idx] != needle[prefixIdx] {
			prefixIdx = lps[prefixIdx-1]
		}
		if haystack[idx] == needle[prefixIdx] {
			if prefixIdx == len(needle)-1 {
				return true
			}
			prefixIdx += 1
		}
	}
	return false
}
