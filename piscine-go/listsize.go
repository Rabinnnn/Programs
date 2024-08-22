package piscine

/*
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
} */

func ListSize(l *List) int {
	current := l.Head
	size := 0

	for current != nil {
		size++
		current = current.Next
	}

	return size
}

/*
func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data, Next: l.Head}
	l.Head = newNode
	if l.Tail == nil {
		l.Tail = newNode
	}
}*/
