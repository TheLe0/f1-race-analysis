package commands

import (
	"github.com/TheLe0/f1-race-analysis/repository"
	"github.com/TheLe0/f1-race-analysis/types"
	"github.com/TheLe0/f1-race-analysis/utils"
)

func loadData() []types.RacerInput {
	var appConfigs = repository.GetConfigs()

	fileParsed := utils.ParseFile(appConfigs.FilePath)
	racersInput := ConvertToRacer(fileParsed)

	return racersInput
}

func setupConfigurations(filePath string, totalRaceLaps uint8) {

	repository.CreateConfiguration(filePath, totalRaceLaps)
}

func AnalyzeRace(filePath string, totalRaceLaps uint8) {

	setupConfigurations(filePath, totalRaceLaps)
	input := loadData()

	for _, line := range input {

		repository.UpsertRacerOnGrid(line)
	}

	ShowResults()
}
