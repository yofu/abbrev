package abbrev

import (
	"errors"
	"strings"
)

type Abbrev struct {
	expr   string
	size   int
	length []int
	pre    []string
	follow []string
}

// Compile parses an abbreviation and returns, if successful,
// an Abbrev object that can be used to match against test.
func Compile(pattern string) (*Abbrev, error) {
	lis := strings.Split(pattern, "/")
	if len(lis)%2 != 0 {
		return nil, errors.New("syntax error")
	}
	size := int(len(lis) / 2)
	length := make([]int, size)
	pre := make([]string, size)
	follow := make([]string, size)
	for i := 0; i < size; i++ {
		pre[i] = lis[2*i]
		follow[i] = lis[2*i+1]
		length[i] = len(pre[i]) + len(follow[i])
	}
	return &Abbrev{pattern, size, length, pre, follow}, nil
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
	pos := 0
match:
	for i := 0; i < abb.size; i++ {
		if len(str[pos:]) > abb.length[i] {
			return false
		}
		if !strings.HasPrefix(str[pos:], abb.pre[i]) {
			return false
		}
		pos += len(abb.pre[i])
		for j, s := range str[pos:] {
			if rune(abb.follow[i][j]) != s {
				if i == abb.size-1 {
					return false
				} else {
					continue match
				}
			}
			pos++
		}
	}
	return true
}

// String returns the source text used to compile the abbreviation.
func (abb *Abbrev) String() string {
	return abb.expr
}

// Longest returns the longest string which matches the abbreviation.
func (abb *Abbrev) Longest() string {
	var buf []byte
	for i := 0; i < abb.size; i++ {
		buf = append(buf, abb.pre[i]...)
		buf = append(buf, abb.follow[i]...)
	}
	return string(buf)
}

// Shortest returns the shortest string which matches the abbreviation.
func (abb *Abbrev) Shortest() string {
	return strings.Join(abb.pre, "")
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
