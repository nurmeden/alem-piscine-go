package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

/*
func ListSort(l *List) *NodeI {
temp := l.Head
clearListLink := l
clearListLink.Head = nil
clearListLink.Tail = nil
for i := temp; i != nil; i = i.Next{
	if l.Head.Data > i.Data{
ListPushBack(clearListLink, i)
	} else if l.Tail.Data < i.Data{
		ListPushBack(clearListLink, i)
	} else {
		u := l.Head.Next
	}
}
}
*/

/*
func ListSort(l *NodeI) *NodeI {
temp := l
for i := temp; i != nil; i = i.Next{
	for j := i.Next; j != nil; j = j.Next{
		if i.Data > j.Data{
			i.Data, j.Data = j.Data, i.Data
		}
	}
}
l= temp
return l
}
*/

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
}
