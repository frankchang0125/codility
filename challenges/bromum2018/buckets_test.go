package buckets

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var ans, expectAns int
	
	ans = Solution(3, 2, []int{1, 2, 0, 1, 1, 0, 0, 1}, []int{0, 3, 0, 2, 0, 3, 0, 0})
	expectAns = 4
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(2, 2, []int{0, 1}, []int{5, 5})
	expectAns = -1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(2, 2, []int{0, 1, 0, 1, 0, 1}, []int{1, 3, 0, 0, 3, 3})
	expectAns = 5
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
}
