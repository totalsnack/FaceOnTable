package doer_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/totalsnack/FaceOnTable"
	"testing"
)

func TestDo(t *testing.T) {
	type args struct {
		s string
		i int
		b bool
	}

	t.Run("Test for all coverage", func(t *testing.T) {
		{
			testCases := map[string]struct {
				args
				res string
				err error
			}{
				"testing a, 1, lower":      {args: args{"a", 1, false}, res: "[1a]", err: nil},
				"testing b, 2, upper":      {args: args{"b", 2, true}, res: "[2B]", err: nil},
				"testing d, 21, upper":     {args: args{"d", 21, true}, res: "D", err: nil},
				"testing d, 13, lower":     {args: args{"d", 13, false}, res: "d", err: nil},
				"testing z, 1, lower":      {args: args{"z", 1, false}, res: "", err: errors.New("invalid s")},
				"testing d, 100500, lower": {args: args{"d", 100500, false}, res: "", err: errors.New("invalid i")}, // i != s
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
	})
}
