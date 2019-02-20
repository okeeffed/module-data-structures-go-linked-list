/*
@author Dennis O'Keeffe

Methods:

1. insertFirst: Insert at the head of list
2. size: Fetch size
3. getFirst: Fetch the first node
4. getLast: Fetch the last node
5. clear: Remove all node
6. removeFirst: Remove head of list
7. removeLast: Remove last node
8. insertLast: Insert at end of list
9. getAt: Fetch node at index
10. removeAt: Remove node at index
11. insertAt: Insert at index
12. forEach: Iterate through list and run function on list
*/

package linkedlist

import (
	"errors"
	"sync"
)

// Node - Setup base node type for linked list
type Node struct {
	data int
	next *Node
}

// List - Basic singularly linked list
type List struct {
	head *Node
	size int
	lock sync.RWMutex
}

func (l *List) insertFirst(n *Node) {
	l.lock.Lock()
	if l.head == nil {
		l.head = n
	} else {
		temp := l.head
		l.head = n
		n.next = temp
	}
	l.size++
	l.lock.Unlock()
}

func (l *List) insertLast(n *Node) {
	l.lock.Lock()
	if l.head == nil {
		l.head = n
	} else {
		last := l.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = n
	}
	l.size++
	l.lock.Unlock()
}

func (l *List) getFirst() *Node {
	l.lock.Lock()
	res := (*l).head
	l.lock.Unlock()
	return res
}

func (l *List) getLast() *Node {
	l.lock.Lock()
	last := l.head
	for {
		if last.next == nil {
			break
		}
		last = last.next
	}

	l.lock.Unlock()
	return last
}

func (l *List) clear() {
	if l.head != nil {
		l.head = nil
	}
}

func (l *List) getAt(index int) (*Node, error) {
	if l.head == nil {
		return nil, errors.New("No elements in list")
	}

	n := l.head
	for i := 0; i <= index; i++ {
		if i == index {
			return n, nil
		}
		n = n.next
	}

	return nil, errors.New("Never made it to correct index")
}

// TODO: removeAt function

// TODO: insertAt function
