package dao

import (
	"context"
	"testing"
	"time"

	"github.com/airunny/timer/api/common"
	v1 "github.com/airunny/timer/api/timer/v1"
	"github.com/airunny/timer/internal/models"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	db, closer, err := NewMySQL(&common.DataConfig{
		Database: &common.DataConfig_Database{
			Source: "root:yann@tcp(127.0.0.1:3306)/wikistock?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&loc=Local",
		},
	}, log.GetLogger())
	assert.Nil(t, err)
	defer closer()

	var (
		now     = time.Now().Local()
		cli     = NewTask(db)
		ctx     = context.Background()
		newData = &models.Task{
			TimerId:      "001",
			RunTime:      now,
			Request:      "request",
			Response:     "response",
			Status:       v1.TaskStatus_Success,
			Retry:        3,
			FailedReason: "reason",
			CreatedAt:    now,
			UpdatedAt:    now,
		}
	)
	newId, err := cli.Add(ctx, newData)
	assert.Nil(t, err)
	assert.Greater(t, newId, int64(0))

	values, _, err := cli.FindByTimerId(ctx, "001", "", 10, -1)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(values))

	values, _, err = cli.FindByTime(ctx, "", now.Add(-time.Hour), now.Add(time.Hour), 10, -1)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(values))
}
