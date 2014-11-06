package abbrev

import (
	"errors"
	"strings"
)

type Abbrev struct {
	length int
	pre    string
	follow string
}

// Compile parses an abbreviation and returns, if successful,
// an Abbrev object that can be used to match against test.
func Compile(pattern string) (*Abbrev, error) {
	lis := strings.Split(pattern, "/")
	if len(lis) != 2 {
		return nil, errors.New("syntax error")
	}
	return &Abbrev{len(pattern)-1, lis[0], lis[1]}, nil
}

// MustCompile is like Compile but panics if the expression cannot be parsed.
func MustCompile(pattern string) *Abbrev {
	abb, err := Compile(pattern)
	if err != nil {
		panic("abbrev: Compile(" + pattern + "):" + err.Error())
	}
	return abb
}

// MatchString reports whether the Abbrev matches the string str.
func (abb *Abbrev) MatchString(str string) bool {
	if len(str) > abb.length {
		return false
	}
	if !strings.HasPrefix(str, abb.pre) {
		return false
	}
	for i, s := range str[len(abb.pre):] {
		if rune(abb.follow[i]) != s {
			return false
		}
	}
	return true
}

// String returns the source text used to compile the abbreviation.
func (abb *Abbrev) String() string {
	return strings.Join([]string{abb.pre, abb.follow}, "/")
}

// Longest returns the longest string which matches the abbreviation.
func (abb *Abbrev) Longest() string {
	return strings.Join([]string{abb.pre, abb.follow}, "")
}

// Shortest returns the shortest string which matches the abbreviation.
func (abb *Abbrev) Shortest() string {
	return abb.pre
}

// MatchString checks whether a textual abbreviation matches the string.
func MatchString(pattern string, str string) (bool, error) {
	abb, err := Compile(pattern)
	if err != nil {
		return false, err
	}
	return abb.MatchString(str), nil
}

// For is like MatchString but panics if the expression cannot be parsed.
func For(pattern string, str string) bool {
	abb, err := Compile(pattern)
	if err != nil {
		panic("abbrev: Compile(" + pattern + "):" + err.Error())
	}
	return abb.MatchString(str)
}

