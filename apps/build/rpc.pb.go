// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/build/pb/rpc.proto

package build

import (
	label "github.com/infraboard/mcenter/apps/label"
	request "github.com/infraboard/mcube/http/request"
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

type RunBuildRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 所属列表
	// @gotags: bson:"service_selector" json:"service_selector"
	ServiceSelector *label.LabelSelector `protobuf:"bytes,1,opt,name=service_selector,json=serviceSelector,proto3" json:"service_selector" bson:"service_selector"`
	// 所属列表
	// @gotags: bson:"service_ids" json:"service_ids"
	ServiceIds []string `protobuf:"bytes,2,rep,name=service_ids,json=serviceIds,proto3" json:"service_ids" bson:"service_ids"`
	// 自动部署相关配置
	// @gotags: bson:"auto_deploy_options" json:"auto_deploy_options"
	AutoDeployOptions *AutoDeployOptions `protobuf:"bytes,3,opt,name=auto_deploy_options,json=autoDeployOptions,proto3" json:"auto_deploy_options" bson:"auto_deploy_options"`
}

func (x *RunBuildRequest) Reset() {
	*x = RunBuildRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_build_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunBuildRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunBuildRequest) ProtoMessage() {}

func (x *RunBuildRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_build_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunBuildRequest.ProtoReflect.Descriptor instead.
func (*RunBuildRequest) Descriptor() ([]byte, []int) {
	return file_apps_build_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *RunBuildRequest) GetServiceSelector() *label.LabelSelector {
	if x != nil {
		return x.ServiceSelector
	}
	return nil
}

func (x *RunBuildRequest) GetServiceIds() []string {
	if x != nil {
		return x.ServiceIds
	}
	return nil
}

func (x *RunBuildRequest) GetAutoDeployOptions() *AutoDeployOptions {
	if x != nil {
		return x.AutoDeployOptions
	}
	return nil
}

// 自动部署相关配置
type AutoDeployOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 是否开启服务的自动部署,如果开启 构建成功后会执行自动部署
	// @gotags: bson:"enabled" json:"enabled"
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled" bson:"enabled"`
	// 自动部署的环境列表, 那些环境允许执行自动部署
	// @gotags: bson:"envs" json:"envs"
	Envs []string `protobuf:"bytes,2,rep,name=envs,proto3" json:"envs" bson:"envs"`
	// 服务标签, 打有相应标签的部署才执行自动部署
	// @gotags: bson:"deploy_selector" json:"deploy_selector"
	DeploySelector *label.LabelSelector `protobuf:"bytes,14,opt,name=deploy_selector,json=deploySelector,proto3" json:"deploy_selector" bson:"deploy_selector"`
}

func (x *AutoDeployOptions) Reset() {
	*x = AutoDeployOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_build_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoDeployOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoDeployOptions) ProtoMessage() {}

func (x *AutoDeployOptions) ProtoReflect() protoreflect.Message {
	mi := &file_apps_build_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoDeployOptions.ProtoReflect.Descriptor instead.
func (*AutoDeployOptions) Descriptor() ([]byte, []int) {
	return file_apps_build_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *AutoDeployOptions) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *AutoDeployOptions) GetEnvs() []string {
	if x != nil {
		return x.Envs
	}
	return nil
}

func (x *AutoDeployOptions) GetDeploySelector() *label.LabelSelector {
	if x != nil {
		return x.DeploySelector
	}
	return nil
}

type QueryBuildRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 所属Domain
	// @gotags: json:"domain" validate:"required"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain" validate:"required"`
	// 所属空间
	// @gotags: json:"namespace"
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace"`
	// 所属服务
	// @gotags: json:"service_id"
	ServiceId string `protobuf:"bytes,4,opt,name=service_id,json=serviceId,proto3" json:"service_id"`
	// 构建Id
	// @gotags: json:"task_ids"
	TaskIds []string `protobuf:"bytes,5,rep,name=task_ids,json=taskIds,proto3" json:"task_ids"`
}

func (x *QueryBuildRequest) Reset() {
	*x = QueryBuildRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_build_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryBuildRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryBuildRequest) ProtoMessage() {}

