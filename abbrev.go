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

func (abb *Abbrev) String() string {
	return abb.pre + abb.follow
}

func MatchString(pattern string, str string) (bool, error) {
	abb, err := Compile(pattern)
	if err != nil {
		return false, err
	}
	return abb.MatchString(str), nil
}

func For(pattern string, str string) bool {
	abb, err := Compile(pattern)
	if err != nil {
		panic("abbrev: Compile(" + pattern + "):" + err.Error())
	}
	return abb.MatchString(str)
}

