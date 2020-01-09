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
	node := &NodeL{data, nil} //создаем элемента листа который мы хотим вставить в конец

	if l.Head == nil { //если начало листа nil
		l.Head = node //делаем наш node как начало
		return //конец работы функции
	}

	l.Tail = node //?_?
	iter := l.Head //создаем итератор с помощью которого мы будем ходить по листу
	for iter.Next != nil { //пока следующий элемент после iter не будет равен nil
		iter = iter.Next //идем дальше по листу
	}
	iter.Next = node //после того как закончился цикл, следующий эелемент после iter равен nil 
			 //соответсвенно это значит что это последний элемент в list, вставляем наш node следующим после последнего элемента
} //конец функции

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
