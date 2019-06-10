package missinginteger

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var ans, expectAns int

	ans = Solution([]int{1, 3, 6, 4, 1, 2})
	expectAns = 5
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{1, 2, 3})
	expectAns = 4
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{-1, -3})
	expectAns = 1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{0})
	expectAns = 1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{-10, 0})
	expectAns = 1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
}
