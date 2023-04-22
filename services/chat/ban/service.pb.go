// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: chat/ban/service.proto

package ban

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

type FetchChatBanForUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ChatId string `protobuf:"bytes,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *FetchChatBanForUserRequest) Reset() {
	*x = FetchChatBanForUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_ban_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchChatBanForUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchChatBanForUserRequest) ProtoMessage() {}

func (x *FetchChatBanForUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_ban_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchChatBanForUserRequest.ProtoReflect.Descriptor instead.
func (*FetchChatBanForUserRequest) Descriptor() ([]byte, []int) {
	return file_chat_ban_service_proto_rawDescGZIP(), []int{0}
}

func (x *FetchChatBanForUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FetchChatBanForUserRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type FetchChatBansForUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ChatId string `protobuf:"bytes,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Cursor string `protobuf:"bytes,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *FetchChatBansForUserRequest) Reset() {
	*x = FetchChatBansForUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_ban_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchChatBansForUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchChatBansForUserRequest) ProtoMessage() {}

func (x *FetchChatBansForUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_ban_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchChatBansForUserRequest.ProtoReflect.Descriptor instead.
func (*FetchChatBansForUserRequest) Descriptor() ([]byte, []int) {
	return file_chat_ban_service_proto_rawDescGZIP(), []int{1}
}

func (x *FetchChatBansForUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FetchChatBansForUserRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *FetchChatBansForUserRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

type FetchPlatformBanForUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *FetchPlatformBanForUserRequest) Reset() {
	*x = FetchPlatformBanForUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_ban_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchPlatformBanForUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchPlatformBanForUserRequest) ProtoMessage() {}

func (x *FetchPlatformBanForUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_ban_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchPlatformBanForUserRequest.ProtoReflect.Descriptor instead.
func (*FetchPlatformBanForUserRequest) Descriptor() ([]byte, []int) {
	return file_chat_ban_service_proto_rawDescGZIP(), []int{2}
}

func (x *FetchPlatformBanForUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type FetchPlatformBansForUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Cursor string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *FetchPlatformBansForUserRequest) Reset() {
	*x = FetchPlatformBansForUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_ban_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchPlatformBansForUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchPlatformBansForUserRequest) ProtoMessage() {}

func (x *FetchPlatformBansForUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_ban_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchPlatformBansForUserRequest.ProtoReflect.Descriptor instead.
func (*FetchPlatformBansForUserRequest) Descriptor() ([]byte, []int) {
	return file_chat_ban_service_proto_rawDescGZIP(), []int{3}
}

func (x *FetchPlatformBansForUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FetchPlatformBansForUserRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

type ChatBan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BannerId  int64  `protobuf:"varint,2,opt,name=banner_id,json=bannerId,proto3" json:"banner_id,omitempty"`
	BannedId  int64  `protobuf:"varint,3,opt,name=banned_id,json=bannedId,proto3" json:"banned_id,omitempty"`
	ChannelId int64  `protobuf:"varint,4,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Reason    string `protobuf:"bytes,5,opt,name=reason,proto3" json:"reason,omitempty"`
	Type      string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Permanent bool   `protobuf:"varint,7,opt,name=permanent,proto3" json:"permanent,omitempty"`
	CreatedAt int64  `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ExpiresAt int64  `protobuf:"varint,9,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *ChatBan) Reset() {
	*x = ChatBan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_ban_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatBan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatBan) ProtoMessage() {}

func (x *ChatBan) ProtoReflect() protoreflect.Message {
	mi := &file_chat_ban_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatBan.ProtoReflect.Descriptor instead.
func (*ChatBan) Descriptor() ([]byte, []int) {
	return file_chat_ban_service_proto_rawDescGZIP(), []int{4}
}

func (x *ChatBan) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChatBan) GetBannerId() int64 {
	if x != nil {
		return x.BannerId
	}
	return 0
}

func (x *ChatBan) GetBannedId() int64 {
	if x != nil {
		return x.BannedId
	}
	return 0
}

func (x *ChatBan) GetChannelId() int64 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *ChatBan) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *ChatBan) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ChatBan) GetPermanent() bool {
	if x != nil {
		return x.Permanent
	}
	return false
}

func (x *ChatBan) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ChatBan) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

type ChatBans struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bans       []*ChatBan `protobuf:"bytes,1,rep,name=bans,proto3" json:"bans,omitempty"`
	NextCursor string     `protobuf:"bytes,2,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
}

func (x *ChatBans) Reset() {
	*x = ChatBans{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_ban_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatBans) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatBans) ProtoMessage() {}

func (x *ChatBans) ProtoReflect() protoreflect.Message {
	mi := &file_chat_ban_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatBans.ProtoReflect.Descriptor instead.
func (*ChatBans) Descriptor() ([]byte, []int) {
	return file_chat_ban_service_proto_rawDescGZIP(), []int{5}
}

func (x *ChatBans) GetBans() []*ChatBan {
	if x != nil {
		return x.Bans
	}
	return nil
}

func (x *ChatBans) GetNextCursor() string {
	if x != nil {
		return x.NextCursor
	}
	return ""
}

type PlatformBan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BannerId  int64  `protobuf:"varint,2,opt,name=banner_id,json=bannerId,proto3" json:"banner_id,omitempty"`
	BannedId  int64  `protobuf:"varint,3,opt,name=banned_id,json=bannedId,proto3" json:"banned_id,omitempty"`
	Reason    string `protobuf:"bytes,5,opt,name=reason,proto3" json:"reason,omitempty"`
	Type      string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Permanent bool   `protobuf:"varint,7,opt,name=permanent,proto3" json:"permanent,omitempty"`
	CreatedAt int64  `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ExpiresAt int64  `protobuf:"varint,9,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *PlatformBan) Reset() {
	*x = PlatformBan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_ban_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformBan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformBan) ProtoMessage() {}

func (x *PlatformBan) ProtoReflect() protoreflect.Message {
	mi := &file_chat_ban_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformBan.ProtoReflect.Descriptor instead.
func (*PlatformBan) Descriptor() ([]byte, []int) {
	return file_chat_ban_service_proto_rawDescGZIP(), []int{6}
}

func (x *PlatformBan) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlatformBan) GetBannerId() int64 {
	if x != nil {
		return x.BannerId
	}
	return 0
}

func (x *PlatformBan) GetBannedId() int64 {
	if x != nil {
		return x.BannedId
	}
	return 0
}

func (x *PlatformBan) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *PlatformBan) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PlatformBan) GetPermanent() bool {
	if x != nil {
		return x.Permanent
	}
	return false
}

func (x *PlatformBan) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *PlatformBan) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

