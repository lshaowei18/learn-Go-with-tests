package integers

import "testing"

//Test struct
type Test struct {
	in1, in2, out int
}

//Test variables
var tests = []Test{
	{2, 2, 4},
}

func TestAdder(t *testing.T) {
	for i, test := range tests {
		sum := Add(test.in1, test.in2)
		if sum != test.out {
			t.Errorf("#%d : Add(%d, %d)= %d; want %d", i, test.in1, test.in2, sum, test.out)
		}
	}
}
