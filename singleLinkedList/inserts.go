package singlelinkedlist

func (list *List) InsertFront(node *Node) {
	if list.head == nil {
		list.head = node
	} else {
		node.next = list.head
		list.head = node
	}

	list.len++
}

func (list *List) InsertFrontAny(a any) {
	node := Node{
		Data: a,
	}
	if list.head == nil {
		list.head = &node
	} else {
		node.next = list.head
		list.head = &node
	}

	list.len++
}

func (list *List) InsertLast(node *Node) {
	if list.head == nil {
		list.InsertFront(node)
		return
	}
	if list.len == 0 {
		list.head = node
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
	list.len++
}

func (list *List) InsertLastAny(a any) {
	node := &Node{Data: a}
	if list.head == nil {
		list.InsertFront(node)
		return
	}

	list.InsertLast(node)
}
