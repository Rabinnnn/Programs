package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil { // if empty return 0
		return 0
	}
	left := BTreeLevelCount(root.Left)
	right := BTreeLevelCount(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}
