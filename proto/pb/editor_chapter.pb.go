// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: pb/editor_chapter.proto

package pb

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

type SocketPingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Req int32 `protobuf:"varint,1,opt,name=Req,proto3" json:"Req,omitempty"` //默认传1
}

func (x *SocketPingReq) Reset() {
	*x = SocketPingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_editor_chapter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocketPingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocketPingReq) ProtoMessage() {}

func (x *SocketPingReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_editor_chapter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocketPingReq.ProtoReflect.Descriptor instead.
func (*SocketPingReq) Descriptor() ([]byte, []int) {
	return file_pb_editor_chapter_proto_rawDescGZIP(), []int{0}
}

func (x *SocketPingReq) GetReq() int32 {
	if x != nil {
		return x.Req
	}
	return 0
}

type SocketPingRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Rsp  int32 `protobuf:"varint,2,opt,name=Rsp,proto3" json:"Rsp,omitempty"` //默认回2
}

func (x *SocketPingRsp) Reset() {
	*x = SocketPingRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_editor_chapter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocketPingRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocketPingRsp) ProtoMessage() {}

func (x *SocketPingRsp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_editor_chapter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocketPingRsp.ProtoReflect.Descriptor instead.
func (*SocketPingRsp) Descriptor() ([]byte, []int) {
	return file_pb_editor_chapter_proto_rawDescGZIP(), []int{1}
}

func (x *SocketPingRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SocketPingRsp) GetRsp() int32 {
	if x != nil {
		return x.Rsp
	}
	return 0
}

type Chapter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color int32 `protobuf:"varint,1,opt,name=Color,proto3" json:"Color,omitempty"` //颜色
	Row   int32 `protobuf:"varint,2,opt,name=Row,proto3" json:"Row,omitempty"`
	Col   int32 `protobuf:"varint,3,opt,name=Col,proto3" json:"Col,omitempty"`
	Type  int32 `protobuf:"varint,4,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *Chapter) Reset() {
	*x = Chapter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_editor_chapter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chapter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chapter) ProtoMessage() {}

func (x *Chapter) ProtoReflect() protoreflect.Message {
	mi := &file_pb_editor_chapter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chapter.ProtoReflect.Descriptor instead.
func (*Chapter) Descriptor() ([]byte, []int) {
	return file_pb_editor_chapter_proto_rawDescGZIP(), []int{2}
}

func (x *Chapter) GetColor() int32 {
	if x != nil {
		return x.Color
	}
	return 0
}

func (x *Chapter) GetRow() int32 {
	if x != nil {
		return x.Row
	}
	return 0
}

func (x *Chapter) GetCol() int32 {
	if x != nil {
		return x.Col
	}
	return 0
}

func (x *Chapter) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type ChapterDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info []*Chapter `protobuf:"bytes,1,rep,name=Info,proto3" json:"Info,omitempty"` //关卡信息
}

func (x *ChapterDetail) Reset() {
	*x = ChapterDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_editor_chapter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChapterDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChapterDetail) ProtoMessage() {}

func (x *ChapterDetail) ProtoReflect() protoreflect.Message {
	mi := &file_pb_editor_chapter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChapterDetail.ProtoReflect.Descriptor instead.
func (*ChapterDetail) Descriptor() ([]byte, []int) {
	return file_pb_editor_chapter_proto_rawDescGZIP(), []int{3}
}

func (x *ChapterDetail) GetInfo() []*Chapter {
	if x != nil {
		return x.Info
	}
	return nil
}

// 编辑关卡
type EditorChapterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChapterId int32      `protobuf:"varint,1,opt,name=ChapterId,proto3" json:"ChapterId,omitempty"` //如果是修改关卡信息，传关卡id，否则不传
	Info      []*Chapter `protobuf:"bytes,2,rep,name=Info,proto3" json:"Info,omitempty"`            //关卡信息
}

func (x *EditorChapterReq) Reset() {
	*x = EditorChapterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_editor_chapter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditorChapterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditorChapterReq) ProtoMessage() {}

func (x *EditorChapterReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_editor_chapter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditorChapterReq.ProtoReflect.Descriptor instead.
func (*EditorChapterReq) Descriptor() ([]byte, []int) {
	return file_pb_editor_chapter_proto_rawDescGZIP(), []int{4}
}

func (x *EditorChapterReq) GetChapterId() int32 {
	if x != nil {
		return x.ChapterId
	}
	return 0
}

func (x *EditorChapterReq) GetInfo() []*Chapter {
	if x != nil {
		return x.Info
	}
	return nil
}

type EditorChapterRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Info *Chapter `protobuf:"bytes,2,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *EditorChapterRsp) Reset() {
	*x = EditorChapterRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_editor_chapter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditorChapterRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditorChapterRsp) ProtoMessage() {}

