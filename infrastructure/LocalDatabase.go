package infrastructure

import (
	"github.com/TheLe0/f1-race-analysis/types"
	"sync"
)

var (
	instance *types.LocalStorage
	once     sync.Once
)

func GetInstance() *types.LocalStorage {
	once.Do(func() {
		instance = &types.LocalStorage{}
	})
	return instance
}
