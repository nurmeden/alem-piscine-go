package piscine

// type NodeI struct {
// 	Data int
// 	Next *NodeI
// }

/*
func ListSort(l *NodeI) *NodeI {
	temp := l                             // temp нужен чтоб непотерять остальные кодыначинаем с l ч
	for i := temp; i != nil; i = i.Next { // проходимся по элементам
		for j := i.Next; j != nil; j = j.Next {
			if i.Data > j.Data { // если больше значение то
				i.Data, j.Data = j.Data, i.Data // меняем
			}
		}
	}
	l = temp // поменяем лист с временным листом
	return l
}*/

/*
func SortListInsert(l *NodeI, data_ref int) *NodeI {
	if data_ref < l.Data {
		n := &NodeI{}
		n.Next = l.Next
		l.Next = n
	}
	// if data_ref > l.Tail.Data {
	//
	for i := l; i != nil; i = i.Next {
		if data_ref > i.Data && data_ref < i.Next.Data {
			n := &NodeI{}
			n.Next = i.Next
			i.Next = n
		}
	}
	return l
}*/

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	n := &NodeI{data_ref, nil}

	n.Next = l
	l = n

	temp := l

	for i := temp; i != nil; i = i.Next {
		for j := i.Next; j != nil; j = j.Next {
			if i.Data > j.Data {
				i.Data, j.Data = j.Data, i.Data
			}
		}
	}

	return temp
}
