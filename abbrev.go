package abbrev

import (
	"errors"
	"strings"
)

func MatchString(pattern string, str string) bool {
	lis := strings.Split(pattern, "/")
	if len(lis) != 2 {
		return false
	}
	if len(str) > len(pattern) - 1 {
		return false
	}
	if !strings.HasPrefix(str, lis[0]) {
		return false
	}
	for i, s := range str[len(lis[0]):] {
		if rune(lis[1][i]) != s {
			return false
		}
	}
	return true
}

type Abbrev struct {
	length int
	pre    string
	follow string
}

func Compile(pattern string) (*Abbrev, error) {
	lis := strings.Split(pattern, "/")
	if len(lis) != 2 {
		return nil, errors.New("syntax error")
	}
	return &Abbrev{len(pattern)-1, lis[0], lis[1]}, nil
}

func MustCompile(pattern string) *Abbrev {
	abb, err := Compile(pattern)
	if err != nil {
		panic("abbrev: Compile(" + pattern + "):" + err.Error())
	}
	return abb
}

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

func (abb *Abbrev) All() string {
	return abb.pre + abb.follow
}
