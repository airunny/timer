package dao

import (
	"context"
	"testing"

	"github.com/airunny/wiki-go-tools/objectid"
	"github.com/airunny/wiki-go-tools/ormhelper"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"timer/api/common"
	"timer/internal/models"
)

func TestNewApplication(t *testing.T) {
	db, closer, err := NewMySQL(&common.DataConfig{
		Database: &common.DataConfig_Database{
			Source: "root:yann@tcp(127.0.0.1:3306)/wikistock?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&loc=Local",
		},
	}, log.GetLogger())
	assert.Nil(t, err)
	defer closer()

	application := NewApplication(db)

	var (
		id      = objectid.ObjectID()
		newData = &models.Application{
			ID:          id,
			Name:        "timer",
			Description: "Distributed timer service",
		}
	)

	// add
	err = application.Add(context.Background(), newData)
	assert.Nil(t, err)

	// get
	data, err := application.Get(context.Background(), id)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, newData.Name, data.Name)
	assert.Equal(t, newData.Description, data.Description)

	// update
	updateData := &models.Application{
		ID:          id,
		Name:        "new timer",
		Description: "new Distributed timer service",
	}
	err = application.Update(context.Background(), updateData)
	assert.Nil(t, err)

	// get
	data, err = application.Get(context.Background(), id)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, updateData.Name, data.Name)
	assert.Equal(t, updateData.Description, data.Description)

	// find
	datas, _, err := application.Find(context.Background(), 10, "")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(datas))

	// delete
	err = application.Delete(context.Background(), id)
	assert.Nil(t, err)

	// delete again
	err = application.Delete(context.Background(), id)
	assert.Nil(t, err)

	// get
	data, err = application.Get(context.Background(), id)
	assert.Equal(t, ormhelper.ErrNotFound, err)
}
