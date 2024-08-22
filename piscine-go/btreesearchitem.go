package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil { // if empty return nil
		return nil
	}
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem) // if the element is less than the root traverse the left side.
	} else if elem > root.Data {
		return BTreeSearchItem(root.Right, elem) // if the element is greater than the root traverse the right side.
	} else {
		return root
	}
}
