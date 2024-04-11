package bootstrap

import (
	"backend/pkg/cache"
	"backend/pkg/conf"
	"backend/pkg/logger"
)

func Init(confPath ...string) {
	InitApplication()
	conf.Init(confPath...)
	logger.Init()
	cache.Init()

}

func Shutdown() error {
	return cache.Store.Close()
}