func (x *QueryBuildRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_build_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryBuildRequest.ProtoReflect.Descriptor instead.
func (*QueryBuildRequest) Descriptor() ([]byte, []int) {
	return file_apps_build_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *QueryBuildRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryBuildRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *QueryBuildRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *QueryBuildRequest) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *QueryBuildRequest) GetTaskIds() []string {
	if x != nil {
		return x.TaskIds
	}
	return nil
}

var File_apps_build_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_build_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70, 0x62, 0x2f,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x1a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f,
	0x70, 0x62, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x01, 0x0a, 0x0f, 0x52, 0x75,
	0x6e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x52, 0x0a,
	0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x73, 0x12, 0x5b, 0x0a, 0x13, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x11, 0x61, 0x75,
	0x74, 0x6f, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x93, 0x01, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x6f, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x65, 0x6e, 0x76, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x65,
	0x6e, 0x76, 0x73, 0x12, 0x50, 0x0a, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0xbb, 0x01, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x73, 0x6b,
	0x49, 0x64, 0x73, 0x32, 0xc4, 0x01, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x5a, 0x0a, 0x08, 0x52,
	0x75, 0x6e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x29, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2e, 0x52, 0x75, 0x6e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x61, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x2b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x74, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73,
	0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_build_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_build_pb_rpc_proto_rawDescData = file_apps_build_pb_rpc_proto_rawDesc
)

func file_apps_build_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_build_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_build_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_build_pb_rpc_proto_rawDescData)
	})
	return file_apps_build_pb_rpc_proto_rawDescData
}

var file_apps_build_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apps_build_pb_rpc_proto_goTypes = []interface{}{
	(*RunBuildRequest)(nil),     // 0: infraboard.mcenter.build.RunBuildRequest
	(*AutoDeployOptions)(nil),   // 1: infraboard.mcenter.build.AutoDeployOptions
	(*QueryBuildRequest)(nil),   // 2: infraboard.mcenter.build.QueryBuildRequest
	(*label.LabelSelector)(nil), // 3: infraboard.mcenter.label.LabelSelector
	(*request.PageRequest)(nil), // 4: infraboard.mcube.page.PageRequest
	(*BuildTask)(nil),           // 5: infraboard.mcenter.build.BuildTask
	(*BuildTaskSet)(nil),        // 6: infraboard.mcenter.build.BuildTaskSet
}
var file_apps_build_pb_rpc_proto_depIdxs = []int32{
	3, // 0: infraboard.mcenter.build.RunBuildRequest.service_selector:type_name -> infraboard.mcenter.label.LabelSelector
	1, // 1: infraboard.mcenter.build.RunBuildRequest.auto_deploy_options:type_name -> infraboard.mcenter.build.AutoDeployOptions
	3, // 2: infraboard.mcenter.build.AutoDeployOptions.deploy_selector:type_name -> infraboard.mcenter.label.LabelSelector
	4, // 3: infraboard.mcenter.build.QueryBuildRequest.page:type_name -> infraboard.mcube.page.PageRequest
	0, // 4: infraboard.mcenter.build.RPC.RunBuild:input_type -> infraboard.mcenter.build.RunBuildRequest
	2, // 5: infraboard.mcenter.build.RPC.QueryBuild:input_type -> infraboard.mcenter.build.QueryBuildRequest
	5, // 6: infraboard.mcenter.build.RPC.RunBuild:output_type -> infraboard.mcenter.build.BuildTask
	6, // 7: infraboard.mcenter.build.RPC.QueryBuild:output_type -> infraboard.mcenter.build.BuildTaskSet
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_apps_build_pb_rpc_proto_init() }
func file_apps_build_pb_rpc_proto_init() {
	if File_apps_build_pb_rpc_proto != nil {
		return
	}
	file_apps_build_pb_build_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_build_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunBuildRequest); i {
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
		file_apps_build_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoDeployOptions); i {
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
		file_apps_build_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryBuildRequest); i {
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
			RawDescriptor: file_apps_build_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_build_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_build_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_build_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_build_pb_rpc_proto = out.File
	file_apps_build_pb_rpc_proto_rawDesc = nil
	file_apps_build_pb_rpc_proto_goTypes = nil
	file_apps_build_pb_rpc_proto_depIdxs = nil
}
