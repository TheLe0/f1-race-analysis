package commands

import (
	"github.com/TheLe0/f1-race-analysis/repository"
	"github.com/TheLe0/f1-race-analysis/types"
	"github.com/common-nighthawk/go-figure"
	"github.com/olekukonko/tablewriter"
	"os"
)

func ShowResults() {
	printTitle()

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{
		"Position",
		"Racer",
		"Completed Laps",
		"Fastest Lap",
		"Average Speed",
		"Total time on race",
		"Difference"})
	
	for _, racerOutput := range repository.GetAllRacersFormatted() {

		table.Append(formatRacerOutputToString(racerOutput))
	}

	table.Render()
}

func formatRacerOutputToString(racerOutput types.RacerOutput) []string {
	return []string{
		racerOutput.Position,
		racerOutput.Racer,
		racerOutput.TotalLaps,
		racerOutput.FastestLap,
		racerOutput.AverageSpeed,
		racerOutput.TotalRacingTime,
		racerOutput.Difference}
}

func printTitle() {
	myFigure := figure.NewColorFigure("F1 Race Analysis", "", "green", true)
	myFigure.Print()
}
