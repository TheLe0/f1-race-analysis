package tests

import (
	"testing"

	"github.com/TheLe0/f1-race-analysis/utils"
)

func TestParseFile(t *testing.T) {

	filePath := "../data/log.csv"
	expectedFileLength := 23

	fileParsed := utils.ParseFile(filePath)

	if len(fileParsed) != expectedFileLength {
		t.Fatalf("Expected array length to be %v but found %v", expectedFileLength, len(fileParsed))
	}
}

func TestParseFileNotFoundPanic(t *testing.T) {

	filePath := "log.csv"

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	utils.ParseFile(filePath)
}
