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
