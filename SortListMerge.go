package student

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	iterator := n2
	for iterator != nil {
		SortListInsert(n1, iterator.Data)
		iterator = iterator.Next
	}

	return n1
}
