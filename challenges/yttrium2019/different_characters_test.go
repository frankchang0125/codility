package differentcharacters

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var ans, expectAns int
	
	ans = Solution("abaacbca", 2)
	expectAns = 3
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution("aabcabc", 1)
	expectAns = 5
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution("zaaaa", 1)
	expectAns = 1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
	
	ans = Solution("aaaa", 2)
	expectAns = -1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
	
	ans = Solution("a", 0)
	expectAns = 1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
	
	ans = Solution("ab", 0)
	expectAns = 2
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
	
	ans = Solution("ab", 2)
	expectAns = 0
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution("abababab", 1)
	expectAns = 7
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution("abababab", 2)
	expectAns = 0
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution("abcccab", 3)
	expectAns = 0
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution("abcccab", 2)
	expectAns = 3
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution("abcccab", 1)
	expectAns = 6
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution("abcccab", 0)
	expectAns = 7
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution("abcacab", 3)
	expectAns = 0
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution("abcacab", 2)
	expectAns = 3
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution("abcacab", 1)
	expectAns = 6
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
}
