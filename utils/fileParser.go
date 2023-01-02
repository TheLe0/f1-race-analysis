package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ParseFile(filePath string) {

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}
}
