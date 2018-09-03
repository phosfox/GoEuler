package main

import "testing"

func TestIsAbundand(t *testing.T) {
	abundandNums := []int{12, 18, 20, 24, 30, 40, 42, 48, 54, 56, 60, 66, 70, 72, 78, 80, 84, 88, 90, 96, 100}
	nonAbundandNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 15, 17, 99, 89, 17, 19, 29, 39}
	for _, n := range abundandNums {
		if !IsAbundand(n) {
			t.Errorf("Expected true %d. got %t", n, IsAbundand(n))
		}
	}

	for _, n := range nonAbundandNums {
		if IsAbundand(n) {
			t.Errorf("Exptected false %d, got %t", n, IsAbundand(n))
		}
	}
}
