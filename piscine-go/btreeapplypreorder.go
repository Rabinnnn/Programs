package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	f(root.Data) // apply function preorder
	BTreeApplyPreorder(root.Left, f)
	BTreeApplyPreorder(root.Right, f)
}
