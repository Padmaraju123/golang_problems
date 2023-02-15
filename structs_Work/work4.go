package main

import "fmt"

type Node struct{
	element int
	next *Node
}

type List struct{
	head *Node
	tail *Node
	len int
}

func (obj1 *List)printing1(){
	comm := obj1.head
	for i:=0;i<obj1.len;i++{
		fmt.Printf("%v==>",comm.element)
		comm = comm.next
	}
	fmt.Println()

}

func (obj2 *List)printing2(){
	comm := obj2.head
	for i:=0;i<obj2.len;i++{
		fmt.Printf("%v==>",comm.element)
		comm = comm.next
	}
	fmt.Println()

}

func merging(obj1 *List,obj2 *List,a int,b int){
	comm := obj1.head
	for p:=0;p<2;p++{
		if p==1{
			
			comm.next.next = obj2.head
			obj2.tail = comm.next.next.next
		}
		comm=comm.next

	}
	

}

func (obj1 *List)adding(val int){
	newnode := Node{element:val}
	if obj1.len==0{
		obj1.head = &newnode
	}else{
		obj1.tail.next = &newnode
	}
	obj1.tail = &newnode
	obj1.len++
}

func main(){
	obj1 := List{}
	obj2 := List{}
	sli1 := []int{0,1,2,3,4,5}
	for _,v := range sli1{
		obj1.adding(v)
	}
	sli2 := []int{1000,1001,1002}
	for _,vv := range sli2{
		obj2.adding(vv)
	}

	obj1.printing1()
	obj2.printing2()

	merging(&obj1,&obj2,3,4)
	obj1.printing1()



	
}