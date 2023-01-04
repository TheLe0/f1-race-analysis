package main

import (
	"github.com/TheLe0/f1-race-analysis/commands"
	"github.com/TheLe0/f1-race-analysis/utils"
)

func main() {

	fileParsed := utils.ParseFile("./data/log.csv")
	racersInput := commands.ConvertToRacer(fileParsed)
	commands.AnalyzeRace(racersInput)
}
