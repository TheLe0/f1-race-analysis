package commands

import "github.com/TheLe0/f1-race-analysis/types"

func AnalyzeRace(input []types.RacerInput) {

	var grid []types.Racer

	for _, line := range input {

		var racer types.Racer

		racer.Name = line.Name
		racer.Number = line.Number

		grid = append(grid, racer)
	}
}
