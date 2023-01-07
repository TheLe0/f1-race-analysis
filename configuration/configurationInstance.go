package configuration

import (
	"github.com/TheLe0/f1-race-analysis/types"
	"sync"
)

var (
	instance *types.Configuration
	once     sync.Once
)

func GetConfiguration() *types.Configuration {
	once.Do(func() {
		instance = &types.Configuration{}
	})
	return instance
}
