package nesting

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var ans, expectAns int

	ans = Solution("(()(())())")
	expectAns = 1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }

    ans = Solution("())")
	expectAns = 0
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
}
