package mushroompicker

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var ans, expectAns int

	ans = Solution([]int{2}, 0, 0)
	expectAns = 2
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{2}, 0, 6)
	expectAns = 2
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{2, 3}, 0, 0)
	expectAns = 2
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{2, 3}, 0, 1)
	expectAns = 5
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{2, 3}, 1, 0)
	expectAns = 3
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{2, 3}, 1, 6)
	expectAns = 5
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{2, 3, 7, 5, 1, 3, 9}, 4, 6)
	expectAns = 25
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{2, 3, 7, 5, 1, 3, 9}, 4, 0)
	expectAns = 1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{2, 3, 7, 5, 1, 3, 9}, 0, 6)
	expectAns = 30
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{2, 3, 7, 5, 1, 3, 9}, 0, 0)
	expectAns = 2
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{2, 3, 7, 5, 1, 3, 9}, 0, 10)
	expectAns = 30
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{2, 3, 7, 5, 1, 3, 9}, 6, 6)
	expectAns = 30
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{2, 3, 7, 5, 1, 3, 9}, 6, 0)
	expectAns = 9
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{2, 3, 7, 5, 1, 3, 9}, 6, 10)
	expectAns = 30
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{2, 3, 7, 5, 1, 3, 9}, 1, 2)
	expectAns = 15
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{2, 3, 7, 5, 1, 3, 9}, 1, 3)
	expectAns = 16
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{2, 3, 7, 5, 1, 3, 9}, 5, 2)
	expectAns = 12
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{2, 3, 7, 5, 1, 3, 9}, 5, 3)
	expectAns = 16
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{2, 3, 5, 2, 1, 3, 9}, 5, 3)
	expectAns = 13
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
}
