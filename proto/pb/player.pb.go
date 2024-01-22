// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: pb/player.proto

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

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid              string          `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	BackPack         map[int32]int32 `protobuf:"bytes,2,rep,name=BackPack,proto3" json:"BackPack,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`                 //背包
	CheckPointRecord map[int32]int32 `protobuf:"bytes,3,rep,name=CheckPointRecord,proto3" json:"CheckPointRecord,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` //通关记录
	LoginDays        []int32         `protobuf:"varint,4,rep,packed,name=LoginDays,proto3" json:"LoginDays,omitempty"`                                                                                                 //连续登录
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_pb_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_pb_player_proto_rawDescGZIP(), []int{0}
}

func (x *Player) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Player) GetBackPack() map[int32]int32 {
	if x != nil {
		return x.BackPack
	}
	return nil
}

func (x *Player) GetCheckPointRecord() map[int32]int32 {
	if x != nil {
		return x.CheckPointRecord
	}
	return nil
}

func (x *Player) GetLoginDays() []int32 {
	if x != nil {
		return x.LoginDays
	}
	return nil
}

// 登录 没有直接创建
type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid       string `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	InviteUid string `protobuf:"bytes,2,opt,name=InviteUid,proto3" json:"InviteUid,omitempty"` //邀请者uid
	Avatar    string `protobuf:"bytes,3,opt,name=Avatar,proto3" json:"Avatar,omitempty"`       //微信头像
	Name      string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`           //微信名称
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_pb_player_proto_rawDescGZIP(), []int{1}
}

func (x *LoginReq) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *LoginReq) GetInviteUid() string {
	if x != nil {
		return x.InviteUid
	}
	return ""
}

func (x *LoginReq) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *LoginReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type LoginRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       int32   `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	PlayerInfo *Player `protobuf:"bytes,2,opt,name=PlayerInfo,proto3" json:"PlayerInfo,omitempty"` //玩家信息
}

func (x *LoginRsp) Reset() {
	*x = LoginRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRsp) ProtoMessage() {}

func (x *LoginRsp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRsp.ProtoReflect.Descriptor instead.
func (*LoginRsp) Descriptor() ([]byte, []int) {
	return file_pb_player_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LoginRsp) GetPlayerInfo() *Player {
	if x != nil {
		return x.PlayerInfo
	}
	return nil
}

// 一键领取七天签到奖励
type SevenDaysRewardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SevenDaysRewardReq) Reset() {
	*x = SevenDaysRewardReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_player_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SevenDaysRewardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SevenDaysRewardReq) ProtoMessage() {}

func (x *SevenDaysRewardReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_player_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SevenDaysRewardReq.ProtoReflect.Descriptor instead.
func (*SevenDaysRewardReq) Descriptor() ([]byte, []int) {
	return file_pb_player_proto_rawDescGZIP(), []int{3}
}

type SevenDaysRewardRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32           `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Rewards map[int32]int32 `protobuf:"bytes,2,rep,name=Rewards,proto3" json:"Rewards,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *SevenDaysRewardRsp) Reset() {
	*x = SevenDaysRewardRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_player_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SevenDaysRewardRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SevenDaysRewardRsp) ProtoMessage() {}

func (x *SevenDaysRewardRsp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_player_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SevenDaysRewardRsp.ProtoReflect.Descriptor instead.
func (*SevenDaysRewardRsp) Descriptor() ([]byte, []int) {
	return file_pb_player_proto_rawDescGZIP(), []int{4}
}

func (x *SevenDaysRewardRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SevenDaysRewardRsp) GetRewards() map[int32]int32 {
	if x != nil {
		return x.Rewards
	}
	return nil
}

type RankStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    string `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Avatar string `protobuf:"bytes,2,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Stars  int32  `protobuf:"varint,4,opt,name=Stars,proto3" json:"Stars,omitempty"` //星星数量
}

func (x *RankStruct) Reset() {
	*x = RankStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_player_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankStruct) ProtoMessage() {}

func (x *RankStruct) ProtoReflect() protoreflect.Message {
	mi := &file_pb_player_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankStruct.ProtoReflect.Descriptor instead.
func (*RankStruct) Descriptor() ([]byte, []int) {
	return file_pb_player_proto_rawDescGZIP(), []int{5}
}

func (x *RankStruct) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *RankStruct) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *RankStruct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RankStruct) GetStars() int32 {
	if x != nil {
		return x.Stars
	}
	return 0
}

// 星星总数排行榜
type StarsRankListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (x *StarsRankListReq) Reset() {
	*x = StarsRankListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_player_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StarsRankListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StarsRankListReq) ProtoMessage() {}

func (x *StarsRankListReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_player_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StarsRankListReq.ProtoReflect.Descriptor instead.
func (*StarsRankListReq) Descriptor() ([]byte, []int) {
	return file_pb_player_proto_rawDescGZIP(), []int{6}
}

func (x *StarsRankListReq) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type StarsRankListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int32         `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	RankList []*RankStruct `protobuf:"bytes,2,rep,name=RankList,proto3" json:"RankList,omitempty"`
}