type PlatformBans struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bans       []*PlatformBan `protobuf:"bytes,1,rep,name=bans,proto3" json:"bans,omitempty"`
	NextCursor string         `protobuf:"bytes,2,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
}

func (x *PlatformBans) Reset() {
	*x = PlatformBans{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_ban_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformBans) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformBans) ProtoMessage() {}

func (x *PlatformBans) ProtoReflect() protoreflect.Message {
	mi := &file_chat_ban_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformBans.ProtoReflect.Descriptor instead.
func (*PlatformBans) Descriptor() ([]byte, []int) {
	return file_chat_ban_service_proto_rawDescGZIP(), []int{7}
}

func (x *PlatformBans) GetBans() []*PlatformBan {
	if x != nil {
		return x.Bans
	}
	return nil
}

func (x *PlatformBans) GetNextCursor() string {
	if x != nil {
		return x.NextCursor
	}
	return ""
}

var File_chat_ban_service_proto protoreflect.FileDescriptor

var file_chat_ban_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x62, 0x61, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6b, 0x69, 0x63, 0x6b, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x62, 0x61, 0x6e, 0x22, 0x4e, 0x0a, 0x1a, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x43, 0x68, 0x61, 0x74, 0x42, 0x61, 0x6e, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x67, 0x0a, 0x1b, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x43, 0x68, 0x61, 0x74, 0x42, 0x61, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72,
	0x22, 0x39, 0x0a, 0x1e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x42, 0x61, 0x6e, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x52, 0x0a, 0x1f, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x61, 0x6e, 0x73,
	0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22,
	0xfa, 0x01, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x74, 0x42, 0x61, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6e,
	0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x61, 0x6e,
	0x6e, 0x65, 0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x22, 0x57, 0x0a, 0x08,
	0x43, 0x68, 0x61, 0x74, 0x42, 0x61, 0x6e, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x62, 0x61, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x69, 0x63, 0x6b, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x62, 0x61, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x42, 0x61, 0x6e, 0x52, 0x04,
	0x62, 0x61, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x75, 0x72,
	0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x43,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0xdf, 0x01, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x42, 0x61, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x70, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x22, 0x5f, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x42, 0x61, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x62, 0x61, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x69, 0x63, 0x6b, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x62, 0x61, 0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x61,
	0x6e, 0x52, 0x04, 0x62, 0x61, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65,
	0x78, 0x74, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x32, 0x92, 0x03, 0x0a, 0x0a, 0x43, 0x68, 0x61,
	0x74, 0x42, 0x61, 0x6e, 0x53, 0x76, 0x63, 0x12, 0x58, 0x0a, 0x13, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x43, 0x68, 0x61, 0x74, 0x42, 0x61, 0x6e, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x29,
	0x2e, 0x6b, 0x69, 0x63, 0x6b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x62, 0x61, 0x6e, 0x2e, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x43, 0x68, 0x61, 0x74, 0x42, 0x61, 0x6e, 0x46, 0x6f, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6b, 0x69, 0x63, 0x6b,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x62, 0x61, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x42, 0x61,
	0x6e, 0x12, 0x5b, 0x0a, 0x14, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x68, 0x61, 0x74, 0x42, 0x61,
	0x6e, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2a, 0x2e, 0x6b, 0x69, 0x63, 0x6b,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x62, 0x61, 0x6e, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43,
	0x68, 0x61, 0x74, 0x42, 0x61, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6b, 0x69, 0x63, 0x6b, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x62, 0x61, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x42, 0x61, 0x6e, 0x73, 0x12, 0x64,
	0x0a, 0x17, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42,
	0x61, 0x6e, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2d, 0x2e, 0x6b, 0x69, 0x63, 0x6b,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x62, 0x61, 0x6e, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x61, 0x6e, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6b, 0x69, 0x63, 0x6b, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x62, 0x61, 0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x42, 0x61, 0x6e, 0x12, 0x67, 0x0a, 0x18, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x61, 0x6e, 0x53, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x2e, 0x2e, 0x6b, 0x69, 0x63, 0x6b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x62, 0x61, 0x6e,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x61,
	0x6e, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x6b, 0x69, 0x63, 0x6b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x62, 0x61, 0x6e,
	0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x61, 0x6e, 0x73, 0x42, 0x0a, 0x5a,
	0x08, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x62, 0x61, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_chat_ban_service_proto_rawDescOnce sync.Once
	file_chat_ban_service_proto_rawDescData = file_chat_ban_service_proto_rawDesc
)

func file_chat_ban_service_proto_rawDescGZIP() []byte {
	file_chat_ban_service_proto_rawDescOnce.Do(func() {
		file_chat_ban_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_ban_service_proto_rawDescData)
	})
	return file_chat_ban_service_proto_rawDescData
}

var file_chat_ban_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_chat_ban_service_proto_goTypes = []interface{}{
	(*FetchChatBanForUserRequest)(nil),      // 0: kick.chat.ban.FetchChatBanForUserRequest
	(*FetchChatBansForUserRequest)(nil),     // 1: kick.chat.ban.FetchChatBansForUserRequest
	(*FetchPlatformBanForUserRequest)(nil),  // 2: kick.chat.ban.FetchPlatformBanForUserRequest
	(*FetchPlatformBansForUserRequest)(nil), // 3: kick.chat.ban.FetchPlatformBansForUserRequest
	(*ChatBan)(nil),                         // 4: kick.chat.ban.ChatBan
	(*ChatBans)(nil),                        // 5: kick.chat.ban.ChatBans
	(*PlatformBan)(nil),                     // 6: kick.chat.ban.PlatformBan
	(*PlatformBans)(nil),                    // 7: kick.chat.ban.PlatformBans
}
var file_chat_ban_service_proto_depIdxs = []int32{
	4, // 0: kick.chat.ban.ChatBans.bans:type_name -> kick.chat.ban.ChatBan
	6, // 1: kick.chat.ban.PlatformBans.bans:type_name -> kick.chat.ban.PlatformBan
	0, // 2: kick.chat.ban.ChatBanSvc.FetchChatBanForUser:input_type -> kick.chat.ban.FetchChatBanForUserRequest
	1, // 3: kick.chat.ban.ChatBanSvc.FetchChatBansForUser:input_type -> kick.chat.ban.FetchChatBansForUserRequest
	2, // 4: kick.chat.ban.ChatBanSvc.FetchPlatformBanForUser:input_type -> kick.chat.ban.FetchPlatformBanForUserRequest
	3, // 5: kick.chat.ban.ChatBanSvc.FetchPlatformBanSForUser:input_type -> kick.chat.ban.FetchPlatformBansForUserRequest
	4, // 6: kick.chat.ban.ChatBanSvc.FetchChatBanForUser:output_type -> kick.chat.ban.ChatBan
	5, // 7: kick.chat.ban.ChatBanSvc.FetchChatBansForUser:output_type -> kick.chat.ban.ChatBans
	6, // 8: kick.chat.ban.ChatBanSvc.FetchPlatformBanForUser:output_type -> kick.chat.ban.PlatformBan
	7, // 9: kick.chat.ban.ChatBanSvc.FetchPlatformBanSForUser:output_type -> kick.chat.ban.PlatformBans
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_chat_ban_service_proto_init() }
func file_chat_ban_service_proto_init() {
	if File_chat_ban_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_ban_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchChatBanForUserRequest); i {
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
		file_chat_ban_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchChatBansForUserRequest); i {
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
		file_chat_ban_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchPlatformBanForUserRequest); i {
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
		file_chat_ban_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchPlatformBansForUserRequest); i {
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
		file_chat_ban_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatBan); i {
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
		file_chat_ban_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatBans); i {
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
		file_chat_ban_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformBan); i {
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
		file_chat_ban_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformBans); i {
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
			RawDescriptor: file_chat_ban_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_ban_service_proto_goTypes,
		DependencyIndexes: file_chat_ban_service_proto_depIdxs,
		MessageInfos:      file_chat_ban_service_proto_msgTypes,
	}.Build()
	File_chat_ban_service_proto = out.File
	file_chat_ban_service_proto_rawDesc = nil
	file_chat_ban_service_proto_goTypes = nil
	file_chat_ban_service_proto_depIdxs = nil
}