package dao

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/airunny/timer/api/common"
	v1 "github.com/airunny/timer/api/timer/v1"
	"github.com/airunny/timer/internal/models"
	"github.com/airunny/wiki-go-tools/objectid"
	"github.com/airunny/wiki-go-tools/ormhelper"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
)

func TestNewTimer(t *testing.T) {
	db, closer, err := NewMySQL(&common.DataConfig{
		Database: &common.DataConfig_Database{
			Source: "root:yann@tcp(127.0.0.1:3306)/wikistock?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&loc=Local",
		},
	}, log.GetLogger())
	assert.Nil(t, err)
	defer closer()

	var (
		now     = time.Now().Local()
		cli     = NewTimer(db)
		id      = objectid.ObjectID()
		ctx     = context.Background()
		newData = &models.Timer{
			ID:              id,
			AppId:           "1",
			Name:            "1",
			Type:            v1.TimerType_CRONJOB,
			CallbackType:    v1.CallbackType_HTTP,
			CallbackAddress: "www.baidu.com",
			CronExpr:        "* * * * *",
			ExpireAt:        now.Unix(),
			Attach:          "attach",
			Status:          v1.TimerStatus_ENABLE,
			Fail:            0,
			Success:         1,
			Extra:           &models.TimerExtra{},
			CreatedAt:       now,
			UpdatedAt:       now,
		}
	)

	// add
	err = cli.Add(ctx, newData)
	assert.Nil(t, err)

	// get
	data, err := cli.Get(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, newData.ID, data.ID)
	assert.Equal(t, newData.AppId, data.AppId)
	assert.Equal(t, newData.Name, data.Name)
	assert.Equal(t, newData.Type, data.Type)
	assert.Equal(t, newData.CallbackType, data.CallbackType)
	assert.Equal(t, newData.CallbackAddress, data.CallbackAddress)
	assert.Equal(t, newData.CronExpr, data.CronExpr)
	assert.Equal(t, newData.ExpireAt, data.ExpireAt)
	assert.Equal(t, newData.Attach, data.Attach)
	assert.Equal(t, newData.Status, data.Status)
	assert.Equal(t, newData.Fail, data.Fail)
	assert.Equal(t, newData.Success, data.Success)
	assert.True(t, reflect.DeepEqual(newData.Extra, data.Extra))

	// update
	newData = &models.Timer{
		ID:              id,
		AppId:           "11",
		Name:            "11",
		Type:            v1.TimerType_ONCE,
		CallbackType:    v1.CallbackType_KAFKA,
		CallbackAddress: "http://www.baidu.com",
		CronExpr:        "0-30/12 * * * *",
		ExpireAt:        now.Unix() + 10,
		Attach:          "attach1",
		Status:          v1.TimerStatus_DISABLE,
		Fail:            1,
		Success:         2,
		Extra: &models.TimerExtra{
			Params: []*v1.Params{
				{
					Key:   "key",
					Value: "value",
					In:    v1.ParamsIn_HEADER,
				},
			},
			FailReason: []string{"1", "2"},
		},
		CreatedAt: now,
		UpdatedAt: now,
	}
	err = cli.Update(ctx, newData)
	assert.Nil(t, err)

	// get
	data, err = cli.Get(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, newData.ID, data.ID)
	assert.Equal(t, newData.AppId, data.AppId)
	assert.Equal(t, newData.Name, data.Name)
	assert.Equal(t, newData.Type, data.Type)
	assert.Equal(t, newData.CallbackType, data.CallbackType)
	assert.Equal(t, newData.CallbackAddress, data.CallbackAddress)
	assert.Equal(t, newData.CronExpr, data.CronExpr)
	assert.Equal(t, newData.ExpireAt, data.ExpireAt)
	assert.Equal(t, newData.Attach, data.Attach)
	assert.Equal(t, newData.Status, data.Status)
	assert.Equal(t, newData.Fail, data.Fail)
	assert.Equal(t, newData.Success, data.Success)
	assert.True(t, reflect.DeepEqual(newData.Extra, data.Extra))

	// update status
	err = cli.UpdateStatus(ctx, id, v1.TimerStatus_ENABLE)
	assert.Nil(t, err)

	// get
	data, err = cli.Get(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, newData.ID, data.ID)
	assert.Equal(t, newData.AppId, data.AppId)
	assert.Equal(t, newData.Name, data.Name)
	assert.Equal(t, newData.Type, data.Type)
	assert.Equal(t, newData.CallbackType, data.CallbackType)
	assert.Equal(t, newData.CallbackAddress, data.CallbackAddress)
	assert.Equal(t, newData.CronExpr, data.CronExpr)
	assert.Equal(t, newData.ExpireAt, data.ExpireAt)
	assert.Equal(t, newData.Attach, data.Attach)
	assert.Equal(t, v1.TimerStatus_ENABLE, data.Status)
	assert.Equal(t, newData.Fail, data.Fail)
	assert.Equal(t, newData.Success, data.Success)
	assert.True(t, reflect.DeepEqual(newData.Extra, data.Extra))

	// delete
	err = cli.Delete(ctx, id)
	assert.Nil(t, err)

	// get
	_, err = cli.Get(ctx, id)
	assert.Equal(t, ormhelper.ErrNotFound, err)

	// find
	values, _, err := cli.Find(ctx, 1, "")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(values))
}
