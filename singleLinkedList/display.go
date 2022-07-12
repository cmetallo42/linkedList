package singlelinkedlist

import "log"

func (list *List) DisplayList() {
	if list.len == 0 {
		log.Println("empty list")
		return
	}
	current := list.head
	for i := 0; uint(i) < list.len; i++{
		log.Println(current)
		current = current.next
	}
} 

func (list *List) DisplayLen() {
	log.Println(list.len)
}

func (list *List) Len() uint {
	return list.len
}

func(list *List) DisplayData(i uint = 0) {
	if list.len == 0 {
		log.Println("empty list")
		return
	}
	if i != 0 {
		log.Println(NodeAt(i))
	}
	current := list.head
	for i := 0; uint(i) < list.len; i++{
		log.Println(current.Data)
		current = current.next
	}
}