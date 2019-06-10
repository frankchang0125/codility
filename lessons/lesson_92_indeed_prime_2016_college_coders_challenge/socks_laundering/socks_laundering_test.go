package sockslaundering

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var ans, expectAns int

	ans = Solution(2, []int{1, 2, 1, 1}, []int{1, 4, 3, 2, 4})
	expectAns = 3
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(2, []int{}, []int{1, 4, 3, 2, 4})
	expectAns = 1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(5, []int{}, []int{2, 2, 2, 2, 2})
	expectAns = 2
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(1, []int{}, []int{2, 2, 2, 2, 2})
	expectAns = 0
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(6, []int{}, []int{2, 2, 2, 2, 2, 2})
	expectAns = 3
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(2, []int{1, 2, 3, 4}, []int{})
	expectAns = 0
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(2, []int{1, 2, 3, 4}, []int{1, 2, 3, 4})
	expectAns = 2
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(4, []int{1, 2, 3, 4}, []int{1, 2, 3, 4})
	expectAns = 4
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(0, []int{1, 2, 3, 4}, []int{1, 2, 3, 4})
	expectAns = 0
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(4, []int{}, []int{})
	expectAns = 0
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(5, []int{2, 3, 4}, []int{2, 2, 2, 2, 2})
	expectAns = 3
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
}
