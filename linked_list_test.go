package linkedlist

import (
	"testing"
)

func TestInsertFirst(t *testing.T) {
	for _, tt := range insertFirstTestCases {
		t.Log(tt.description)
		l := List{}
		l.insertFirst(&tt.input)

		if tt.expectedLen != l.size {
			t.Errorf("FAIL: Return value expected %d but got %d", tt.expectedLen, l.size)
			return
		}

		if tt.expectedData != l.head.data {
			t.Errorf("FAIL: Return value expected %d but got %d", tt.expectedData, l.head.data)
			return
		}

		hasNext := l.head.next != nil
		if tt.expectedNext != hasNext {
			t.Errorf("FAIL: Return value expected %d but got %d", tt.expectedData, l.head.data)
			return
		}

		t.Logf("PASS: Expected length %d and got %d", tt.expectedLen, l.size)
		t.Logf("PASS: Expected data %d and got %d", tt.expectedData, l.head.data)
		t.Logf("PASS: Expected data %t and got %t", tt.expectedNext, hasNext)
	}
}

func TestInsertFirstFromArray(t *testing.T) {
	for _, tt := range insertFirstTestCases {
		t.Log(tt.description)
		l := List{}
		n1 := Node{data: 12}
		n2 := Node{data: 2}
		n3 := Node{data: 8}

		l.insertFirst(&n1)
		l.insertFirst(&n2)
		l.insertFirst(&n3)

		if &n3 != l.head {
			t.Errorf("FAIL: First list value expected %+v but got %+v", &n3, l.head)
			return
		}

		t.Logf("PASS: First list value expected %+v but got %+v", &n3, l.head)
	}
}

func TestInsertLastFromArray(t *testing.T) {
	for _, tt := range insertFirstTestCases {
		t.Log(tt.description)
		l := List{}
		n1 := Node{data: 12}
		n2 := Node{data: 2}
		n3 := Node{data: 8}

		l.insertLast(&n1)
		l.insertLast(&n2)
		l.insertLast(&n3)

		last := l.getLast()

		if &n3 != last {
			t.Errorf("FAIL: Last list value expected %+v but got %+v", &n3, l.head)
			return
		}

		t.Logf("PASS: Last list value expected %+v but got %+v", &n3, l.head)
	}
}

func TestGetFirst(t *testing.T) {
	for _, tt := range getFirstTestCases {
		t.Log(tt.description)
		l := List{}
		l.insertFirst(&tt.input)

		firstNode := l.getFirst()

		if firstNode != &tt.input {
			t.Errorf("FAIL: Return value expected %+v but got %+v", &tt.input, firstNode)
			return
		}

		t.Logf("PASS: Expected %+v and got %+v", &tt.input, firstNode)
	}
}
