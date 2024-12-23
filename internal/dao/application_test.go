package dao

import (
	"context"
	"testing"

	"github.com/airunny/timer/api/common"
	v1 "github.com/airunny/timer/api/timer/v1"
	"github.com/airunny/timer/internal/models"
	"github.com/airunny/wiki-go-tools/objectid"
	"github.com/airunny/wiki-go-tools/ormhelper"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
)

func TestNewApplication(t *testing.T) {
	db, closer, err := NewMySQL(&common.DataConfig{
		Database: &common.DataConfig_Database{
			Source: "root:yann@tcp(127.0.0.1:3306)/wikistock?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&loc=Local",
		},
	}, log.GetLogger())
	assert.Nil(t, err)
	defer closer()

	var (
		application = NewApplication(db)
		id          = objectid.ObjectID()
		newData     = &models.Application{
			ID:            id,
			Name:          "timer",
			Description:   "Distributed timer service",
			Secret:        "secret",
			Authorization: true,
			Status:        v1.ApplicationStatus_ON,
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
	assert.Equal(t, newData.Secret, data.Secret)
	assert.Equal(t, newData.Authorization, data.Authorization)
	assert.Equal(t, newData.Status, data.Status)

	// update
	updateData := &models.Application{
		ID:            id,
		Name:          "new timer",
		Description:   "new Distributed timer service",
		Authorization: false,
		Status:        v1.ApplicationStatus_OFF,
	}
	err = application.Update(context.Background(), updateData)
	assert.Nil(t, err)

	// get
	data, err = application.Get(context.Background(), id)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, updateData.Name, data.Name)
	assert.Equal(t, updateData.Description, data.Description)
	assert.Equal(t, newData.Secret, data.Secret)
	assert.Equal(t, updateData.Authorization, data.Authorization)
	assert.Equal(t, updateData.Status, data.Status)

	// update status
	err = application.UpdateStatus(context.Background(), id, v1.ApplicationStatus_ON)
	assert.Nil(t, err)

	// get
	data, err = application.Get(context.Background(), id)
	assert.Nil(t, err)
	assert.Equal(t, v1.ApplicationStatus_ON, data.Status)

	// update secret
	err = application.UpdateSecret(context.Background(), id, "timer_secret")
	assert.Nil(t, err)

	// get
	data, err = application.Get(context.Background(), id)
	assert.Nil(t, err)
	assert.Equal(t, "timer_secret", data.Secret)

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
