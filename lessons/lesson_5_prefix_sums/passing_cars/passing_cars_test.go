package passingcars

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var ans, expectAns int
	var arr []int

	ans = Solution([]int{0, 1, 0, 1, 1})
	expectAns = 5
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{0, 0})
	expectAns = 0
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{0, 1})
	expectAns = 1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{1, 0})
	expectAns = 0
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{1, 1})
	expectAns = 0
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	arr = make([]int, 10000*2)
	for i := range arr {
		if i & 1 == 0 {
			arr[i] = 1
		}
	}
	ans = Solution(arr)
	expectAns = ((10000*10000)-10000)/2
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	arr = make([]int, 20000*2)
	for i := range arr {
		if i & 1 == 0 {
			arr[i] = 1
		}
	}
	ans = Solution(arr)
	expectAns = 199990000
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	arr = make([]int, 44721*2)
	for i := range arr {
		if i & 1 == 0 {
			arr[i] = 1
		}
	}
	ans = Solution(arr)
	expectAns = 999961560
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	arr = make([]int, 50000*2)
	for i := range arr {
		if i & 1 == 0 {
			arr[i] = 1
		}
	}
	ans = Solution(arr)
	expectAns = -1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
}
