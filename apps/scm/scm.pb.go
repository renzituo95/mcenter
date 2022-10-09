// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/scm/pb/scm.proto

package scm

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
	// gitlab 代码仓库
	TYPE_GITLAB TYPE = 0
	// github 代码仓库
	TYPE_GITHUB TYPE = 1
)

// Enum value maps for TYPE.
var (
	TYPE_name = map[int32]string{
		0: "GITLAB",
		1: "GITHUB",
	}
	TYPE_value = map[string]int32{
		"GITLAB": 0,
		"GITHUB": 1,
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
	return file_apps_scm_pb_scm_proto_enumTypes[0].Descriptor()
}

func (TYPE) Type() protoreflect.EnumType {
	return &file_apps_scm_pb_scm_proto_enumTypes[0]
}

func (x TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TYPE.Descriptor instead.
func (TYPE) EnumDescriptor() ([]byte, []int) {
	return file_apps_scm_pb_scm_proto_rawDescGZIP(), []int{0}
}

type WebHookEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 事件类型
	// @gotags: bson:"object_kind" json:"object_kind" validate:"required"
	ObjectKind string `protobuf:"bytes,1,opt,name=object_kind,json=objectKind,proto3" json:"object_kind" bson:"object_kind" validate:"required"`
	// 事件名称
	// @gotags: bson:"event_name" json:"event_name" validate:"required"
	EventName string `protobuf:"bytes,2,opt,name=event_name,json=eventName,proto3" json:"event_name" bson:"event_name" validate:"required"`
	// 关联分支
	// @gotags: bson:"ref" json:"ref" validate:"required"
	Ref string `protobuf:"bytes,3,opt,name=ref,proto3" json:"ref" bson:"ref" validate:"required"`
	// 触发者用户ID
	// @gotags: bson:"user_id" json:"user_id"
	UserId int64 `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id" bson:"user_id"`
	// 触发者用户名称
	// @gotags: bson:"user_name" json:"user_name"
	UserName string `protobuf:"bytes,5,opt,name=user_name,json=userName,proto3" json:"user_name" bson:"user_name"`
	// 用户头像
	// @gotags: bson:"user_avatar" json:"user_avatar"
	UserAvatar string `protobuf:"bytes,6,opt,name=user_avatar,json=userAvatar,proto3" json:"user_avatar" bson:"user_avatar"`
	// 用户头像
	// @gotags: bson:"project" json:"project"
	Project *Project `protobuf:"bytes,7,opt,name=project,proto3" json:"project" bson:"project"`
	// Commit信息
	// @gotags: bson:"commits" json:"commits"
	Commits []*Commit `protobuf:"bytes,8,rep,name=commits,proto3" json:"commits" bson:"commits"`
}

func (x *WebHookEvent) Reset() {
	*x = WebHookEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_scm_pb_scm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebHookEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebHookEvent) ProtoMessage() {}

func (x *WebHookEvent) ProtoReflect() protoreflect.Message {
	mi := &file_apps_scm_pb_scm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebHookEvent.ProtoReflect.Descriptor instead.
func (*WebHookEvent) Descriptor() ([]byte, []int) {
	return file_apps_scm_pb_scm_proto_rawDescGZIP(), []int{0}
}

func (x *WebHookEvent) GetObjectKind() string {
	if x != nil {
		return x.ObjectKind
	}
	return ""
}

func (x *WebHookEvent) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

func (x *WebHookEvent) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *WebHookEvent) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *WebHookEvent) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *WebHookEvent) GetUserAvatar() string {
	if x != nil {
		return x.UserAvatar
	}
	return ""
}

func (x *WebHookEvent) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *WebHookEvent) GetCommits() []*Commit {
	if x != nil {
		return x.Commits
	}
	return nil
}

