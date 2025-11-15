package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, "", "")
}

func isBST(node *TreeNode, min, max string) bool {
	if node == nil {
		return true
	}
	if (min != "" && node.Data <= min) || (max != "" && node.Data >= max) {
		return false
	}
	return isBST(node.Left, min, node.Data) && isBST(node.Right, node.Data, max)
}
