// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/endpoint/pb/endpoint.proto

package endpoint

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

// Endpoint Service's features
type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 端点名称
	// @gotags: bson:"_id" json:"id" validate:"required,lte=64"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id" validate:"required,lte=64"`
	// 创建时间
	// @gotags: bson:"create_at" json:"create_at,omitempty"
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty" bson:"create_at"`
	// 更新时间
	// @gotags: bson:"update_at" json:"update_at,omitempty"
	UpdateAt int64 `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty" bson:"update_at"`
	// 该功能属于那个服务
	// @gotags: bson:"service_id" json:"service_id,omitempty" validate:"required,lte=64"
	ServiceId string `protobuf:"bytes,4,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty" bson:"service_id" validate:"required,lte=64"`
	// 服务那个版本的功能
	// @gotags: bson:"version" json:"version,omitempty" validate:"required,lte=64"
	Version string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty" bson:"version" validate:"required,lte=64"`
	// 路由条目信息
	// @gotags: bson:"entry" json:"entry" validate:"required"
	Entry *Entry `protobuf:"bytes,6,opt,name=entry,proto3" json:"entry" bson:"entry" validate:"required"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_endpoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_endpoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_endpoint_proto_rawDescGZIP(), []int{0}
}

func (x *Endpoint) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Endpoint) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Endpoint) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Endpoint) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *Endpoint) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Endpoint) GetEntry() *Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

// Entry 路由条目
type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 函数名称
	// @gotags: bson:"function_name" json:"function_name"
	FunctionName string `protobuf:"bytes,2,opt,name=function_name,json=functionName,proto3" json:"function_name" bson:"function_name"`
	// HTTP path 用于自动生成http api
	// @gotags: bson:"path" json:"path"
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path" bson:"path"`
	// HTTP method 用于自动生成http api
	// @gotags: bson:"method" json:"method"
	Method string `protobuf:"bytes,4,opt,name=method,proto3" json:"method" bson:"method"`
	// 资源名称
	// @gotags: bson:"resource" json:"resource"
	Resource string `protobuf:"bytes,5,opt,name=resource,proto3" json:"resource" bson:"resource"`
	// 是否校验用户身份 (acccess_token 校验)
	// @gotags: bson:"auth_enable" json:"auth_enable"
	AuthEnable bool `protobuf:"varint,6,opt,name=auth_enable,json=authEnable,proto3" json:"auth_enable" bson:"auth_enable"`
	// 验证码校验(开启双因子认证需要) (code 校验)
	// @gotags: bson:"code_enable" json:"code_enable"
	CodeEnable bool `protobuf:"varint,13,opt,name=code_enable,json=codeEnable,proto3" json:"code_enable" bson:"code_enable"`
	// 是否校验用户权限
	// @gotags: bson:"permission_enable" json:"permission_enable"
	PermissionEnable bool `protobuf:"varint,7,opt,name=permission_enable,json=permissionEnable,proto3" json:"permission_enable" bson:"permission_enable"`
	// 允许的通过的身份标识符, 比如角色, 用户类型之类
	// @gotags: bson:"allow" json:"allow"
	Allow []string `protobuf:"bytes,12,rep,name=allow,proto3" json:"allow" bson:"allow"`
	// 是否开启操作审计, 开启后这次操作将被记录
	// @gotags: bson:"audit_log" json:"audit_log"
	AuditLog bool `protobuf:"varint,9,opt,name=audit_log,json=auditLog,proto3" json:"audit_log" bson:"audit_log"`
	// 名称空间不能为空
	// @gotags: bson:"required_namespace" json:"required_namespace"
	RequiredNamespace bool `protobuf:"varint,10,opt,name=required_namespace,json=requiredNamespace,proto3" json:"required_namespace" bson:"required_namespace"`
	// 标签
	// @gotags: bson:"labels" json:"labels"
	Labels map[string]string `protobuf:"bytes,8,rep,name=labels,proto3" json:"labels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"labels"`
	// 扩展属性
	// @gotags: bson:"extension" json:"extension"
	Extension map[string]string `protobuf:"bytes,11,rep,name=extension,proto3" json:"extension" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"extension"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_endpoint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_endpoint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_endpoint_proto_rawDescGZIP(), []int{1}
}

func (x *Entry) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *Entry) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Entry) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Entry) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *Entry) GetAuthEnable() bool {
	if x != nil {
		return x.AuthEnable
	}
	return false
}

func (x *Entry) GetCodeEnable() bool {
	if x != nil {
		return x.CodeEnable
	}
	return false
}

func (x *Entry) GetPermissionEnable() bool {
	if x != nil {
		return x.PermissionEnable
	}
	return false
}

func (x *Entry) GetAllow() []string {
	if x != nil {
		return x.Allow
	}
	return nil
}

func (x *Entry) GetAuditLog() bool {
	if x != nil {
		return x.AuditLog
	}
	return false
}

func (x *Entry) GetRequiredNamespace() bool {
	if x != nil {
		return x.RequiredNamespace
	}
	return false
}

func (x *Entry) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Entry) GetExtension() map[string]string {
	if x != nil {
		return x.Extension
	}
	return nil
}

type EndpointSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*Endpoint `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *EndpointSet) Reset() {
	*x = EndpointSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_endpoint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointSet) ProtoMessage() {}

