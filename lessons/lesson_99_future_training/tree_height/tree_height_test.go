package treeheight

import (
	"testing"
)

type Tree struct {
	X int
	L *Tree
	R *Tree
}

func TestSolution(t *testing.T) {
	var ans, expectAns int
	var tree *Tree

	tree = &Tree{
		X: 5,
		L: &Tree{
			X: 3,
			L: &Tree{
				X: 20,
			},
			R: &Tree{
				X: 21,
			},
		},
		R: &Tree{
			X: 10,
			L: &Tree{
				X: 1,
			},
		},
	}
	ans = Solution(tree)
	expectAns = 2
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
}
