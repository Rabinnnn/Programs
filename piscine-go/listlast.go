package piscine

// import "fmt"

// NodeL represents a node in the linked list.
/**type NodeL struct {
	Data interface{}
	Next *NodeL
}

// List represents a linked list.
type List struct {
	Head *NodeL
	Tail *NodeL
} **/

// ListLast returns the Data interface{} of the last element in the linked list l.
func ListLast(l *List) interface{} {
	// Check if the linked list is empty
	if l.Head == nil {
		return nil
	}

	// Start from the head of the list
	current := l.Head

	// Traverse the list until the last node
	for current.Next != nil {
		current = current.Next
	}

	// Return the Data of the last node
	return current.Data
}

// ListPushBack adds a new node with the given data to the end of the linked list.
/**func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data, Next: nil}

	if l.Tail == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}

func main() {
	link := &List{}
	link2 := &List{}

	ListPushBack(link, "three")
	ListPushBack(link, 3)
	ListPushBack(link, "1")

	fmt.Println(ListLast(link))  // Output: 1
	fmt.Println(ListLast(link2)) // Output: <nil>
} **/