func (x *EditorChapterRsp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_editor_chapter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditorChapterRsp.ProtoReflect.Descriptor instead.
func (*EditorChapterRsp) Descriptor() ([]byte, []int) {
	return file_pb_editor_chapter_proto_rawDescGZIP(), []int{5}
}

func (x *EditorChapterRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *EditorChapterRsp) GetInfo() *Chapter {
	if x != nil {
		return x.Info
	}
	return nil
}

// 获取当前所有关卡
type GetAllChapterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllChapterReq) Reset() {
	*x = GetAllChapterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_editor_chapter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllChapterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllChapterReq) ProtoMessage() {}

func (x *GetAllChapterReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_editor_chapter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllChapterReq.ProtoReflect.Descriptor instead.
func (*GetAllChapterReq) Descriptor() ([]byte, []int) {
	return file_pb_editor_chapter_proto_rawDescGZIP(), []int{6}
}

type GetAllChapterRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       int32   `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	ChapterIds []int32 `protobuf:"varint,2,rep,packed,name=ChapterIds,proto3" json:"ChapterIds,omitempty"` //关卡id
}

func (x *GetAllChapterRsp) Reset() {
	*x = GetAllChapterRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_editor_chapter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllChapterRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllChapterRsp) ProtoMessage() {}

func (x *GetAllChapterRsp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_editor_chapter_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllChapterRsp.ProtoReflect.Descriptor instead.
func (*GetAllChapterRsp) Descriptor() ([]byte, []int) {
	return file_pb_editor_chapter_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllChapterRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetAllChapterRsp) GetChapterIds() []int32 {
	if x != nil {
		return x.ChapterIds
	}
	return nil
}

// 获取关卡详情
type ChapterDetailByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChapterId int32 `protobuf:"varint,1,opt,name=ChapterId,proto3" json:"ChapterId,omitempty"`
	Tp        int32 `protobuf:"varint,2,opt,name=Tp,proto3" json:"Tp,omitempty"` //1 GM编辑页面 2 正常游戏
}

func (x *ChapterDetailByIdReq) Reset() {
	*x = ChapterDetailByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_editor_chapter_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChapterDetailByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChapterDetailByIdReq) ProtoMessage() {}

func (x *ChapterDetailByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_editor_chapter_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChapterDetailByIdReq.ProtoReflect.Descriptor instead.
func (*ChapterDetailByIdReq) Descriptor() ([]byte, []int) {
	return file_pb_editor_chapter_proto_rawDescGZIP(), []int{8}
}

func (x *ChapterDetailByIdReq) GetChapterId() int32 {
	if x != nil {
		return x.ChapterId
	}
	return 0
}

func (x *ChapterDetailByIdReq) GetTp() int32 {
	if x != nil {
		return x.Tp
	}
	return 0
}

type ChapterDetailByIdRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int32          `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Info      *ChapterDetail `protobuf:"bytes,2,opt,name=Info,proto3" json:"Info,omitempty"`
	ChapterId int32          `protobuf:"varint,3,opt,name=ChapterId,proto3" json:"ChapterId,omitempty"` //关卡id
	Tp        int32          `protobuf:"varint,4,opt,name=Tp,proto3" json:"Tp,omitempty"`               //1 GM编辑页面 2 正常游戏
}

func (x *ChapterDetailByIdRsp) Reset() {
	*x = ChapterDetailByIdRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_editor_chapter_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChapterDetailByIdRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChapterDetailByIdRsp) ProtoMessage() {}

