package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ParseFile(filePath string) [][]string {

	file, err := os.Open(filePath)

	var listString [][]string

	if err != nil {
		fmt.Printf("Could not open the file due to the following error \n")
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		listString = append(listString, parseLine(fileScanner.Text()))
	}

	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to the following error \n")
		panic(err)
	}

	return listString
}

func parseLine(line string) []string {

	stringArray := strings.Split(line, ";")

	return stringArray
}
