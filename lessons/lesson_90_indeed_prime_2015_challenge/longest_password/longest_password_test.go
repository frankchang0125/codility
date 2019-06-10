package longestpassword

import (
    "testing"
)

func TestSolution(t *testing.T) {
    var ans, expectAns int

	ans = Solution("test 5 a0A pass007 ?xy1")
	expectAns = 7
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }

    ans = Solution("")
	expectAns = -1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }

    ans = Solution("#")
	expectAns = -1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }

    ans = Solution("1a2b3c4d5 1a2b3c4d5e6f7")
	expectAns = 13
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }

    ans = Solution("1a2b3c4d5 1a2b3c4d5e6f7?")
	expectAns = 9
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
}
