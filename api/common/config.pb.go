// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: common/config.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ServerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http      *ServerConfig_HTTP  `protobuf:"bytes,1,opt,name=http,json=http,proto3" json:"http"`                  // http信息
	Grpc      *ServerConfig_GRPC  `protobuf:"bytes,2,opt,name=grpc,json=grpc,proto3" json:"grpc"`                  // grpc信息
	Trace     *ServerConfig_Trace `protobuf:"bytes,3,opt,name=trace,json=trace,proto3" json:"trace"`               // 链路追踪信息
	OssDomain string              `protobuf:"bytes,4,opt,name=oss_domain,json=ossDomain,proto3" json:"oss_domain"` // 对象存储domain
}

func (x *ServerConfig) Reset() {
	*x = ServerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfig) ProtoMessage() {}

func (x *ServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_common_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfig.ProtoReflect.Descriptor instead.
func (*ServerConfig) Descriptor() ([]byte, []int) {
	return file_common_config_proto_rawDescGZIP(), []int{0}
}

func (x *ServerConfig) GetHttp() *ServerConfig_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *ServerConfig) GetGrpc() *ServerConfig_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

func (x *ServerConfig) GetTrace() *ServerConfig_Trace {
	if x != nil {
		return x.Trace
	}
	return nil
}

func (x *ServerConfig) GetOssDomain() string {
	if x != nil {
		return x.OssDomain
	}
	return ""
}

type DataConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database *DataConfig_Database `protobuf:"bytes,1,opt,name=database,json=database,proto3" json:"database"`
	Redis    *DataConfig_Redis    `protobuf:"bytes,2,opt,name=redis,json=redis,proto3" json:"redis"`
}

func (x *DataConfig) Reset() {
	*x = DataConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataConfig) ProtoMessage() {}

func (x *DataConfig) ProtoReflect() protoreflect.Message {
	mi := &file_common_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataConfig.ProtoReflect.Descriptor instead.
func (*DataConfig) Descriptor() ([]byte, []int) {
	return file_common_config_proto_rawDescGZIP(), []int{1}
}

func (x *DataConfig) GetDatabase() *DataConfig_Database {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *DataConfig) GetRedis() *DataConfig_Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

type GRPCClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint       string `protobuf:"bytes,1,opt,name=endpoint,json=endpoint,proto3" json:"endpoint"`
	TimeoutSeconds int64  `protobuf:"varint,2,opt,name=timeout_seconds,json=timeoutSeconds,proto3" json:"timeout_seconds"`
}

func (x *GRPCClient) Reset() {
	*x = GRPCClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GRPCClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GRPCClient) ProtoMessage() {}

func (x *GRPCClient) ProtoReflect() protoreflect.Message {
	mi := &file_common_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GRPCClient.ProtoReflect.Descriptor instead.
func (*GRPCClient) Descriptor() ([]byte, []int) {
	return file_common_config_proto_rawDescGZIP(), []int{2}
}

func (x *GRPCClient) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *GRPCClient) GetTimeoutSeconds() int64 {
	if x != nil {
		return x.TimeoutSeconds
	}
	return 0
}

type ServerConfig_HTTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network        string `protobuf:"bytes,1,opt,name=network,json=network,proto3" json:"network"`
	Addr           string `protobuf:"bytes,2,opt,name=addr,json=addr,proto3" json:"addr"`
	TimeoutSeconds int64  `protobuf:"varint,3,opt,name=timeout_seconds,json=timeoutSeconds,proto3" json:"timeout_seconds"`
}

func (x *ServerConfig_HTTP) Reset() {
	*x = ServerConfig_HTTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfig_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfig_HTTP) ProtoMessage() {}

func (x *ServerConfig_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_common_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfig_HTTP.ProtoReflect.Descriptor instead.
func (*ServerConfig_HTTP) Descriptor() ([]byte, []int) {
	return file_common_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ServerConfig_HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *ServerConfig_HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *ServerConfig_HTTP) GetTimeoutSeconds() int64 {
	if x != nil {
		return x.TimeoutSeconds
	}
	return 0
}

type ServerConfig_GRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network        string `protobuf:"bytes,1,opt,name=network,json=network,proto3" json:"network"`
	Addr           string `protobuf:"bytes,2,opt,name=addr,json=addr,proto3" json:"addr"`
	TimeoutSeconds int64  `protobuf:"varint,3,opt,name=timeout_seconds,json=timeoutSeconds,proto3" json:"timeout_seconds"`
}

