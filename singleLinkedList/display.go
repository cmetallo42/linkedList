package singlelinkedlist

import "log"

func (list *List) Display() {
	for {
		log.Println(list.head)
		if list.head.next == nil {
			break
		}
		list.head = list.head.next
	}
}

func (list *List) Len() uint {
	return list.len
}
