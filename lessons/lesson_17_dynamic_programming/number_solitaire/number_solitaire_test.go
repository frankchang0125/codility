package numbersolitaire

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var ans, expectAns int

	ans = Solution([]int{1, -2, 0, 9, -1, -2})
	expectAns = 8
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }

    ans = Solution([]int{1})
	expectAns = 1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }

    ans = Solution([]int{-5})
	expectAns = -5
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }

    ans = Solution([]int{-1, -2, -3, -4, -5})
	expectAns = -6
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }

    ans = Solution([]int{-1, -2, -3, -4, -5, -6, -7, -8})
	expectAns = -11
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }

    ans = Solution([]int{-1, -2, -3, 4, -5, -6, -7, -8})
	expectAns = -5
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }

    ans = Solution([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	expectAns = 55
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }

    ans = Solution([]int{1, 2, 3, 4, 5, 6, -7, 8, 9, 10})
	expectAns = 48
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
}
