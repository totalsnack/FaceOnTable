package doer_test

import (
	"fmt"
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
				"testing a, 1, lower":  {args: args{"a", 1, false}, res: "[a1]", err: nil},
				"testing b, 2, upper":  {args: args{"b", 2, true}, res: "[B2]", err: nil},
				"testing d, 21, upper": {args: args{"d", 21, true}, res: "D", err: nil},
				"testing d, 13, lower": {args: args{"d", 13, false}, res: "[d2]", err: nil},
				"testing z, 1, lower":  {args: args{"z", 1, false}, res: "", err: fmt.Errorf("invalid s")},
				"testing f, 0, lower":  {args: args{"f", 0, false}, res: "", err: fmt.Errorf("invalid i")},
			}
			for testName, test := range testCases {
				t.Run(testName, func(t *testing.T) {
					if got, err := doer.Do(test.s, test.i, test.b); got != test.res && err != test.err {
						fmt.Println("!!!!!!!!!!!!!!!", got, err)
						t.Fail()
					}
				})
			}
		}
	})
}
