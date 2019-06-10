package distinct

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var ans, expectAns int

	ans = Solution([]int{2, 1, 1, 2, 3, 1})
	expectAns = 3
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{0, 1, 2, 3, 4})
	expectAns = 5
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{})
	expectAns = 0
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{0})
	expectAns = 1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{1})
	expectAns = 1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{0, 1})
	expectAns = 2
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{-1, 1})
	expectAns = 2
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
}
