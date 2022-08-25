package piscine

// type TreeNode struct {
// 	Left, Right, Parent *TreeNode
// 	Data                string
// }

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	} else if elem > root.Data {
		return BTreeSearchItem(root.Right, elem)
	} else {
		return root
	}
}

// func BTreeInsertData(root *TreeNode, data string) *TreeNode {
// 	if root == nil {
// 		return &TreeNode{Data: data}
// 	}

// 	if data < root.Data {
// 		root.Left = BTreeInsertData(root.Left, data)
// 		root.Left.Parent = root
// 	} else if data > root.Data {
// 		root.Right = BTreeInsertData(root.Right, data)
// 		root.Right.Parent = root
// 	}
// 	return root
// }
