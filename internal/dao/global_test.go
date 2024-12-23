package dao

import (
	"context"
	"testing"

	"github.com/airunny/timer/api/common"
	"github.com/airunny/timer/internal/models"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
)

func TestNewGlobal(t *testing.T) {
	db, closer, err := NewMySQL(&common.DataConfig{
		Database: &common.DataConfig_Database{
			Source: "root:yann@tcp(127.0.0.1:3306)/wikistock?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&loc=Local",
		},
	}, log.GetLogger())
	assert.Nil(t, err)
	defer closer()

	var (
		cli     = NewGlobal(db)
		ctx     = context.Background()
		newData = &models.Global{
			Namespace: "timer",
			Name:      "name",
			Value:     "value",
		}
	)

	newId, err := cli.Add(ctx, newData)
	assert.Nil(t, err)
	assert.Greater(t, newId, int64(0))

	data, err := cli.GetByName(ctx, newData.Namespace, newData.Name)
	assert.Nil(t, err)
	assert.Equal(t, newData.Namespace, data.Namespace)
	assert.Equal(t, newData.Name, data.Name)
	assert.Equal(t, newData.Value, data.Value)

	err = cli.UpdateByKey(ctx, &models.Global{
		Namespace: "timer",
		Name:      "name",
		Value:     "value1",
	})
	assert.Nil(t, err)

	data, err = cli.GetByName(ctx, newData.Namespace, newData.Name)
	assert.Nil(t, err)
	assert.Equal(t, newData.Namespace, data.Namespace)
	assert.Equal(t, newData.Name, data.Name)
	assert.Equal(t, "value1", data.Value)

	err = cli.DeleteByKey(ctx, "timer", "name")
	assert.Nil(t, err)

	_, err = cli.GetByName(ctx, "timer", "name")

}
