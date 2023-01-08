// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/gateway/pb/gateway.proto

package gateway

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

type TYPE int32

const (
	// 基于Etcd对接的Traefik网关
	TYPE_TRAEFIK_ETCD TYPE = 0
)

// Enum value maps for TYPE.
var (
	TYPE_name = map[int32]string{
		0: "TRAEFIK_ETCD",
	}
	TYPE_value = map[string]int32{
		"TRAEFIK_ETCD": 0,
	}
)

func (x TYPE) Enum() *TYPE {
	p := new(TYPE)
	*p = x
	return p
}

func (x TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_gateway_pb_gateway_proto_enumTypes[0].Descriptor()
}

func (TYPE) Type() protoreflect.EnumType {
	return &file_apps_gateway_pb_gateway_proto_enumTypes[0]
}

func (x TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TYPE.Descriptor instead.
func (TYPE) EnumDescriptor() ([]byte, []int) {
	return file_apps_gateway_pb_gateway_proto_rawDescGZIP(), []int{0}
}

type Gateway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 网关ID
	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 创建时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 更新时间
	// @gotags: bson:"update_at" json:"update_at"
	UpdateAt int64 `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 网关信息
	// @gotags: bson:"spec" json:"spec"
	Spec *CreateGatewayRequest `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec" bson:"spec"`
}

func (x *Gateway) Reset() {
	*x = Gateway{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_gateway_pb_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gateway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gateway) ProtoMessage() {}

func (x *Gateway) ProtoReflect() protoreflect.Message {
	mi := &file_apps_gateway_pb_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gateway.ProtoReflect.Descriptor instead.
func (*Gateway) Descriptor() ([]byte, []int) {
	return file_apps_gateway_pb_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *Gateway) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Gateway) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Gateway) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Gateway) GetSpec() *CreateGatewayRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

type GatewaySet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*Gateway `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *GatewaySet) Reset() {
	*x = GatewaySet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_gateway_pb_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewaySet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewaySet) ProtoMessage() {}

func (x *GatewaySet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_gateway_pb_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewaySet.ProtoReflect.Descriptor instead.
func (*GatewaySet) Descriptor() ([]byte, []int) {
	return file_apps_gateway_pb_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *GatewaySet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GatewaySet) GetItems() []*Gateway {
	if x != nil {
		return x.Items
	}
	return nil
}

// CreateGatewayRequest 创建网关
type CreateGatewayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 所属域
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 所属空间
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace" bson:"namespace"`
	// 创建者
	// @gotags: bson:"create_by" json:"create_by"
	CreateBy string `protobuf:"bytes,3,opt,name=create_by,json=createBy,proto3" json:"create_by" bson:"create_by"`
	// 是否公开
	// @gotags: bson:"is_public" json:"is_public"
	IsPublic bool `protobuf:"varint,4,opt,name=is_public,json=isPublic,proto3" json:"is_public" bson:"is_public"`
	// 是否是该空间下的默认网关, 一个空间内只允许有1个默认网关
	// @gotags: bson:"is_default" json:"is_default"
	IsDefault bool `protobuf:"varint,5,opt,name=is_default,json=isDefault,proto3" json:"is_default" bson:"is_default"`
	// 网关的名称
	// @gotags: bson:"name" json:"name" validate:"required,lte=30"
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name" bson:"name" validate:"required,lte=30"`
	// 网关的描述
	// @gotags: bson:"description" json:"description" validate:"lte=400"
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description" bson:"description" validate:"lte=400"`
	// 网关类型
	// @gotags: bson:"type" json:"type"
	Type TYPE `protobuf:"varint,8,opt,name=type,proto3,enum=infraboard.mcenter.gateway.TYPE" json:"type" bson:"type"`
	// Traefik网关配置, 通过该配置可以操作网关
	// @gotags: bson:"traefik_config" json:"traefik_config"
	TraefikConfig *TraefikConfig `protobuf:"bytes,9,opt,name=traefik_config,json=traefikConfig,proto3" json:"traefik_config" bson:"traefik_config"`
	// 网关标签
	// @gotags: bson:"labels" json:"labels"
	Labels map[string]string `protobuf:"bytes,15,rep,name=labels,proto3" json:"labels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"labels"`
}

func (x *CreateGatewayRequest) Reset() {
	*x = CreateGatewayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_gateway_pb_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGatewayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGatewayRequest) ProtoMessage() {}

func (x *CreateGatewayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_gateway_pb_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGatewayRequest.ProtoReflect.Descriptor instead.
func (*CreateGatewayRequest) Descriptor() ([]byte, []int) {
	return file_apps_gateway_pb_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *CreateGatewayRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CreateGatewayRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateGatewayRequest) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *CreateGatewayRequest) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

func (x *CreateGatewayRequest) GetIsDefault() bool {
	if x != nil {
		return x.IsDefault
	}
	return false
}

func (x *CreateGatewayRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateGatewayRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateGatewayRequest) GetType() TYPE {
	if x != nil {
		return x.Type
	}
	return TYPE_TRAEFIK_ETCD
}

func (x *CreateGatewayRequest) GetTraefikConfig() *TraefikConfig {
	if x != nil {
		return x.TraefikConfig
	}
	return nil
}

