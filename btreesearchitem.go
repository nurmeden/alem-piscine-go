package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root.Right == nil && root.Left == nil {
		return 1
	}
	if root.Right == nil && root.Left != nil {
		return 1 + BTreeLevelCount(root.Left)
	}
	if root.Right != nil && root.Left == nil {
		return 1 + BTreeLevelCount(root.Right)
	}
	if root.Right != nil && root.Left != nil {
		return BTreeLevelCount(root.Left) + BTreeLevelCount(root.Right)
	}
	return 0
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data == elem {
		return root
	}
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	} else {
		return BTreeSearchItem(root.Right, elem)
	}
}
