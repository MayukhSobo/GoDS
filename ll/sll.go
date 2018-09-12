package main

import "fmt"

// Node A node for storing data into the link
type Node struct {
	Data int
	Next *Node
}

// LL Singly link list
type LL struct {
	Head *Node
	Curr *Node
	Size int
}

// Push push back each element in the link list
func Push(ll *LL, data int) {
	if ll.Head == nil {
		var head Node
		head.Data = data
		head.Next = nil
		ll.Head = &head
		ll.Curr = &head
	} else {
		curr := ll.Head
		for {
			if curr.Next == nil {
				break
			}
			curr = curr.Next
		}
		var temp Node
		temp.Data = data
		temp.Next = nil
		curr.Next = &temp
		ll.Curr = curr.Next
	}
	ll.Size++
}

// PushBulk Push all the elements of the list
func PushBulk(ll *LL, data []int) {
	for _, each := range data {
		Push(ll, each)
	}
}

// PushAt Push an element at an index
func PushAt(ll *LL, data int, index int) {

}

func main() {
	ll := LL{nil, nil, 0}
	Push(&ll, 10)
	Push(&ll, 20)
	Push(&ll, 30)
	Push(&ll, 40)
	Push(&ll, 50)
	Push(&ll, 60)
	Push(&ll, 70)
	Push(&ll, 80)
	PushBulk(&ll, []int{90, 100, 110, 120})
	temp := ll.Head
	for {
		fmt.Println(temp.Data)
		temp = temp.Next
		if temp == nil {
			break
		}
	}
	// fmt.Println(ll.Size)
}
