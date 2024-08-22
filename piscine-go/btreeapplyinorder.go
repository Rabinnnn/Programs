package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil { // empty tree
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data) // apply function to each element
	BTreeApplyInorder(root.Right, f)
}
