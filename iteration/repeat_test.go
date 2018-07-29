package iteration

import "testing"

func TestRepeat(t *testing.T) {

	got := Repeat("a")
	expected := "aaaaa"

	if got != expected {
		t.Errorf("got %s; expected %s", got, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func TestExampleRepeat(t *testing.T) {
	got := ExampleRepeat("a", 3)
	expected := "aaa"

	if got != expected {
		t.Errorf("got %s; expected %s", got, expected)
	}
}

func BenchmarkExampleRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExampleRepeat("a", 100)
	}
}

func TestStringsRepeat(t *testing.T) {
	got := StringsRepeat("a", 4)
	expected := "aaaa"

	if got != expected {
		t.Errorf("got %s; expected %s", got, expected)
	}
}

func BenchmarkStringsRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringsRepeat("a", 100)
	}
}
