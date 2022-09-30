package piscine

import (
	"fmt"
)

func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	i := l
	for i.Next != nil {
		i = i.Next
	}
	i.Next = n
	return l
}

type NodeI struct {
	Data interface{}
	Next *NodeI
}

type List struct {
	Head *NodeI
	Tail *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}
	var current *NodeI
	var current2 *NodeI
	var tmp int

	current = l
	for current.Next != nil {
		current2 = current.Next
		for current2 != nil {
			if current.Data.(int) > current2.Data.(int) {
				tmp = current.Data.(int)
				current.Data = current2.Data
				current2.Data = tmp
			}
			current2 = current2.Next
		}
		current = current.Next
	}
	return l
}
