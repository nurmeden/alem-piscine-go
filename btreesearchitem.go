package piscine

<<<<<<< HEAD
=======
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

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

>>>>>>> 922cc1cdb52d99f448a5678d8eb8fd7a08b3bb38
func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data == elem {
		return root
<<<<<<< HEAD
	}
	if root.Data < elem {
		return BTreeSearchItem(root.Right, elem)
	} else {
		return BTreeSearchItem(root.Left, elem)
=======
	}
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	} else {
		return BTreeSearchItem(root.Right, elem)
	}
}

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	f(root.Data)
	BTreeApplyPreorder(root.Left, f)
	BTreeApplyPreorder(root.Right, f)
}

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyPostorder(root.Left, f)
	BTreeApplyPostorder(root.Right, f)
	f(root.Data)
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if root.Data > data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
>>>>>>> 922cc1cdb52d99f448a5678d8eb8fd7a08b3bb38
	}
	return nil
}
