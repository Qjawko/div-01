package student

type NodeI struct {
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	node := &NodeI{Data: data_ref}
	if l.Data >= node.Data {
		node.Next = l
		l = node
	} else {
		iterator := l
		for iterator.Next != nil && iterator.Next.Data < node.Data {
			iterator = iterator.Next
		}

		node.Next = iterator.Next
		iterator.Next = node
	}

	return l
}

//1 -> 4 -> 9
//1 -> node -> 4
