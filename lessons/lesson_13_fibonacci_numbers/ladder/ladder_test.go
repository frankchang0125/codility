package ladder

import (
    "reflect"
    "testing"
)

func TestSolution(t *testing.T) {
	var ans, expectAns []int

	ans = Solution([]int{4, 4, 5, 5, 1}, []int{3, 2, 4, 3, 1})
	expectAns = []int{5, 1, 8, 0, 1}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
}
