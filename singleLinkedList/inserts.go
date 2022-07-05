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
	for {
		if list.head.next == nil {
			list.head.next = node
			break	
		}
		list.head = list.head.next
	}

	list.len++
}

func (list *List) InsertLastAny(a any) {
	node := Node{
		Data: a,
	}
	list.len++

	if list.head == nil {
		list.head = &node
		return
	}

	for {
		if list.head.next == nil {
			list.head.next = &node
			return	
		}
	
		list.head = list.head.next
	}
}