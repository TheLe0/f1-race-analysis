package commands

import (
	"github.com/TheLe0/f1-race-analysis/types"
)

func ConvertToRacer(parsedString [][]string) []types.RacerInput {

	var racerInputList []types.RacerInput

	for _, line := range parsedString {

		var racerInput types.RacerInput

		racerInput.LocalTime = line[0]
		racerInput.Number = line[1]
		racerInput.Name = line[2]
		racerInput.LapNumber = line[3]
		racerInput.LapTime = line[4]
		racerInput.AverageLapSpeed = line[5]

		racerInputList = append(racerInputList, racerInput)
	}

	return racerInputList
}
