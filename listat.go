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
	if l == nil {
		return
	}

	var new_node *NodeL = new(NodeL)
	new_node.Data = data

	if l.Head == nil {
		l.Head = new_node
		l.Tail = new_node
		return
	}

	l.Tail.Next = new_node
	l.Tail = new_node
}
func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil {
		return nil
	}
	c := 0
	for l.Next != nil {
		if c == pos {
			return l
		}
		l = l.Next
		c++
	}
	return nil
}
