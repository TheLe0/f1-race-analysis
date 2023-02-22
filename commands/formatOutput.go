package commands

import (
	"fmt"
	"os"

	"github.com/TheLe0/f1-race-analysis/repository"
	"github.com/TheLe0/f1-race-analysis/types"
	"github.com/common-nighthawk/go-figure"
	"github.com/olekukonko/tablewriter"
)

func ShowResults() {
	printTitle()
	printFastestLap()
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

func printFastestLap() {
	fmt.Printf("🥇 Fastest Lap: " + repository.GetFastestLapFormatted() + "\n")
}

func printTitle() {
	myFigure := figure.NewColorFigure("F1 Race Analysis", "", "green", true)
	myFigure.Print()
}
