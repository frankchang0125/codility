package oddoccurrencesinarray

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var ans, expectAns int

	ans = Solution([]int{9, 3, 9, 3, 9, 7, 9})
	expectAns = 7
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }

    ans = Solution([]int{9})
	expectAns = 9
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }

    ans = Solution([]int{9, 9, 9, 9, 9, 9, 5})
	expectAns = 5
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }

    ans = Solution([]int{9, 9, 9, 9, 9, 9, 9})
	expectAns = 9
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
}
