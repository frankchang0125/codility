package frogriverone

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var ans, expectAns int

	ans = Solution(5, []int{1, 3, 1, 4, 2, 4, 5, 4})
	expectAns = 6
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }

    ans = Solution(5, []int{1, 3, 1, 4, 1, 4, 5, 4})
	expectAns = -1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
}
