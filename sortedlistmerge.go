package piscine

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}

	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	for i := n1; i != nil; i = i.Next {
		listPushBack(n2, i.Data)
	}

	for i := n2; i != nil; i = i.Next {
		for j := i.Next; j != nil; j = j.Next {
			if i.Data > j.Data {
				i.Data, j.Data = j.Data, i.Data
			}
		}
	}

	return n2
}
