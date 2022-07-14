package doublylinkedlist

type (
	Node struct {
		Data any
		next *Node
		prev *Node
	}

	DList struct {
		len  uint
		head *Node
		tail *Node
	}
)