func (x *CreateGatewayRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// Traefik网关配置, 通过操作Etcd来操作Traefik
type TraefikConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// etcd地址
	// @gotags: bson:"endpoints" json:"endpoints"
	Endpoints []string `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints" bson:"endpoints"`
	// 用户名
	// @gotags: bson:"username" json:"username"
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username" bson:"username"`
	// 用户密码
	// @gotags: bson:"password" json:"password"
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password" bson:"password"`
	// 前缀
	// @gotags: bson:"root_key" json:"root_key"
	RootKey string `protobuf:"bytes,4,opt,name=root_key,json=rootKey,proto3" json:"root_key" bson:"root_key"`
	// 开启TLS
	// @gotags: bson:"enable_tls" json:"enable_tls"
	EnableTls bool `protobuf:"varint,5,opt,name=enable_tls,json=enableTls,proto3" json:"enable_tls" bson:"enable_tls"`
}

func (x *TraefikConfig) Reset() {
	*x = TraefikConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_gateway_pb_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraefikConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraefikConfig) ProtoMessage() {}

func (x *TraefikConfig) ProtoReflect() protoreflect.Message {
	mi := &file_apps_gateway_pb_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraefikConfig.ProtoReflect.Descriptor instead.
func (*TraefikConfig) Descriptor() ([]byte, []int) {
	return file_apps_gateway_pb_gateway_proto_rawDescGZIP(), []int{3}
}

func (x *TraefikConfig) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *TraefikConfig) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *TraefikConfig) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *TraefikConfig) GetRootKey() string {
	if x != nil {
		return x.RootKey
	}
	return ""
}

func (x *TraefikConfig) GetEnableTls() bool {
	if x != nil {
		return x.EnableTls
	}
	return false
}

var File_apps_gateway_pb_gateway_proto protoreflect.FileDescriptor

var file_apps_gateway_pb_gateway_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70,
	0x62, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0x99, 0x01, 0x0a, 0x07,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x12, 0x44, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x30, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x5d, 0x0a, 0x0a, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x39, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xf4, 0x03, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x54, 0x59, 0x50, 0x45, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x50, 0x0a, 0x0e, 0x74, 0x72,
	0x61, 0x65, 0x66, 0x69, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x54, 0x72, 0x61, 0x65, 0x66, 0x69, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x74,
	0x72, 0x61, 0x65, 0x66, 0x69, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x54, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9f, 0x01,
	0x0a, 0x0d, 0x54, 0x72, 0x61, 0x65, 0x66, 0x69, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x74, 0x4b, 0x65, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x6c, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x6c, 0x73, 0x2a,
	0x18, 0x0a, 0x04, 0x54, 0x59, 0x50, 0x45, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x52, 0x41, 0x45, 0x46,
	0x49, 0x4b, 0x5f, 0x45, 0x54, 0x43, 0x44, 0x10, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_gateway_pb_gateway_proto_rawDescOnce sync.Once
	file_apps_gateway_pb_gateway_proto_rawDescData = file_apps_gateway_pb_gateway_proto_rawDesc
)

func file_apps_gateway_pb_gateway_proto_rawDescGZIP() []byte {
	file_apps_gateway_pb_gateway_proto_rawDescOnce.Do(func() {
		file_apps_gateway_pb_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_gateway_pb_gateway_proto_rawDescData)
	})
	return file_apps_gateway_pb_gateway_proto_rawDescData
}

var file_apps_gateway_pb_gateway_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_gateway_pb_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_apps_gateway_pb_gateway_proto_goTypes = []interface{}{
	(TYPE)(0),                    // 0: infraboard.mcenter.gateway.TYPE
	(*Gateway)(nil),              // 1: infraboard.mcenter.gateway.Gateway
	(*GatewaySet)(nil),           // 2: infraboard.mcenter.gateway.GatewaySet
	(*CreateGatewayRequest)(nil), // 3: infraboard.mcenter.gateway.CreateGatewayRequest
	(*TraefikConfig)(nil),        // 4: infraboard.mcenter.gateway.TraefikConfig
	nil,                          // 5: infraboard.mcenter.gateway.CreateGatewayRequest.LabelsEntry
}
var file_apps_gateway_pb_gateway_proto_depIdxs = []int32{
	3, // 0: infraboard.mcenter.gateway.Gateway.spec:type_name -> infraboard.mcenter.gateway.CreateGatewayRequest
	1, // 1: infraboard.mcenter.gateway.GatewaySet.items:type_name -> infraboard.mcenter.gateway.Gateway
	0, // 2: infraboard.mcenter.gateway.CreateGatewayRequest.type:type_name -> infraboard.mcenter.gateway.TYPE
	4, // 3: infraboard.mcenter.gateway.CreateGatewayRequest.traefik_config:type_name -> infraboard.mcenter.gateway.TraefikConfig
	5, // 4: infraboard.mcenter.gateway.CreateGatewayRequest.labels:type_name -> infraboard.mcenter.gateway.CreateGatewayRequest.LabelsEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_apps_gateway_pb_gateway_proto_init() }
func file_apps_gateway_pb_gateway_proto_init() {
	if File_apps_gateway_pb_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_gateway_pb_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gateway); i {
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
		file_apps_gateway_pb_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewaySet); i {
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
		file_apps_gateway_pb_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGatewayRequest); i {
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
		file_apps_gateway_pb_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraefikConfig); i {
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
			RawDescriptor: file_apps_gateway_pb_gateway_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_gateway_pb_gateway_proto_goTypes,
		DependencyIndexes: file_apps_gateway_pb_gateway_proto_depIdxs,
		EnumInfos:         file_apps_gateway_pb_gateway_proto_enumTypes,
		MessageInfos:      file_apps_gateway_pb_gateway_proto_msgTypes,
	}.Build()
	File_apps_gateway_pb_gateway_proto = out.File
	file_apps_gateway_pb_gateway_proto_rawDesc = nil
	file_apps_gateway_pb_gateway_proto_goTypes = nil
	file_apps_gateway_pb_gateway_proto_depIdxs = nil
}
