package student

type NodeI struct {
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	node := &NodeI{Data: data_ref}
	if l.Data > node.Data {
		node.Next = l
		return node
	}

	iterator := l
	for iterator.Next != nil {
		if iterator.Data < node.Data && iterator.Next.Data > node.Data {
			node.Next = iterator.Next
			iterator.Next = node
		}

		iterator = iterator.Next
	}

	if iterator.Data < node.Data {
		iterator.Next = node
	}
	return l
}

//1 -> 4 -> 9
//1 -> node -> 4