type Commit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id
	// @gotags: bson:"id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
	// commit message
	// @gotags: bson:"message" json:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message" bson:"message"`
	// title
	// @gotags: bson:"title" json:"title"
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title" bson:"title"`
	// 文本格式时间
	// @gotags: bson:"timestamp" json:"timestamp"
	Timestamp string `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp" bson:"timestamp"`
	// commit对应的url
	// @gotags: bson:"url" json:"url"
	Url string `protobuf:"bytes,5,opt,name=url,proto3" json:"url" bson:"url"`
	// 作者
	// @gotags: bson:"author" json:"author"
	Author *Author `protobuf:"bytes,6,opt,name=author,proto3" json:"author" bson:"author"`
	// 新加的文件
	// @gotags: bson:"added" json:"added"
	Added []string `protobuf:"bytes,7,rep,name=added,proto3" json:"added" bson:"added"`
	// 修改的文件
	// @gotags: bson:"modified" json:"modified"
	Modified []string `protobuf:"bytes,8,rep,name=modified,proto3" json:"modified" bson:"modified"`
	// 删除的文件
	// @gotags: bson:"removed" json:"removed"
	Removed []string `protobuf:"bytes,9,rep,name=removed,proto3" json:"removed" bson:"removed"`
}

func (x *Commit) Reset() {
	*x = Commit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_scm_pb_scm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commit) ProtoMessage() {}

func (x *Commit) ProtoReflect() protoreflect.Message {
	mi := &file_apps_scm_pb_scm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commit.ProtoReflect.Descriptor instead.
func (*Commit) Descriptor() ([]byte, []int) {
	return file_apps_scm_pb_scm_proto_rawDescGZIP(), []int{1}
}

func (x *Commit) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Commit) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Commit) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Commit) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *Commit) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Commit) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Commit) GetAdded() []string {
	if x != nil {
		return x.Added
	}
	return nil
}

func (x *Commit) GetModified() []string {
	if x != nil {
		return x.Modified
	}
	return nil
}

