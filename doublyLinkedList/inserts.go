package doublylinkedlist

func (dlist *DList) InsertFirst(node *Node) {
	if dlist.head == nil {
		dlist.head = node
		dlist.tail = node
	} else {
		node.next = dlist.head
		dlist.head.prev = node
		dlist.head = node
	}
	dlist.len++
}
