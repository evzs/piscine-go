package piscine

import (
	"fmt"
)

func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
		return
	}

	i := l.Head
	for i.Next != nil {
		i = i.Next
	}
	i.Next = n
}

func ListRemoveIf(l *List, data_ref interface{}) {
	tmp := l.Head
	prev := l.Head

	for tmp != nil && tmp.Data == data_ref {
		l.Head = tmp.Next
		tmp = l.Head
	}
	for tmp != nil {
		for tmp != nil && tmp.Data != data_ref {
			prev = tmp
			tmp = tmp.Next
		}

		if tmp == nil {
			return
		}
		prev.Next = tmp.Next
		tmp = prev.Next
	}
}
