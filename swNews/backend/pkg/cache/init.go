package cache

import (
	"backend/pkg/logger"
)

var Store Driver = NewLocalStore()

func Init() {

	err := Store.Ping()
	if err != nil {
		logger.Panic(err)
	}
	if err = Store.Restore(DefaultCacheFile); err != nil {
		logger.Warn(err)
	}
}
