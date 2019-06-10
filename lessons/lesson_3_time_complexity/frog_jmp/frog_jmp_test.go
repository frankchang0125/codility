package frogjmp

import (
	"testing"
)

func TestSolution(t *testing.T) {
    var ans, expectAns int
    
    ans = Solution(1, 1, 30)
	expectAns = 0
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
    
    ans = Solution(1, 2, 30)
	expectAns = 1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }

	ans = Solution(10, 85, 30)
	expectAns = 3
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
    
    ans = Solution(85, 10, 30)
	expectAns = 3
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
    
    ans = Solution(1, 100, 30)
	expectAns = 4
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
    
    ans = Solution(1, 101, 1)
	expectAns = 100
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
    
    ans = Solution(1, 10, 100000)
	expectAns = 1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
    
    ans = Solution(1, 100000, 10)
	expectAns = 10000
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
}
