package student

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Data == elem {
		return root
	}
		//5         2
	if root.Data > elem {
		root = BTreeSearchItem(root.Left, elem)
	} else {
		root = BTreeSearchItem(root.Right, elem)
	}

	return root
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}  //create node with new data ex : 6 
	}
    //     7        6
	if root.Data > data {
		root.Left = BTreeInsertData(root.Left, data) // nil   6
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data) // 7   6
		root.Right.Parent = root
	}

	return root
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	node = BTreeSearchItem(root, node.Data)
	node.Data = rplc.Data

	return root
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}
