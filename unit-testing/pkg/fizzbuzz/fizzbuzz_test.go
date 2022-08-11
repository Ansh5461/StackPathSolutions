package fizzbuzz

import (
	"fmt"
	"testing"
)

func stringcompare(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			fmt.Println("Error here")
			return false
		}
	}
	return true
}

func TestFizzBuzz(t *testing.T) {
	got := FizzBuzz(16, 3, 5)
	want := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16"}

	if !stringcompare(got, want) {
		t.Errorf("Wanted is not equal to got, want %v   \ngot %v", want, got)
	}
}
