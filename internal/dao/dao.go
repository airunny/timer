package dao

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewMySQL,
	NewRedis,
	NewRedisCache,

	NewGlobal,
	NewApplication,
	NewToken,
	NewUser,
	NewTimer,
	NewTask,
)
