// The inorder traversal of a binary tree
package learninggo

// tree
type node struct {
	left  *node
	right *node
	val   int
}

// helper function to create a tree. Zero represents an empty array val
func parseArrayToTree(treeArr []int) node {
	if len(treeArr) == 0 {
		return node{}
	}
	els := []*node{}
	for i, val := range treeArr {
		if val != 0 {
			n := node{val: val}
			els = append(els, &n)
			if i > 0 && val != 0 { // not a root, not empty
				pos := (i - 1) / 2 //father
				if i%2 == 0 {      //right node
					els[pos].right = &n
				} else { //left node
					els[pos].left = &n
				}
			}
		}
	}
	return *els[0]
}

// solution with recursion
func getInOrderRecursive(root node) []int {
	if root.val == 0 {
		return []int{}
	}
	res := []int{}
	if root.left != nil {
		res = append(res, getInOrderRecursive(*root.left)...)
	}
	res = append(res, root.val)
	if root.right != nil {
		res = append(res, getInOrderRecursive(*root.right)...)
	}
	return res
}

// solution with a stack
// append at end consume from end
func getInOrderStack(root node) []int {
	if root.val == 0 {
		return []int{}
	}
	visited := map[node]bool{}
	pile := []node{root}
	res := []int{}
	for len(pile) > 0 {
		el := pile[len(pile)-1]
		pile = pile[:len(pile)-1]

		if _, ok := visited[el]; !ok {
			if el.right != nil {
				pile = append(pile, *el.right)
			}
			pile = append(pile, el)
			if el.left != nil {
				pile = append(pile, *el.left)
			}
			visited[el] = true
		} else {
			res = append(res, el.val)
		}
	}
	return res
}