func (x *ServerConfig_GRPC) Reset() {
	*x = ServerConfig_GRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfig_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfig_GRPC) ProtoMessage() {}

func (x *ServerConfig_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_common_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfig_GRPC.ProtoReflect.Descriptor instead.
func (*ServerConfig_GRPC) Descriptor() ([]byte, []int) {
	return file_common_config_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ServerConfig_GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *ServerConfig_GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *ServerConfig_GRPC) GetTimeoutSeconds() int64 {
	if x != nil {
		return x.TimeoutSeconds
	}
	return 0
}

type ServerConfig_Trace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind     string `protobuf:"bytes,1,opt,name=kind,json=kind,proto3" json:"kind"`
	Endpoint string `protobuf:"bytes,2,opt,name=endpoint,json=endpoint,proto3" json:"endpoint"`
	Fraction string `protobuf:"bytes,3,opt,name=fraction,json=fraction,proto3" json:"fraction"`
}

func (x *ServerConfig_Trace) Reset() {
	*x = ServerConfig_Trace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfig_Trace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfig_Trace) ProtoMessage() {}

func (x *ServerConfig_Trace) ProtoReflect() protoreflect.Message {
	mi := &file_common_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfig_Trace.ProtoReflect.Descriptor instead.
func (*ServerConfig_Trace) Descriptor() ([]byte, []int) {
	return file_common_config_proto_rawDescGZIP(), []int{0, 2}
}

func (x *ServerConfig_Trace) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ServerConfig_Trace) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *ServerConfig_Trace) GetFraction() string {
	if x != nil {
		return x.Fraction
	}
	return ""
}

type DataConfig_Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver             string `protobuf:"bytes,1,opt,name=driver,json=driver,proto3" json:"driver"`
	Source             string `protobuf:"bytes,2,opt,name=source,json=source,proto3" json:"source"`
	Level              int32  `protobuf:"varint,3,opt,name=level,json=level,proto3" json:"level"`
	MaxOpen            int32  `protobuf:"varint,4,opt,name=max_open,json=maxOpen,proto3" json:"max_open"`
	MaxIdle            int32  `protobuf:"varint,5,opt,name=max_idle,json=maxIdle,proto3" json:"max_idle"`
	MaxLifeTimeSeconds int32  `protobuf:"varint,6,opt,name=max_life_time_seconds,json=maxLifeTimeSeconds,proto3" json:"max_life_time_seconds"`
}

func (x *DataConfig_Database) Reset() {
	*x = DataConfig_Database{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataConfig_Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataConfig_Database) ProtoMessage() {}

func (x *DataConfig_Database) ProtoReflect() protoreflect.Message {
	mi := &file_common_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataConfig_Database.ProtoReflect.Descriptor instead.
func (*DataConfig_Database) Descriptor() ([]byte, []int) {
	return file_common_config_proto_rawDescGZIP(), []int{1, 0}
}

func (x *DataConfig_Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *DataConfig_Database) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *DataConfig_Database) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *DataConfig_Database) GetMaxOpen() int32 {
	if x != nil {
		return x.MaxOpen
	}
	return 0
}

func (x *DataConfig_Database) GetMaxIdle() int32 {
	if x != nil {
		return x.MaxIdle
	}
	return 0
}

func (x *DataConfig_Database) GetMaxLifeTimeSeconds() int32 {
	if x != nil {
		return x.MaxLifeTimeSeconds
	}
	return 0
}

type DataConfig_Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address             string `protobuf:"bytes,1,opt,name=address,json=address,proto3" json:"address"`
	Password            string `protobuf:"bytes,2,opt,name=password,json=password,proto3" json:"password"`
	Db                  int32  `protobuf:"varint,3,opt,name=db,json=db,proto3" json:"db"`
	MaxIdle             int32  `protobuf:"varint,4,opt,name=max_idle,json=maxIdle,proto3" json:"max_idle"`
	ReadTimeoutSeconds  int64  `protobuf:"varint,5,opt,name=read_timeout_seconds,json=readTimeoutSeconds,proto3" json:"read_timeout_seconds"`
	WriteTimeoutSeconds int64  `protobuf:"varint,6,opt,name=write_timeout_seconds,json=writeTimeoutSeconds,proto3" json:"write_timeout_seconds"`
}

func (x *DataConfig_Redis) Reset() {
	*x = DataConfig_Redis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataConfig_Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataConfig_Redis) ProtoMessage() {}

