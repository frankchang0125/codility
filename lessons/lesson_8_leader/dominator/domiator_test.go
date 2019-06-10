package dominator

import (
	"testing"
)

func TestSolution1(t *testing.T) {
	var ans, expectAns int

	ans = Solution1([]int{3, 4, 3, 2, 3, -1, 3, 3})
	expectAns = 7
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
}

func TestSolution2(t *testing.T) {
	var ans, expectAns int

	ans = Solution2([]int{3, 4, 3, 2, 3, -1, 3, 3})
	expectAns = 7
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
}