func (x *ChapterDetailByIdRsp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_editor_chapter_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChapterDetailByIdRsp.ProtoReflect.Descriptor instead.
func (*ChapterDetailByIdRsp) Descriptor() ([]byte, []int) {
	return file_pb_editor_chapter_proto_rawDescGZIP(), []int{9}
}

func (x *ChapterDetailByIdRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ChapterDetailByIdRsp) GetInfo() *ChapterDetail {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ChapterDetailByIdRsp) GetChapterId() int32 {
	if x != nil {
		return x.ChapterId
	}
	return 0
}

func (x *ChapterDetailByIdRsp) GetTp() int32 {
	if x != nil {
		return x.Tp
	}
	return 0
}

var File_pb_editor_chapter_proto protoreflect.FileDescriptor

var file_pb_editor_chapter_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x62, 0x2f, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x68, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x21, 0x0a,
	0x0d, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x10,
	0x0a, 0x03, 0x52, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x52, 0x65, 0x71,
	0x22, 0x35, 0x0a, 0x0d, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x73,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x73, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x52, 0x73, 0x70, 0x22, 0x57, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x6f, 0x77, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x52, 0x6f, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x6f,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x43, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x30, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x1f, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x51, 0x0a, 0x10, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x43, 0x68, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x47, 0x0a, 0x10, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x43,
	0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x12,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x22, 0x46, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x68,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a,
	0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x44, 0x0a, 0x14, 0x43, 0x68,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x54, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x54, 0x70,
	0x22, 0x7f, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x42, 0x79, 0x49, 0x64, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x54,
	0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pb_editor_chapter_proto_rawDescOnce sync.Once
	file_pb_editor_chapter_proto_rawDescData = file_pb_editor_chapter_proto_rawDesc
)

func file_pb_editor_chapter_proto_rawDescGZIP() []byte {
	file_pb_editor_chapter_proto_rawDescOnce.Do(func() {
		file_pb_editor_chapter_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_editor_chapter_proto_rawDescData)
	})
	return file_pb_editor_chapter_proto_rawDescData
}

var file_pb_editor_chapter_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pb_editor_chapter_proto_goTypes = []interface{}{
	(*SocketPingReq)(nil),        // 0: pb.SocketPingReq
	(*SocketPingRsp)(nil),        // 1: pb.SocketPingRsp
	(*Chapter)(nil),              // 2: pb.Chapter
	(*ChapterDetail)(nil),        // 3: pb.ChapterDetail
	(*EditorChapterReq)(nil),     // 4: pb.EditorChapterReq
	(*EditorChapterRsp)(nil),     // 5: pb.EditorChapterRsp
	(*GetAllChapterReq)(nil),     // 6: pb.GetAllChapterReq
	(*GetAllChapterRsp)(nil),     // 7: pb.GetAllChapterRsp
	(*ChapterDetailByIdReq)(nil), // 8: pb.ChapterDetailByIdReq
	(*ChapterDetailByIdRsp)(nil), // 9: pb.ChapterDetailByIdRsp
}
var file_pb_editor_chapter_proto_depIdxs = []int32{
	2, // 0: pb.ChapterDetail.Info:type_name -> pb.Chapter
	2, // 1: pb.EditorChapterReq.Info:type_name -> pb.Chapter
	2, // 2: pb.EditorChapterRsp.Info:type_name -> pb.Chapter
	3, // 3: pb.ChapterDetailByIdRsp.Info:type_name -> pb.ChapterDetail
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pb_editor_chapter_proto_init() }
func file_pb_editor_chapter_proto_init() {
	if File_pb_editor_chapter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_editor_chapter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocketPingReq); i {
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
		file_pb_editor_chapter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocketPingRsp); i {
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
		file_pb_editor_chapter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chapter); i {
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
		file_pb_editor_chapter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChapterDetail); i {
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
		file_pb_editor_chapter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditorChapterReq); i {
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
		file_pb_editor_chapter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditorChapterRsp); i {
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
		file_pb_editor_chapter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllChapterReq); i {
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
		file_pb_editor_chapter_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllChapterRsp); i {
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
		file_pb_editor_chapter_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChapterDetailByIdReq); i {
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
		file_pb_editor_chapter_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChapterDetailByIdRsp); i {
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
			RawDescriptor: file_pb_editor_chapter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_editor_chapter_proto_goTypes,
		DependencyIndexes: file_pb_editor_chapter_proto_depIdxs,
		MessageInfos:      file_pb_editor_chapter_proto_msgTypes,
	}.Build()
	File_pb_editor_chapter_proto = out.File
	file_pb_editor_chapter_proto_rawDesc = nil
	file_pb_editor_chapter_proto_goTypes = nil
	file_pb_editor_chapter_proto_depIdxs = nil
}
