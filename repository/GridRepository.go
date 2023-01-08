package repository

import (
	"fmt"
	"github.com/TheLe0/f1-race-analysis/infrastructure"
	"github.com/TheLe0/f1-race-analysis/types"
	"github.com/TheLe0/f1-race-analysis/utils"
	"strconv"
)

var localStorage = *infrastructure.GetInstance()
var grid = localStorage.Grid

func FindByRacerNumber(racerNumber string) (types.Racer, bool) {

	for _, racer := range grid {
		if racer.Number == racerNumber {
			return racer, true
		}
	}

	return types.Racer{}, false
}

func DeleteByRacerNumber(racerNumber string) {

	var newGrid []types.Racer

	for _, racer := range grid {
		if racer.Number != racerNumber {
			newGrid = append(newGrid, racer)
		}
	}

	grid = newGrid
}

func insertRacer(racerInput types.RacerInput) {

	var racer types.Racer

	racer.Name = racerInput.Name
	racer.Number = racerInput.Number
	racer.StartTime = utils.ParseTimeStringToMilliseconds(racerInput.LocalTime)
	racer.Laps = utils.ParseStringToUint8(racerInput.LapNumber)
	racer.FastestLap = utils.ParseLapTimeStringToMilliseconds(racerInput.LapTime)
	racer.TotalSpeed = utils.ParseSpeedString(racerInput.AverageLapSpeed)

	grid = append(grid, racer)
}

func updateRacer(racer types.Racer, racerInput types.RacerInput) {

	racer.Laps = utils.ParseStringToUint8(racerInput.LapNumber)
	racer.TotalSpeed += utils.ParseSpeedString(racerInput.AverageLapSpeed)
	racer.AverageSpeed = float32(racer.TotalSpeed) / float32(racer.Laps)

	currentLapTime := utils.ParseLapTimeStringToMilliseconds(racerInput.LapTime)

	if racer.FastestLap >= currentLapTime {
		racer.FastestLap = currentLapTime
	}

	if GetConfigs().RaceLaps == racer.Laps {
		racer.FinishTime = utils.ParseTimeStringToMilliseconds(racerInput.LocalTime)
	}

	DeleteByRacerNumber(racer.Number)
	grid = append(grid, racer)
}

func UpsertRacerOnGrid(racerInput types.RacerInput) {

	racer, alreadyExists := FindByRacerNumber(racerInput.Number)

	if alreadyExists {
		updateRacer(racer, racerInput)
	} else {
		insertRacer(racerInput)
	}
}

func GetAllRacersFormatted() []types.RacerOutput {

	var gridFormatted []types.RacerOutput
	var difference uint32 = 0
	var lastDuration uint32 = 0

	for index, racer := range grid {

		var row types.RacerOutput

		if GetConfigs().RaceLaps == racer.Laps {
			durationTime := racer.FinishTime - racer.StartTime

			if index == 0 {
				difference = 0
				lastDuration = durationTime
			} else {
				difference += durationTime - lastDuration
				lastDuration = durationTime
			}

			row.TotalRacingTime = fmt.Sprintf("%s", utils.ParseMillisecondsToTimeString(durationTime))
			row.Difference = fmt.Sprintf("%s", utils.ParseMillisecondsToTimeString(difference))

		} else {
			row.TotalRacingTime = fmt.Sprintf("Not Completed")
			row.Difference = fmt.Sprintf("Not Completed")
		}

		row.Position = strconv.Itoa(index + 1)
		row.Racer = fmt.Sprintf("%s - %s", racer.Number, racer.Name)
		row.TotalLaps = fmt.Sprintf("%d", racer.Laps)
		row.FastestLap = fmt.Sprintf("%s", utils.ParseMillisecondsToLapTimeString(racer.FastestLap))
		row.AverageSpeed = fmt.Sprintf("%s", utils.ParseSpeedToString(racer.AverageSpeed))

		gridFormatted = append(gridFormatted, row)
	}

	return gridFormatted
}
