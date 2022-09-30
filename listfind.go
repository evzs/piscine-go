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

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	i := l.Head
	for i != nil {
		if comp(i.Data, ref) {
			return &i.Data
		}
		i = i.Next
	}
	return nil

}

func CompStr(a, b interface{}) bool {
	if a.(string) == b.(string) {
		return true
	}
	return false
}
