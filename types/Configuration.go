package types

type Configuration struct {
	FilePath string
	RaceLaps uint8
}

func (config *Configuration) SetupConfigs(filePath string, raceLaps uint8) {
	config.FilePath = filePath
	config.RaceLaps = raceLaps
}
