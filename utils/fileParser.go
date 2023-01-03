package utils

import (
	"bufio"
	"fmt"
	"github.com/TheLe0/f1-race-analysis/types"
	"os"
	"strings"
)

func ParseFile(filePath string) []types.RacerInput {

	file, err := os.Open(filePath)

	var racers []types.RacerInput

	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	i := 0

	for fileScanner.Scan() {
		racers[i] = ParseLiner(fileScanner.Text())
		i++
	}

	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}

	return racers
}

func ParseLiner(line string) types.RacerInput {

	line = strings.Replace(line, "\t\t", "–", -1)
	line = strings.Replace(line, " – ", "–", -1)
	line = strings.Replace(line, "\t\t\t", "–", -1)
	line = strings.Replace(line, "\t", "–", -1)

	list := strings.Split(line, "–")

	var input types.RacerInput

	input.LocalTime = list[0]
	input.Number = list[1]
	input.Name = list[2]
	input.LapNumber = list[3]
	input.LapTime = list[4]
	input.AverageLapSpeed = list[5]

	return input
}