func (x *DataConfig_Redis) ProtoReflect() protoreflect.Message {
	mi := &file_common_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataConfig_Redis.ProtoReflect.Descriptor instead.
func (*DataConfig_Redis) Descriptor() ([]byte, []int) {
	return file_common_config_proto_rawDescGZIP(), []int{1, 1}
}

func (x *DataConfig_Redis) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DataConfig_Redis) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *DataConfig_Redis) GetDb() int32 {
	if x != nil {
		return x.Db
	}
	return 0
}

func (x *DataConfig_Redis) GetMaxIdle() int32 {
	if x != nil {
		return x.MaxIdle
	}
	return 0
}

func (x *DataConfig_Redis) GetReadTimeoutSeconds() int64 {
	if x != nil {
		return x.ReadTimeoutSeconds
	}
	return 0
}

func (x *DataConfig_Redis) GetWriteTimeoutSeconds() int64 {
	if x != nil {
		return x.WriteTimeoutSeconds
	}
	return 0
}

var File_common_config_proto protoreflect.FileDescriptor

var file_common_config_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0xd0, 0x03,
	0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2d,
	0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x2d, 0x0a,
	0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x12, 0x30, 0x0a, 0x05,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6f, 0x73, 0x73, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6f, 0x73, 0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x5d, 0x0a,
	0x04, 0x48, 0x54, 0x54, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x1a, 0x5d, 0x0a, 0x04,
	0x47, 0x52, 0x50, 0x43, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x1a, 0x53, 0x0a, 0x05, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x82, 0x04, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x37, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x72, 0x65, 0x64, 0x69,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x65, 0x64, 0x69,
	0x73, 0x52, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x1a, 0xb9, 0x01, 0x0a, 0x08, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x61, 0x78, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d,
	0x61, 0x78, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x64,
	0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x49, 0x64, 0x6c,
	0x65, 0x12, 0x31, 0x0a, 0x15, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x12, 0x6d, 0x61, 0x78, 0x4c, 0x69, 0x66, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x1a, 0xce, 0x01, 0x0a, 0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x64, 0x62, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x64, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x49, 0x64, 0x6c, 0x65, 0x12,
	0x30, 0x0a, 0x14, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x72,
	0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x12, 0x32, 0x0a, 0x15, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x13, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x51, 0x0a, 0x0a, 0x47, 0x52, 0x50, 0x43, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x27, 0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x42, 0x13, 0x5a, 0x11, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_config_proto_rawDescOnce sync.Once
	file_common_config_proto_rawDescData = file_common_config_proto_rawDesc
)

func file_common_config_proto_rawDescGZIP() []byte {
	file_common_config_proto_rawDescOnce.Do(func() {
		file_common_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_config_proto_rawDescData)
	})
	return file_common_config_proto_rawDescData
}

var file_common_config_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_common_config_proto_goTypes = []interface{}{
	(*ServerConfig)(nil),        // 0: common.ServerConfig
	(*DataConfig)(nil),          // 1: common.DataConfig
	(*GRPCClient)(nil),          // 2: common.GRPCClient
	(*ServerConfig_HTTP)(nil),   // 3: common.ServerConfig.HTTP
	(*ServerConfig_GRPC)(nil),   // 4: common.ServerConfig.GRPC
	(*ServerConfig_Trace)(nil),  // 5: common.ServerConfig.Trace
	(*DataConfig_Database)(nil), // 6: common.DataConfig.Database
	(*DataConfig_Redis)(nil),    // 7: common.DataConfig.Redis
}
var file_common_config_proto_depIdxs = []int32{
	3, // 0: common.ServerConfig.http:type_name -> common.ServerConfig.HTTP
	4, // 1: common.ServerConfig.grpc:type_name -> common.ServerConfig.GRPC
	5, // 2: common.ServerConfig.trace:type_name -> common.ServerConfig.Trace
	6, // 3: common.DataConfig.database:type_name -> common.DataConfig.Database
	7, // 4: common.DataConfig.redis:type_name -> common.DataConfig.Redis
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_common_config_proto_init() }
func file_common_config_proto_init() {
	if File_common_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GRPCClient); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfig_HTTP); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfig_GRPC); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfig_Trace); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataConfig_Database); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataConfig_Redis); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_config_proto_goTypes,
		DependencyIndexes: file_common_config_proto_depIdxs,
		MessageInfos:      file_common_config_proto_msgTypes,
	}.Build()
	File_common_config_proto = out.File
	file_common_config_proto_rawDesc = nil
	file_common_config_proto_goTypes = nil
	file_common_config_proto_depIdxs = nil
}
