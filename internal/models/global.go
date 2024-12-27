package models

import (
	"encoding/json"
	"time"
)

const (
	NamespaceMetadata string = "metadata" // 元数据
	NamespaceBucket   string = "bucket"   // bucket信息
)

type Global struct {
	ID        int64     `gorm:"column:id;type:bigint unsigned not null primary key auto_increment;comment:主键"`
	Namespace string    `gorm:"column:namespace;type:varchar(64) not null;default:'';index:idx_namespace_name,unique;comment:namespace"`
	Name      string    `gorm:"column:name;type:varchar(64) not null;default:'';index:idx_namespace_name,unique;comment:key名称"`
	Value     string    `gorm:"column:value;type:text;comment:value"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp not null;default:current_timestamp();comment:创建时间"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp not null default current_timestamp() on update current_timestamp();comment:最后一次更新时间"`
}

func (Global) TableName() string {
	return "global"
}

type MetadataBucket struct {
	NumberOfBuckets int64    `json:"number_of_buckets"` // bucket的数量
	Buckets         []string `json:"buckets"`           // bucket列表
}

func (m *MetadataBucket) String() string {
	str, _ := json.Marshal(m)
	return string(str)
}

func (m *MetadataBucket) FromString(in string) error {
	return json.Unmarshal([]byte(in), m)
}
