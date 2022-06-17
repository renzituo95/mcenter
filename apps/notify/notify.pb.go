// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: apps/notify/pb/notify.proto

package notify

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

type SendSMSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 短信模版的Id
	// @gotags: bson:"template_id" json:"template_id"
	TemplateId string `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id" bson:"template_id"`
	// 模版参数
	// @gotags: bson:"template_params" json:"template_params"
	TemplateParams []string `protobuf:"bytes,2,rep,name=template_params,json=templateParams,proto3" json:"template_params" bson:"template_params"`
	// 手机号码
	// @gotags: bson:"phone_numbers" json:"phone_numbers"
	PhoneNumbers []string `protobuf:"bytes,3,rep,name=phone_numbers,json=phoneNumbers,proto3" json:"phone_numbers" bson:"phone_numbers"`
}

func (x *SendSMSRequest) Reset() {
	*x = SendSMSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_notify_pb_notify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSMSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSMSRequest) ProtoMessage() {}

func (x *SendSMSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_notify_pb_notify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSMSRequest.ProtoReflect.Descriptor instead.
func (*SendSMSRequest) Descriptor() ([]byte, []int) {
	return file_apps_notify_pb_notify_proto_rawDescGZIP(), []int{0}
}

func (x *SendSMSRequest) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *SendSMSRequest) GetTemplateParams() []string {
	if x != nil {
		return x.TemplateParams
	}
	return nil
}

func (x *SendSMSRequest) GetPhoneNumbers() []string {
	if x != nil {
		return x.PhoneNumbers
	}
	return nil
}

type SendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendResponse) Reset() {
	*x = SendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_notify_pb_notify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResponse) ProtoMessage() {}

func (x *SendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_notify_pb_notify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResponse.ProtoReflect.Descriptor instead.
func (*SendResponse) Descriptor() ([]byte, []int) {
	return file_apps_notify_pb_notify_proto_rawDescGZIP(), []int{1}
}

var File_apps_notify_pb_notify_proto protoreflect.FileDescriptor

var file_apps_notify_pb_notify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x70, 0x62,
	0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x22, 0x7f, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64,
	0x53, 0x4d, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x64, 0x0a, 0x03, 0x52, 0x50, 0x43,
	0x12, 0x5d, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d, 0x53, 0x12, 0x29, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d, 0x53, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_notify_pb_notify_proto_rawDescOnce sync.Once
	file_apps_notify_pb_notify_proto_rawDescData = file_apps_notify_pb_notify_proto_rawDesc
)

func file_apps_notify_pb_notify_proto_rawDescGZIP() []byte {
	file_apps_notify_pb_notify_proto_rawDescOnce.Do(func() {
		file_apps_notify_pb_notify_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_notify_pb_notify_proto_rawDescData)
	})
	return file_apps_notify_pb_notify_proto_rawDescData
}

var file_apps_notify_pb_notify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apps_notify_pb_notify_proto_goTypes = []interface{}{
	(*SendSMSRequest)(nil), // 0: infraboard.mcenter.notify.SendSMSRequest
	(*SendResponse)(nil),   // 1: infraboard.mcenter.notify.SendResponse
}
var file_apps_notify_pb_notify_proto_depIdxs = []int32{
	0, // 0: infraboard.mcenter.notify.RPC.SendSMS:input_type -> infraboard.mcenter.notify.SendSMSRequest
	1, // 1: infraboard.mcenter.notify.RPC.SendSMS:output_type -> infraboard.mcenter.notify.SendResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apps_notify_pb_notify_proto_init() }
func file_apps_notify_pb_notify_proto_init() {
	if File_apps_notify_pb_notify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_notify_pb_notify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSMSRequest); i {
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
		file_apps_notify_pb_notify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendResponse); i {
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
			RawDescriptor: file_apps_notify_pb_notify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_notify_pb_notify_proto_goTypes,
		DependencyIndexes: file_apps_notify_pb_notify_proto_depIdxs,
		MessageInfos:      file_apps_notify_pb_notify_proto_msgTypes,
	}.Build()
	File_apps_notify_pb_notify_proto = out.File
	file_apps_notify_pb_notify_proto_rawDesc = nil
	file_apps_notify_pb_notify_proto_goTypes = nil
	file_apps_notify_pb_notify_proto_depIdxs = nil
}
