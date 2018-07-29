package reflection

import "testing"

type Test struct {
	Name string
}

func TestWalk(t *testing.T) {
	t.Run("1 string value", func(t *testing.T) {
		expected := "christ"
		var got []string
		x := Test{expected}
		walk(x, func(input string) {
			got = append(got, input)
		})
		if len(got) != 1 {
			t.Errorf("got : %d, wants 1.", len(got))
		}
	})
}