func (x *EndpointSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_endpoint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointSet.ProtoReflect.Descriptor instead.
func (*EndpointSet) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_endpoint_proto_rawDescGZIP(), []int{2}
}

func (x *EndpointSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *EndpointSet) GetItems() []*Endpoint {
	if x != nil {
		return x.Items
	}
	return nil
}

// RegistryRequest 服务注册请求
type RegistryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"client_id" validate:"required"
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id" validate:"required"`
	// @gotags: json:"client_secret" validate:"required"
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret" validate:"required"`
	// @gotags: json:"version" validate:"required,lte=32"
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version" validate:"required,lte=32"`
	// @gotags: json:"entries"
	Entries []*Entry `protobuf:"bytes,4,rep,name=entries,proto3" json:"entries"`
}

func (x *RegistryRequest) Reset() {
	*x = RegistryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_endpoint_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryRequest) ProtoMessage() {}

func (x *RegistryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_endpoint_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryRequest.ProtoReflect.Descriptor instead.
func (*RegistryRequest) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_endpoint_proto_rawDescGZIP(), []int{3}
}

func (x *RegistryRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *RegistryRequest) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *RegistryRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *RegistryRequest) GetEntries() []*Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

// RegistryReponse todo
type RegistryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"message"
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
}

func (x *RegistryResponse) Reset() {
	*x = RegistryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_endpoint_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryResponse) ProtoMessage() {}

func (x *RegistryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_endpoint_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryResponse.ProtoReflect.Descriptor instead.
func (*RegistryResponse) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_endpoint_proto_rawDescGZIP(), []int{4}
}

func (x *RegistryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_apps_endpoint_pb_endpoint_proto protoreflect.FileDescriptor

var file_apps_endpoint_pb_endpoint_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f,
	0x70, 0x62, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0xc7,
	0x01, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38,
	0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x22, 0xd7, 0x04, 0x0a, 0x05, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f,
	0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f,
	0x67, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x46, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x4f, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a, 0x0e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x60, 0x0a, 0x0b, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x3b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0xab, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x22, 0x2c, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_endpoint_pb_endpoint_proto_rawDescOnce sync.Once
	file_apps_endpoint_pb_endpoint_proto_rawDescData = file_apps_endpoint_pb_endpoint_proto_rawDesc
)

func file_apps_endpoint_pb_endpoint_proto_rawDescGZIP() []byte {
	file_apps_endpoint_pb_endpoint_proto_rawDescOnce.Do(func() {
		file_apps_endpoint_pb_endpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_endpoint_pb_endpoint_proto_rawDescData)
	})
	return file_apps_endpoint_pb_endpoint_proto_rawDescData
}

var file_apps_endpoint_pb_endpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_apps_endpoint_pb_endpoint_proto_goTypes = []interface{}{
	(*Endpoint)(nil),         // 0: infraboard.mcenter.endpoint.Endpoint
	(*Entry)(nil),            // 1: infraboard.mcenter.endpoint.Entry
	(*EndpointSet)(nil),      // 2: infraboard.mcenter.endpoint.EndpointSet
	(*RegistryRequest)(nil),  // 3: infraboard.mcenter.endpoint.RegistryRequest
	(*RegistryResponse)(nil), // 4: infraboard.mcenter.endpoint.RegistryResponse
	nil,                      // 5: infraboard.mcenter.endpoint.Entry.LabelsEntry
	nil,                      // 6: infraboard.mcenter.endpoint.Entry.ExtensionEntry
}
var file_apps_endpoint_pb_endpoint_proto_depIdxs = []int32{
	1, // 0: infraboard.mcenter.endpoint.Endpoint.entry:type_name -> infraboard.mcenter.endpoint.Entry
	5, // 1: infraboard.mcenter.endpoint.Entry.labels:type_name -> infraboard.mcenter.endpoint.Entry.LabelsEntry
	6, // 2: infraboard.mcenter.endpoint.Entry.extension:type_name -> infraboard.mcenter.endpoint.Entry.ExtensionEntry
	0, // 3: infraboard.mcenter.endpoint.EndpointSet.items:type_name -> infraboard.mcenter.endpoint.Endpoint
	1, // 4: infraboard.mcenter.endpoint.RegistryRequest.entries:type_name -> infraboard.mcenter.endpoint.Entry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_apps_endpoint_pb_endpoint_proto_init() }
func file_apps_endpoint_pb_endpoint_proto_init() {
	if File_apps_endpoint_pb_endpoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_endpoint_pb_endpoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endpoint); i {
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
		file_apps_endpoint_pb_endpoint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_apps_endpoint_pb_endpoint_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointSet); i {
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
		file_apps_endpoint_pb_endpoint_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistryRequest); i {
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
		file_apps_endpoint_pb_endpoint_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistryResponse); i {
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
			RawDescriptor: file_apps_endpoint_pb_endpoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_endpoint_pb_endpoint_proto_goTypes,
		DependencyIndexes: file_apps_endpoint_pb_endpoint_proto_depIdxs,
		MessageInfos:      file_apps_endpoint_pb_endpoint_proto_msgTypes,
	}.Build()
	File_apps_endpoint_pb_endpoint_proto = out.File
	file_apps_endpoint_pb_endpoint_proto_rawDesc = nil
	file_apps_endpoint_pb_endpoint_proto_goTypes = nil
	file_apps_endpoint_pb_endpoint_proto_depIdxs = nil
}
