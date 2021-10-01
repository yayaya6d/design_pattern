package iterator

import "errors"

type Node struct {
	Val  int
	Next *Node
}

type LinkedList interface {
	Collection
	AddNode(int)
}

type LinkedListIter struct {
	Iterator
	cur *Node
}

func (lli *LinkedListIter) Next() (int, error) {
	if lli.cur == nil {
		return -1, errors.New("please make sure list has next element by HasNext()")
	}

	ret := lli.cur.Val
	lli.cur = lli.cur.Next

	return ret, nil
}

func (lli *LinkedListIter) HasNext() bool {
	return lli.cur != nil
}

type linkedList struct {
	head *Node
	end  *Node
}

func NewLinkedList() LinkedList {
	return &linkedList{
		head: nil,
		end:  nil,
	}
}

func (ll *linkedList) CreateIterator() Iterator {
	return &LinkedListIter{
		cur: ll.head,
	}
}

func (ll *linkedList) AddNode(val int) {
	newNode := &Node{
		Val: val,
	}

	if ll.head == nil {
		ll.head = newNode
		ll.end = ll.head
	} else {
		if ll.head == ll.end {
			ll.end = newNode
			ll.head.Next = newNode
		} else {
			ll.end.Next = newNode
			ll.end = newNode
		}
	}
}
