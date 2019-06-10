package maxcounters

import (
	"reflect"
	"testing"
)

func TestSolution1(t *testing.T) {
	var ans, expectAns []int

	ans = Solution1(5, []int{3, 4, 4, 6, 1, 4, 4})
	expectAns = []int{3, 2, 2, 4, 2}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution1(1, []int{1})
	expectAns = []int{1}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution1(1, []int{1, 2})
	expectAns = []int{1}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution1(1, []int{1, 2, 2})
	expectAns = []int{1}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution1(1, []int{1, 1, 2, 1, 1})
	expectAns = []int{4}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution1(2, []int{1, 1, 1, 2})
	expectAns = []int{3, 1}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution1(2, []int{1, 1, 1, 3})
	expectAns = []int{3, 3}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution1(2, []int{3, 1, 1, 1})
	expectAns = []int{3, 0}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution1(2, []int{3, 1, 1, 2})
	expectAns = []int{2, 1}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution1(2, []int{1, 2, 3, 1, 2})
	expectAns = []int{2, 2}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution1(2, []int{1, 1, 3, 1, 1})
	expectAns = []int{4, 2}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution1(2, []int{1, 1, 3, 2, 2})
	expectAns = []int{2, 4}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
}

func TestSolution2(t *testing.T) {
	var ans, expectAns []int

	ans = Solution2(7, []int{3, 4, 4, 6, 1, 4, 4})
	expectAns = []int{3, 2, 2, 4, 2}
	if reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution2(1, []int{1})
	expectAns = []int{1}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution2(1, []int{1, 2})
	expectAns = []int{1}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution2(1, []int{1, 2, 2})
	expectAns = []int{1}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution2(1, []int{1, 1, 2, 1, 1})
	expectAns = []int{4}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution2(2, []int{1, 1, 1, 2})
	expectAns = []int{3, 1}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution2(2, []int{1, 1, 1, 3})
	expectAns = []int{3, 3}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution2(2, []int{3, 1, 1, 1})
	expectAns = []int{3, 0}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution2(2, []int{3, 1, 1, 2})
	expectAns = []int{2, 1}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution2(2, []int{1, 2, 3, 1, 2})
	expectAns = []int{2, 2}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution2(2, []int{1, 1, 3, 1, 1})
	expectAns = []int{4, 2}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution2(2, []int{1, 1, 3, 2, 2})
	expectAns = []int{2, 4}
	if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
}
