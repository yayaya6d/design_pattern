package iterator

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type linkedListTestSuite struct {
	suite.Suite
}

func TestLinkedListTestSuite(t *testing.T) {
	suite.Run(t, new(linkedListTestSuite))
}

func (s *linkedListTestSuite) TestAddNodeAndTraversal_AddSomeNode_TraversalListCorrectly() {
	// arrange
	linkedList := NewLinkedList()
	expectedList := []int{1, 2, 3, 4}

	// act
	linkedList.AddNode(1)
	linkedList.AddNode(2)
	linkedList.AddNode(3)
	linkedList.AddNode(4)

	actualList := make([]int, 0)
	iter := linkedList.CreateIterator()
	for iter.HasNext() {
		val, err := iter.Next()
		if err != nil {
			s.Failf("error occurred while traversal, msg: %s", err.Error())
		}
		actualList = append(actualList, val)
	}

	// assert
	s.Equal(expectedList, actualList)
}

func (s *linkedListTestSuite) TestHasNext_DoNotAddAnyNode_ReturnFalse() {
	// arrange
	linkedList := NewLinkedList()

	// act
	hasNext := linkedList.CreateIterator().HasNext()

	// assert
	s.False(hasNext)
}

func (s *linkedListTestSuite) TestNext_DoNotAddAnyNode_PanicWhenVallNext() {
	// arrange
	linkedList := NewLinkedList()

	// act
	val, err := linkedList.CreateIterator().Next()

	// assert
	s.Equal(-1, val)
	s.Error(err)
}