func (x *StarsRankListRsp) Reset() {
	*x = StarsRankListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_player_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StarsRankListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StarsRankListRsp) ProtoMessage() {}

func (x *StarsRankListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_player_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StarsRankListRsp.ProtoReflect.Descriptor instead.
func (*StarsRankListRsp) Descriptor() ([]byte, []int) {
	return file_pb_player_proto_rawDescGZIP(), []int{7}
}

func (x *StarsRankListRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *StarsRankListRsp) GetRankList() []*RankStruct {
	if x != nil {
		return x.RankList
	}
	return nil
}

var File_pb_player_proto protoreflect.FileDescriptor

var file_pb_player_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x62, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xbe, 0x02, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55,
	0x69, 0x64, 0x12, 0x34, 0x0a, 0x08, 0x42, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x63, 0x6b, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x2e, 0x42, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x42, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x63, 0x6b, 0x12, 0x4c, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44,
	0x61, 0x79, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x44, 0x61, 0x79, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x42, 0x61, 0x63, 0x6b, 0x50, 0x61, 0x63, 0x6b,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x43, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x66, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x55, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4a,
	0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2a,
	0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x0a,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x65,
	0x76, 0x65, 0x6e, 0x44, 0x61, 0x79, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x22, 0xa3, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x76, 0x65, 0x6e, 0x44, 0x61, 0x79, 0x73, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x6e, 0x44, 0x61, 0x79, 0x73, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x52, 0x73, 0x70, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x60, 0x0a, 0x0a, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x55, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x53, 0x74, 0x61, 0x72, 0x73, 0x22, 0x24, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x72,
	0x73, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03,
	0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x4e, 0x75, 0x6d, 0x22, 0x52,
	0x0a, 0x10, 0x53, 0x74, 0x61, 0x72, 0x73, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x61,
	0x6e, 0x6b, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pb_player_proto_rawDescOnce sync.Once
	file_pb_player_proto_rawDescData = file_pb_player_proto_rawDesc
)

func file_pb_player_proto_rawDescGZIP() []byte {
	file_pb_player_proto_rawDescOnce.Do(func() {
		file_pb_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_player_proto_rawDescData)
	})
	return file_pb_player_proto_rawDescData
}

var file_pb_player_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pb_player_proto_goTypes = []interface{}{
	(*Player)(nil),             // 0: pb.Player
	(*LoginReq)(nil),           // 1: pb.LoginReq
	(*LoginRsp)(nil),           // 2: pb.LoginRsp
	(*SevenDaysRewardReq)(nil), // 3: pb.SevenDaysRewardReq
	(*SevenDaysRewardRsp)(nil), // 4: pb.SevenDaysRewardRsp
	(*RankStruct)(nil),         // 5: pb.RankStruct
	(*StarsRankListReq)(nil),   // 6: pb.StarsRankListReq
	(*StarsRankListRsp)(nil),   // 7: pb.StarsRankListRsp
	nil,                        // 8: pb.Player.BackPackEntry
	nil,                        // 9: pb.Player.CheckPointRecordEntry
	nil,                        // 10: pb.SevenDaysRewardRsp.RewardsEntry
}
var file_pb_player_proto_depIdxs = []int32{
	8,  // 0: pb.Player.BackPack:type_name -> pb.Player.BackPackEntry
	9,  // 1: pb.Player.CheckPointRecord:type_name -> pb.Player.CheckPointRecordEntry
	0,  // 2: pb.LoginRsp.PlayerInfo:type_name -> pb.Player
	10, // 3: pb.SevenDaysRewardRsp.Rewards:type_name -> pb.SevenDaysRewardRsp.RewardsEntry
	5,  // 4: pb.StarsRankListRsp.RankList:type_name -> pb.RankStruct
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_pb_player_proto_init() }
func file_pb_player_proto_init() {
	if File_pb_player_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_pb_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_pb_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRsp); i {
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
		file_pb_player_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SevenDaysRewardReq); i {
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
		file_pb_player_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SevenDaysRewardRsp); i {
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
		file_pb_player_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankStruct); i {
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
		file_pb_player_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StarsRankListReq); i {
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
		file_pb_player_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StarsRankListRsp); i {
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
			RawDescriptor: file_pb_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_player_proto_goTypes,
		DependencyIndexes: file_pb_player_proto_depIdxs,
		MessageInfos:      file_pb_player_proto_msgTypes,
	}.Build()
	File_pb_player_proto = out.File
	file_pb_player_proto_rawDesc = nil
	file_pb_player_proto_goTypes = nil
	file_pb_player_proto_depIdxs = nil
}
