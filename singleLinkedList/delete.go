package singlelinkedlist

func (list *List) NodeAt(i uint) *Node {
	current := list.head
	for i > 1 {
		current = current.next
		i--
	}
	return current
}

func (list *List) DeleteIndex(i uint) {
	tmpPrev := list.NodeAt(i - 1)
	tmpPrev.next = list.NodeAt(i).next
	list.len--
}

func (list *List) DeleteLast() {
	current := list.head
	for current.next.next != nil {
		current = current.next
	}

	current.next = nil

	list.len--
}
