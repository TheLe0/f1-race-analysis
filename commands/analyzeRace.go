package commands

import (
	"github.com/TheLe0/f1-race-analysis/configuration"
	"github.com/TheLe0/f1-race-analysis/repository"
	"github.com/TheLe0/f1-race-analysis/types"
	"github.com/TheLe0/f1-race-analysis/utils"
)

var appConfigs = *configuration.GetConfiguration()

func loadData() []types.RacerInput {
	fileParsed := utils.ParseFile(appConfigs.FilePath)
	racersInput := ConvertToRacer(fileParsed)

	return racersInput
}

func setupConfigurations(filePath string, totalRaceLaps uint8) {

	appConfigs.FilePath = filePath
	appConfigs.RaceLaps = totalRaceLaps
}

func AnalyzeRace(filePath string, totalRaceLaps uint8) {

	setupConfigurations(filePath, totalRaceLaps)
	input := loadData()

	for _, line := range input {

		repository.UpsertRacerOnGrid(line)
	}

	ShowResults()
}
