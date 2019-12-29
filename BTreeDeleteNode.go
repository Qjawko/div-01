package student

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {

	if root == nil {
		return root
	}

	if node == nil {
		return root
	}

	isRoot := false
	if node == root {
		isRoot = true
	}

	isLeft := false
	if !isRoot && node.Data < node.Parent.Data {
		isLeft = true
	}

	if node.Right == nil && node.Left == nil {
		// 0 children
		if isRoot {
			return nil
		}

		if isLeft {
			node.Parent.Left = nil
		} else {
			node.Parent.Right = nil
		}
	} else if node.Left == nil {
		// 1 child
		if isRoot {
			return node.Right
		}

		if isLeft {
			node.Parent.Left = node.Right
		} else {
			node.Parent.Right = node.Right
		}
	} else if node.Right == nil {
		// 1 child
		if isRoot {
			return node.Left
		}

		if isLeft {
			node.Parent.Left = node.Left
		} else {
			node.Parent.Right = node.Left
		}
	} else {
		// 2 children
		min := BTreeFindMin(node.Right)
		node.Data = min.Data
		if isRoot {
			return BTreeDeleteNode(node, min)
		}
		if isLeft {
			node.Parent.Left = BTreeDeleteNode(node, min)
		} else {
			node.Parent.Right = BTreeDeleteNode(node, min)
		}
	}

	return root
}

// BTreeFindMin q
func BTreeFindMin(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	iterator := root
	for iterator.Left != nil {
		iterator = iterator.Left
	}
	return iterator
}
