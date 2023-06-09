// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: chat/identity/service.proto

package identity

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

type GetIdentityForChannelAndUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetIdentityForChannelAndUserRequest) Reset() {
	*x = GetIdentityForChannelAndUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_identity_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIdentityForChannelAndUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdentityForChannelAndUserRequest) ProtoMessage() {}

func (x *GetIdentityForChannelAndUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_identity_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdentityForChannelAndUserRequest.ProtoReflect.Descriptor instead.
func (*GetIdentityForChannelAndUserRequest) Descriptor() ([]byte, []int) {
	return file_chat_identity_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetIdentityForChannelAndUserRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *GetIdentityForChannelAndUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UpdateIdentityForChannelAndUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string   `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	UserId    string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Color     string   `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	Badges    []*Badge `protobuf:"bytes,4,rep,name=badges,proto3" json:"badges,omitempty"`
}

func (x *UpdateIdentityForChannelAndUserRequest) Reset() {
	*x = UpdateIdentityForChannelAndUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_identity_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIdentityForChannelAndUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIdentityForChannelAndUserRequest) ProtoMessage() {}

func (x *UpdateIdentityForChannelAndUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_identity_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIdentityForChannelAndUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateIdentityForChannelAndUserRequest) Descriptor() ([]byte, []int) {
	return file_chat_identity_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateIdentityForChannelAndUserRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *UpdateIdentityForChannelAndUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateIdentityForChannelAndUserRequest) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *UpdateIdentityForChannelAndUserRequest) GetBadges() []*Badge {
	if x != nil {
		return x.Badges
	}
	return nil
}

type ChatIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color  string   `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"`
	Badges []*Badge `protobuf:"bytes,2,rep,name=badges,proto3" json:"badges,omitempty"`
}

func (x *ChatIdentity) Reset() {
	*x = ChatIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_identity_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatIdentity) ProtoMessage() {}

func (x *ChatIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_chat_identity_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatIdentity.ProtoReflect.Descriptor instead.
func (*ChatIdentity) Descriptor() ([]byte, []int) {
	return file_chat_identity_service_proto_rawDescGZIP(), []int{2}
}

func (x *ChatIdentity) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *ChatIdentity) GetBadges() []*Badge {
	if x != nil {
		return x.Badges
	}
	return nil
}

type Badge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Badge) Reset() {
	*x = Badge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_identity_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Badge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Badge) ProtoMessage() {}

func (x *Badge) ProtoReflect() protoreflect.Message {
	mi := &file_chat_identity_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Badge.ProtoReflect.Descriptor instead.
func (*Badge) Descriptor() ([]byte, []int) {
	return file_chat_identity_service_proto_rawDescGZIP(), []int{3}
}

func (x *Badge) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_chat_identity_service_proto protoreflect.FileDescriptor

var file_chat_identity_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6b,
	0x69, 0x63, 0x6b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x22, 0x5d, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x46, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x6e, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0xa9, 0x01, 0x0a, 0x26, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x46, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x6e, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x31, 0x0a, 0x06, 0x62, 0x61, 0x64,
	0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6b, 0x69, 0x63, 0x6b,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x42,
	0x61, 0x64, 0x67, 0x65, 0x52, 0x06, 0x62, 0x61, 0x64, 0x67, 0x65, 0x73, 0x22, 0x57, 0x0a, 0x0c,
	0x43, 0x68, 0x61, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x12, 0x31, 0x0a, 0x06, 0x62, 0x61, 0x64, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6b, 0x69, 0x63, 0x6b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x42, 0x61, 0x64, 0x67, 0x65, 0x52, 0x06, 0x62,
	0x61, 0x64, 0x67, 0x65, 0x73, 0x22, 0x1b, 0x0a, 0x05, 0x42, 0x61, 0x64, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x32, 0x8f, 0x02, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x53, 0x76, 0x63, 0x12, 0x7b, 0x0a, 0x1e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x46, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x41, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x37, 0x2e, 0x6b, 0x69, 0x63, 0x6b, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x46, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x41, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x6b, 0x69, 0x63, 0x6b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x7f, 0x0a, 0x1f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x46, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41,
	0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3a, 0x2e, 0x6b, 0x69, 0x63, 0x6b, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x46, 0x6f, 0x72, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6b, 0x69, 0x63, 0x6b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x42, 0x0f, 0x5a, 0x0d, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_identity_service_proto_rawDescOnce sync.Once
	file_chat_identity_service_proto_rawDescData = file_chat_identity_service_proto_rawDesc
)

func file_chat_identity_service_proto_rawDescGZIP() []byte {
	file_chat_identity_service_proto_rawDescOnce.Do(func() {
		file_chat_identity_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_identity_service_proto_rawDescData)
	})
	return file_chat_identity_service_proto_rawDescData
}

var file_chat_identity_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chat_identity_service_proto_goTypes = []interface{}{
	(*GetIdentityForChannelAndUserRequest)(nil),    // 0: kick.chat.identity.GetIdentityForChannelAndUserRequest
	(*UpdateIdentityForChannelAndUserRequest)(nil), // 1: kick.chat.identity.UpdateIdentityForChannelAndUserRequest
	(*ChatIdentity)(nil),                           // 2: kick.chat.identity.ChatIdentity
	(*Badge)(nil),                                  // 3: kick.chat.identity.Badge
}
var file_chat_identity_service_proto_depIdxs = []int32{
	3, // 0: kick.chat.identity.UpdateIdentityForChannelAndUserRequest.badges:type_name -> kick.chat.identity.Badge
	3, // 1: kick.chat.identity.ChatIdentity.badges:type_name -> kick.chat.identity.Badge
	0, // 2: kick.chat.identity.ChatIdentitySvc.FetchIdentityForChannelAndUser:input_type -> kick.chat.identity.GetIdentityForChannelAndUserRequest
	1, // 3: kick.chat.identity.ChatIdentitySvc.UpdateIdentityForChannelAndUser:input_type -> kick.chat.identity.UpdateIdentityForChannelAndUserRequest
	2, // 4: kick.chat.identity.ChatIdentitySvc.FetchIdentityForChannelAndUser:output_type -> kick.chat.identity.ChatIdentity
	2, // 5: kick.chat.identity.ChatIdentitySvc.UpdateIdentityForChannelAndUser:output_type -> kick.chat.identity.ChatIdentity
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_chat_identity_service_proto_init() }
func file_chat_identity_service_proto_init() {
	if File_chat_identity_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_identity_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIdentityForChannelAndUserRequest); i {
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
		file_chat_identity_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIdentityForChannelAndUserRequest); i {
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
		file_chat_identity_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatIdentity); i {
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
		file_chat_identity_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Badge); i {
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
			RawDescriptor: file_chat_identity_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_identity_service_proto_goTypes,
		DependencyIndexes: file_chat_identity_service_proto_depIdxs,
		MessageInfos:      file_chat_identity_service_proto_msgTypes,
	}.Build()
	File_chat_identity_service_proto = out.File
	file_chat_identity_service_proto_rawDesc = nil
	file_chat_identity_service_proto_goTypes = nil
	file_chat_identity_service_proto_depIdxs = nil
}
