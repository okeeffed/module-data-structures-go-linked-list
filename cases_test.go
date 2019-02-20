package linkedlist

type testCase struct {
	description  string
	input        Node
	expectedLen  int
	expectedData int
	expectedNext bool
}

var insertFirstTestCases = []testCase{
	{
		description:  "insertFirst: should insert the correct node at the start head of the list",
		input:        Node{data: 8},
		expectedLen:  1,
		expectedData: 8,
		expectedNext: false,
	},
	{
		description:  "insertFirst: should insert the correct node at the start head of the list",
		input:        Node{data: 2, next: &Node{data: 5}},
		expectedLen:  1,
		expectedData: 2,
		expectedNext: true,
	},
}
