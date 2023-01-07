package main

import (
	"fmt"
	"github.com/TheLe0/f1-race-analysis/commands"
)

func main() {

	var raceLaps uint8
	var filePath string

	fmt.Print("File path of the race log: ")
	_, err := fmt.Scan(&filePath)

	if err != nil {
		fmt.Printf("Error reading the input for the file path, please inform a valid information \n")
		return
	}

	fmt.Print("Total number of laps of the race: ")
	_, err = fmt.Scan(&raceLaps)

	if err != nil {
		fmt.Printf("Error reading the input for the total of laps, please inform a valid information \n")
		return
	}

	commands.AnalyzeRace(filePath, raceLaps)
}