func (x *Commit) GetRemoved() []string {
	if x != nil {
		return x.Removed
	}
	return nil
}

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 作者名称
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"name"`
	// 作者邮箱
	// @gotags: bson:"email" json:"email"
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email" bson:"email"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_scm_pb_scm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_apps_scm_pb_scm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_apps_scm_pb_scm_proto_rawDescGZIP(), []int{2}
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Author) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 项目id
	// @gotags: bson:"id" json:"id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" bson:"id"`
	// 描述
	// @gotags: bson:"description" json:"description"
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description" bson:"description"`
	// 名称
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" bson:"name"`
	// ssh 地址
	// @gotags: bson:"ssh_url_to_repo" json:"ssh_url_to_repo"
	GitSshUrl string `protobuf:"bytes,4,opt,name=git_ssh_url,json=gitSshUrl,proto3" json:"ssh_url_to_repo" bson:"ssh_url_to_repo"`
	// http 地址
	// @gotags: bson:"http_url_to_repo" json:"http_url_to_repo"
	GitHttpUrl string `protobuf:"bytes,5,opt,name=git_http_url,json=gitHttpUrl,proto3" json:"http_url_to_repo" bson:"http_url_to_repo"`
	// namespace
	// @gotags: bson:"path_with_namespace" json:"path_with_namespace"
	NamespacePath string `protobuf:"bytes,6,opt,name=namespace_path,json=namespacePath,proto3" json:"path_with_namespace" bson:"path_with_namespace"`
	// 是否已经同步
	// @gotags: bson:"has_synced" json:"has_synced"
	HasSynced bool `protobuf:"varint,7,opt,name=has_synced,json=hasSynced,proto3" json:"has_synced" bson:"has_synced"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_scm_pb_scm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_apps_scm_pb_scm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_apps_scm_pb_scm_proto_rawDescGZIP(), []int{3}
}

func (x *Project) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Project) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetGitSshUrl() string {
	if x != nil {
		return x.GitSshUrl
	}
	return ""
}

func (x *Project) GetGitHttpUrl() string {
	if x != nil {
		return x.GitHttpUrl
	}
	return ""
}

func (x *Project) GetNamespacePath() string {
	if x != nil {
		return x.NamespacePath
	}
	return ""
}

func (x *Project) GetHasSynced() bool {
	if x != nil {
		return x.HasSynced
	}
	return false
}

type ProjectSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"items"
	Items []*Project `protobuf:"bytes,9,rep,name=items,proto3" json:"items"`
}

func (x *ProjectSet) Reset() {
	*x = ProjectSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_scm_pb_scm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectSet) ProtoMessage() {}

func (x *ProjectSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_scm_pb_scm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectSet.ProtoReflect.Descriptor instead.
func (*ProjectSet) Descriptor() ([]byte, []int) {
	return file_apps_scm_pb_scm_proto_rawDescGZIP(), []int{4}
}

func (x *ProjectSet) GetItems() []*Project {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_apps_scm_pb_scm_proto protoreflect.FileDescriptor

var file_apps_scm_pb_scm_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x63, 0x6d, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x63,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x63, 0x6d, 0x22,
	0xac, 0x02, 0x0a, 0x0c, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4b, 0x69, 0x6e,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72,
	0x65, 0x66, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x39, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x73, 0x63, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x38, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x63, 0x6d, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x22, 0xfc,
	0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x36, 0x0a, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73,
	0x63, 0x6d, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x22, 0x32, 0x0a,
	0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x22, 0xd7, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x67, 0x69, 0x74, 0x5f, 0x73, 0x73, 0x68, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x69, 0x74, 0x53, 0x73, 0x68,
	0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0c, 0x67, 0x69, 0x74, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x69, 0x74, 0x48, 0x74,
	0x74, 0x70, 0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a,
	0x68, 0x61, 0x73, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x68, 0x61, 0x73, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x22, 0x43, 0x0a, 0x0a, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x63,
	0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x2a, 0x1e, 0x0a, 0x04, 0x54, 0x59, 0x50, 0x45, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x49, 0x54, 0x4c,
	0x41, 0x42, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x49, 0x54, 0x48, 0x55, 0x42, 0x10, 0x01,
	0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x63, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_apps_scm_pb_scm_proto_rawDescOnce sync.Once
	file_apps_scm_pb_scm_proto_rawDescData = file_apps_scm_pb_scm_proto_rawDesc
)

func file_apps_scm_pb_scm_proto_rawDescGZIP() []byte {
	file_apps_scm_pb_scm_proto_rawDescOnce.Do(func() {
		file_apps_scm_pb_scm_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_scm_pb_scm_proto_rawDescData)
	})
	return file_apps_scm_pb_scm_proto_rawDescData
}

var file_apps_scm_pb_scm_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_scm_pb_scm_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_apps_scm_pb_scm_proto_goTypes = []interface{}{
	(TYPE)(0),            // 0: infraboard.mcenter.scm.TYPE
	(*WebHookEvent)(nil), // 1: infraboard.mcenter.scm.WebHookEvent
	(*Commit)(nil),       // 2: infraboard.mcenter.scm.Commit
	(*Author)(nil),       // 3: infraboard.mcenter.scm.Author
	(*Project)(nil),      // 4: infraboard.mcenter.scm.Project
	(*ProjectSet)(nil),   // 5: infraboard.mcenter.scm.ProjectSet
}
var file_apps_scm_pb_scm_proto_depIdxs = []int32{
	4, // 0: infraboard.mcenter.scm.WebHookEvent.project:type_name -> infraboard.mcenter.scm.Project
	2, // 1: infraboard.mcenter.scm.WebHookEvent.commits:type_name -> infraboard.mcenter.scm.Commit
	3, // 2: infraboard.mcenter.scm.Commit.author:type_name -> infraboard.mcenter.scm.Author
	4, // 3: infraboard.mcenter.scm.ProjectSet.items:type_name -> infraboard.mcenter.scm.Project
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_apps_scm_pb_scm_proto_init() }
func file_apps_scm_pb_scm_proto_init() {
	if File_apps_scm_pb_scm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_scm_pb_scm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebHookEvent); i {
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
		file_apps_scm_pb_scm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commit); i {
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
		file_apps_scm_pb_scm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_apps_scm_pb_scm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_apps_scm_pb_scm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectSet); i {
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
			RawDescriptor: file_apps_scm_pb_scm_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_scm_pb_scm_proto_goTypes,
		DependencyIndexes: file_apps_scm_pb_scm_proto_depIdxs,
		EnumInfos:         file_apps_scm_pb_scm_proto_enumTypes,
		MessageInfos:      file_apps_scm_pb_scm_proto_msgTypes,
	}.Build()
	File_apps_scm_pb_scm_proto = out.File
	file_apps_scm_pb_scm_proto_rawDesc = nil
	file_apps_scm_pb_scm_proto_goTypes = nil
	file_apps_scm_pb_scm_proto_depIdxs = nil
}
