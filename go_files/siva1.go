package main

import "fmt"

//class node
type Node struct {
	element int
	next    *Node
}

//class list
type List struct {
	head *Node //nil
	tail *Node //nil
	len  int   //0
}

func (obj *List) printing() {
	comm := obj.head //n1
	for i := 0; i < obj.len; i++ {
		fmt.Printf("%v==>", comm.element)
		comm = comm.next
	}
	fmt.Println()

}

//adding node at beginning
func (obj *List) insertbegg(val int) {
	new := Node{element: val}
	new.next = obj.head
	obj.head = &new
	obj.len++
}

//delete at beginning
func (obj *List) deletebegg() {
	comm := obj.head //x
	obj.head = comm.next
	obj.len--
}

//insert at end
func (obj *List) insertend(val int) {
	newnode := Node{element: val}
	if obj.len == 0 {
		obj.head = &newnode
		obj.tail = &newnode
	} else {
		obj.tail.next = &newnode
	}
	obj.len++
}

//delete at end
func (obj *List)deleteend() {
	if obj.len == 0 {
		fmt.Println("no elements")
		return
	}
	obj.tail.next = nil
	obj.tail = nil
	obj.len--

}

//adding nodes
func (obj *List) adding(val int) {
	newnode := Node{element: val}
	if obj.len == 0 {
		obj.head = &newnode
	} else {
		obj.tail.next = &newnode
	}
	obj.tail = &newnode
	obj.len++

}

func main() {
	obj := List{}
	obj.adding(10)
	obj.adding(20)
	obj.adding(30)
	obj.printing()
	obj.insertbegg(100)
	obj.printing()
	obj.deletebegg()
	obj.printing()
	obj.adding(200)
	obj.printing()
	obj.deleteend()
	obj.printing()

}
