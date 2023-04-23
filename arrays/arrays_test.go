package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("array of 5 elements", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("expected %d got %d, for array %v", got, want, numbers)
		}
	})
}
