package main_test

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	type args struct {
		s string
		i int
		b bool
	}
	type want struct {
		res string
		err error
	}
	testCases := map[string]struct {
		args args
		res  string
		err  error
	}{
		"testing a, 1, lower": {
			args{"a", 1, false},
			"[a1]",
			nil,
		},
		"testing b, 2, upper": {
			args: args{"", 2, true},
			res:  "[B2]",
			err:  nil,
		},
		"testing d, 21, upper": {
			args: args{"d", 21, true},
			res:  "D",
			err:  nil,
		},
		"testing d, 13, lower": {
			args: args{"d", 13, false},
			res:  "[B2]",
			err:  nil,
		},
		"testing z, 1, lower": {
			args: args{"z", 1, false},
			res:  "",
			err:  fmt.Errorf("invalid s"),
		},
		"testing f, 0, lower": {
			args: args{"f", 0, false},
			res:  "",
			err:  fmt.Errorf("invalid s"),
		},
	}
}
