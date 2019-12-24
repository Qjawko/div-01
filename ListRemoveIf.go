package student

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	node := &NodeL{data, nil}

	if l.Head == nil {
		l.Head = node
		return
	}

	l.Tail = node
	iter := l.Head
	for iter.Next != nil {
		iter = iter.Next
	}
	iter.Next = node
}

func ListRemoveIf(l *List, data_ref interface{}) {
	newList := &List{}
	iterator := l.Head

	for iterator != nil {
		if iterator.Data != data_ref {
			ListPushBack(newList, iterator.Data)
		}
		iterator = iterator.Next
	}
	*l = *newList
}
