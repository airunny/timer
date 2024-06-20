package dao

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewMySQL,
	NewRedis,
	NewRedisCache,
	NewUser,
	NewApplication,
	NewTimerRecord,
)
