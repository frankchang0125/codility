package cyclicrotation

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	var ans, expectAns []int

	ans = Solution([]int{3, 8, 9, 7, 6}, 0)
	expectAns = []int{3, 8, 9, 7, 6}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{3, 8, 9, 7, 6}, 1)
	expectAns = []int{6, 3, 8, 9, 7}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{3, 8, 9, 7, 6}, 2)
	expectAns = []int{7, 6, 3, 8, 9}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{3, 8, 9, 7, 6}, 3)
	expectAns = []int{9, 7, 6, 3, 8}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{3, 8, 9, 7, 6}, 4)
	expectAns = []int{8, 9, 7, 6, 3}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{3, 8, 9, 7, 6}, 5)
	expectAns = []int{3, 8, 9, 7, 6}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{3, 8, 9, 7, 6}, 6)
	expectAns = []int{6, 3, 8, 9, 7}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{3, 8, 9, 7, 6}, 7)
	expectAns = []int{7, 6, 3, 8, 9}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{0, 0, 0}, 1)
	expectAns = []int{0, 0, 0}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{1, 2, 3, 4}, 4)
	expectAns = []int{1, 2, 3, 4}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{1, 2, 3, 4}, 8)
	expectAns = []int{1, 2, 3, 4}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{1, 2, 3, 4}, 1)
	expectAns = []int{4, 1, 2, 3}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution([]int{}, 5)
	expectAns = []int{}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
}
