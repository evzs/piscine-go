package piscine

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

func ListReverse(l *List) interface{} {
	if l.Head == nil {
		return nil
	} else {
		current := l.Head
		c := -1
		for current.Next != nil {
			current = current.Next
			c++
		}
		return c
	}
}
