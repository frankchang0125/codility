package numberofdiscintersections

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var ans, expectAns int

	ans = Solution([]int{1, 5, 2, 1, 4, 0})
	expectAns = 11
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
	
	ans = Solution([]int{1, 1, 1})
	expectAns = 3
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
	
	ans = Solution([]int{})
	expectAns = 0
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
	
	ans = Solution([]int{10})
	expectAns = 0
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
	
	ans = Solution([]int{1, 1})
	expectAns = 1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
	
	extremeLarge := make([]int, 100000)
	for i := 0; i < len(extremeLarge); i++ {
		extremeLarge[i] = 10000000
	}

	ans = Solution(extremeLarge)
	expectAns = -1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
}