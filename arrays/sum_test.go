package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	got := Sum([]int{1, 2, 3, 4})
	expected := 10

	if got != expected {
		t.Errorf("got %d; expected %d", got, expected)
	}
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{3, 9})
	expected := []int{3, 12}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v; expected %v", got, expected)
	}
}

func TestSumAll2(t *testing.T) {

	got := SumAll2([]int{1, 2}, []int{3, 9})
	expected := []int{3, 12}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v; expected %v", got, expected)
	}
}

//Sum all numbers in a slice except for the first
func TestSumTails(t *testing.T) {

	got := SumTails([]int{1, 2}, []int{3, 9})
	expected := []int{2, 9}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v; expected %v", got, expected)
	}
}

func TestSumTails2(t *testing.T) {

	got := SumTails2([]int{1, 2}, []int{3, 9})
	expected := []int{2, 9}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v; expected %v", got, expected)
	}
}

func BenchmarkSumAll1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 45})
	}
}

func BenchmarkSumAll2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll2([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 45})
	}
}

func BenchmarkSumTail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumTails([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 45})
	}
}

func BenchmarkSumTail2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumTails2([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 45})
	}
}
