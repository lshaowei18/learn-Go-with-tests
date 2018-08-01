package reflection

import (
	"reflect"
	"testing"
)

type Test struct {
	Name     string
	Input    interface{}
	Expected []string
}

func TestWalk(t *testing.T) {
	tests := []Test{
		{
			Name:     "Struct with one field",
			Input:    struct{ string }{"chris"},
			Expected: []string{"chris"},
		},
		{
			Name:     "Struct with two field",
			Input:    struct{ a, b string }{"chris", "rock"},
			Expected: []string{"chris", "rock"},
		},
		{
			Name: "Struct with two different field types",
			Input: struct {
				a int
				b string
			}{1, "rock"},
			Expected: []string{"rock"},
		},
		{
			Name: "Pointer struct",
			Input: &struct {
				a int
				b string
			}{1, "rock"},
			Expected: []string{"rock"},
		},
		{
			Name: "Slices",
			Input: []struct {
				a int
				b string
			}{
				{1, "rock"},
				{2, "jack"},
			},
			Expected: []string{"rock", "jack"},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(test.Expected, got) {
				t.Errorf("got : %v, wants %v.", got, test.Expected)
			}
		})
	}
}

func BenchmarkWalk(b *testing.B) {
	input := Test{
		Name: "Slices",
		Input: []struct {
			a int
			b string
		}{
			{1, "rock"},
			{2, "jack"},
		},
		Expected: []string{"rock", "jack"},
	}

	var got []string

	for i := 0; i < b.N; i++ {
		walk(input, func(input string) {
			got = append(got, input)
		})
	}
}

func BenchmarkWalk1(b *testing.B) {
	input := Test{
		Name: "Slices",
		Input: []struct {
			a int
			b string
		}{
			{1, "rock"},
			{2, "jack"},
		},
		Expected: []string{"rock", "jack"},
	}

	var got []string

	for i := 0; i < b.N; i++ {
		Walk1(input, func(input string) {
			got = append(got, input)
		})
	}
}
