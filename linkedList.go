package ds

import (
	"fmt"
)

type ILinkedList interface {
	Insert(int)
	DeleteLast()
	Display()
	Reverse() *LinkedList
}

type LinkedList struct {
	d    int
	link *LinkedList
}

func (head *LinkedList) Insert(i int) {
	node := &LinkedList{d: i}
	if head == nil {
		head = node
	}
	temp := head
	for temp.link != nil {
		temp = temp.link
	}
	temp.link = node
}

func (head *LinkedList) DeleteLast() {
	if head.link == nil {
		return
	}
	temp := head.link
	prev := head.link
	for temp.link != nil {
		prev = temp
		temp = temp.link
	}

	prev.link = nil
}

func (head *LinkedList) Display() {
	if head.link == nil {
		return
	}
	temp := head.link
	for temp.link != nil {
		fmt.Println(temp.d)
		temp = temp.link
	}
	fmt.Println(temp.d)
}

func (head *LinkedList) Reverse() *LinkedList {
	var prev *LinkedList
	var next *LinkedList
	head = head.link
	for head != nil {
		next = head.link
		head.link = prev
		prev = head
		head = next
	}
	return prev
}
