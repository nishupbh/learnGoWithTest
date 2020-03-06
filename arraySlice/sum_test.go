package arraySlice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("test case using array", func(t *testing.T) {
		numbers := []int{1, 2, 1, 42, 5}
		got := sum(numbers)
		want := 51
		if got != want {
			t.Errorf("expected %d got %d for given number %v", want, got, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {

	got := sumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}


func TestSumAllTails( t *testing.T){
	checkSums := func(t *testing.T,got,want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}	
	}

	t.Run("make the sum of all tails", func(t *testing.T){
		got := sumAllTrails([]int{1, 2}, []int{0, 9})
        want := []int{2, 9}
		checkSums(t, got, want)
		
	})

	t.Run("make the sum of all tails", func(t *testing.T){
		got := sumAllTrails([]int{}, []int{0, 9})
        want := []int{0, 9}
		checkSums(t, got, want)
		
	})
}