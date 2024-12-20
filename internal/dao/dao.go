package dao

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewMySQL,
	NewRedis,
	NewRedisCache,
	NewToken,
	NewUser,
	NewApplication,
	NewTimer,
	NewTask,
)
