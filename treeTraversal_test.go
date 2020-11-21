package learninggo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Testing tree parsing
// Creating a very simple tree
func TestParseRoot(t *testing.T) {
	tree1 := parseArrayToTree([]int{1})
	assert.Equal(t, tree1.val, 1)
	assert.Nil(t, tree1.left)
	assert.Nil(t, tree1.right)

}

// a tree with root and children
func TestParseChildren(t *testing.T) {
	tree2 := parseArrayToTree([]int{1, 2, 3})
	assert.Equal(t, tree2.val, 1)
	assert.Equal(t, tree2.left.val, 2)
	assert.Nil(t, tree2.left.right)
	assert.Nil(t, tree2.left.left)
	assert.Equal(t, tree2.right.val, 3)
	assert.Nil(t, tree2.right.right)
	assert.Nil(t, tree2.right.left)
}

// testing a tree with zero
func TestParseOneChildre(t *testing.T) {
	tree3 := parseArrayToTree([]int{1, 0, 2, 3})
	assert.Equal(t, tree3.val, 1)
	assert.Nil(t, tree3.left)
	assert.Equal(t, tree3.right.val, 2)
	assert.Nil(t, tree3.right.right)
	assert.Equal(t, tree3.right.left.val, 3)
}

// just a passing test with a bigger entry
func TestParseLongEntry(t *testing.T) {
	tree3 := parseArrayToTree([]int{1, 2, 3, 0, 1, 2, 1, 3, 2, 1, 0, 1, 0, 1, 2, 3, 4, 2, 1, 2, 1, 2, 3})
	assert.NotNil(t, tree3)
	assert.NotNil(t, tree3.right)
	assert.NotNil(t, tree3.left)
}

// Test table for inorder algorithm
var orderingTests = [][][]int{
	{{1, 0, 2, 3}, {1, 3, 2}},
	{{1, 4, 2, 3}, {3, 4, 1, 2}},
	{{1, 5, 2, 0, 4, 7, 0, 9, 10, 23, 14, 13, 0, 16}, {5, 13, 9, 4, 16, 10, 1, 23, 7, 14, 2}},
	{{}, {}},
	{{1}, {1}},
	{{1, 2}, {2, 1}},
	{{1, 0, 2}, {1, 2}},
}

// recursive approach
func TestInOrderRecursive(t *testing.T) {
	for _, check := range orderingTests {
		assert.Equal(t, check[1], getInOrderRecursive(parseArrayToTree(check[0])))
	}
}

// stack approach
func TestInOrderStack(t *testing.T) {
	for _, check := range orderingTests {
		assert.Equal(t, check[1], getInOrderStack(parseArrayToTree(check[0])))
	}
}
