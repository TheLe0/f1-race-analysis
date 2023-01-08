package repository

import "github.com/TheLe0/f1-race-analysis/types"

var configs = localStorage.Configs

func CreateConfiguration(filePath string, totalRaceLaps uint8) {

	configs.SetupConfigs(filePath, totalRaceLaps)
}

func GetConfigs() types.Configuration {
	return configs
}
