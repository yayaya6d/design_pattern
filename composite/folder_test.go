package composite

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type folderTestSuite struct {
	suite.Suite
	sut Component
}

func TestFolderTestSuite(t *testing.T) {
	suite.Run(t, new(folderTestSuite))
}

func (s *folderTestSuite) SetupSuite() {
	file1 := File{
		name: "file1",
	}
	file2 := File{
		name: "file2",
	}
	rootFolder := Folder{
		name: "root",
	}
	folder1 := Folder{
		name: "folder1",
	}

	folder1.Add(&file2)
	rootFolder.Add(&folder1)
	rootFolder.Add(&file1)

	s.sut = &rootFolder
}

func (s *folderTestSuite) TestSearch_GivenTargetName_ReturnExpectedResult() {
	// arrange
	testCases := []struct {
		name           string
		target         string
		expectedResult bool
	}{
		{
			name:           "given root",
			target:         "root",
			expectedResult: true,
		},
		{
			name:           "given folder1",
			target:         "folder1",
			expectedResult: true,
		},
		{
			name:           "given file1",
			target:         "file1",
			expectedResult: true,
		},
		{
			name:           "given file2",
			target:         "file2",
			expectedResult: true,
		},
		{
			name:           "given non-exist name",
			target:         "unknown",
			expectedResult: false,
		},
	}

	for _, c := range testCases {
		s.Run(c.name, func() {
			// act
			actualResult := s.sut.Search(c.target)

			// assert
			s.Equal(actualResult, c.expectedResult)
		})
	}
}
