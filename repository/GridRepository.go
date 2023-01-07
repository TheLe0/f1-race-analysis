package repository

import (
	"github.com/TheLe0/f1-race-analysis/configuration"
	"github.com/TheLe0/f1-race-analysis/infrastructure"
	"github.com/TheLe0/f1-race-analysis/types"
	"github.com/TheLe0/f1-race-analysis/utils"
)

var localStorage = *infrastructure.GetInstance()
var grid = localStorage.Grid
var appConfigs = *configuration.GetConfiguration()

func FindByRacerNumber(racerNumber string) *types.Racer {

	var foundRacer *types.Racer = nil

	for _, racer := range grid {
		if racer.Number == racerNumber {
			foundRacer = &racer
		}
	}

	return foundRacer
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

	currentLapTime := utils.ParseLapTimeStringToMilliseconds(racerInput.LapTime)

	if racer.FastestLap <= currentLapTime {
		racer.FastestLap = currentLapTime
	}

	if appConfigs.RaceLaps == racer.Laps {
		racer.FinishTime = utils.ParseTimeStringToMilliseconds(racerInput.LocalTime)
		racer.AverageSpeed = float32(racer.TotalSpeed) / float32(racer.Laps)
	}
}

func UpsertRacerOnGrid(racerInput types.RacerInput) {

	racer := FindByRacerNumber(racerInput.Number)

	if racer != nil {
		updateRacer(*racer, racerInput)
	} else {
		insertRacer(racerInput)
	}
}
