package doer_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/totalsnack/FaceOnTable"
	"testing"
)

func TestDo(t *testing.T) {
	testCases := map[string]struct {
		s   string
		i   int
		b   bool
		res string
		err error
	}{
		"testing a, 1, lower":      {s: "a", i: 1, b: false, res: "[1a]", err: nil},
		"testing b, 2, upper":      {s: "b", i: 2, b: true, res: "[2B]", err: nil},
		"testing d, 21, upper":     {s: "d", i: 21, b: true, res: "D", err: nil},
		"testing d, 13, lower":     {s: "d", i: 13, b: false, res: "d", err: nil},
		"testing z, 1, lower":      {s: "z", i: 1, b: false, res: "", err: errors.New("invalid s")},
		"testing d, 100500, lower": {s: "d", i: 100500, b: false, res: "", err: errors.New("invalid i")}, // i != s
	}
	for testName, test := range testCases {
		t.Run(testName, func(t *testing.T) {
			got, err := doer.Do(test.s, test.i, test.b)
			if err != nil {
				assert.ErrorContains(t, err, test.err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.res, got)
		})
	}
}
