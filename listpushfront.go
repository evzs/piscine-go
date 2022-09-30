package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data string) {
	n := &NodeL{Data: data}
	var new_node *NodeL = new(NodeL)
	new_node.Data = data
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}

	new_node.Next = l.Head
	l.Head = new_node
}
