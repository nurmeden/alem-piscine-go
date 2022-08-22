package piscine

// import (
// 	"fmt"
// )

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	result := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = result
		l.Tail = result
	} else {
		l.Tail.Next = result
		l.Tail = result
	}
}

// func main() {
// 	link := &List{}

// 	ListPushBack(link, "Hello")
// 	ListPushBack(link, "man")
// 	ListPushBack(link, "how are you")

// 	for link.Head != nil {
// 		fmt.Println(link.Head.Data)
// 		link.Head = link.Head.Next
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// type node struct {
// 	data int
// 	next *node
// }

// func insert(head *node, data int) *node {
// 	n := &node{data: data}

// 	if head == nil {
// 		return n
// 	} else {
// 		n.next = head
// 		return n
// 	}
// }

// func PrintList(head *node) {
// 	for head != nil {
// 		fmt.Print(head.data, " -> ")
// 		head = head.next
// 	}
// 	fmt.Println(nil)
// }

// func main() {
// 	var link *node
// 	link = insert(link, 1)
// 	link = insert(link, 2)
// 	link = insert(link, 3)

// 	PrintList(link)
// }
