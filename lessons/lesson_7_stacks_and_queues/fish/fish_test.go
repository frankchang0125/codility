package fish

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var ans, expectAns int

	ans = Solution([]int{4, 3, 2, 1, 5}, []int{0, 1, 0, 0, 0})
	expectAns = 2
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
}
