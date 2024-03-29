// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: profile/v1/profile.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type RegisterCelebRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *RegisterCelebRequest_User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *RegisterCelebRequest) Reset() {
	*x = RegisterCelebRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterCelebRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterCelebRequest) ProtoMessage() {}

func (x *RegisterCelebRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterCelebRequest.ProtoReflect.Descriptor instead.
func (*RegisterCelebRequest) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterCelebRequest) GetUser() *RegisterCelebRequest_User {
	if x != nil {
		return x.User
	}
	return nil
}

type ProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ProfileRequest) Reset() {
	*x = ProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileRequest) ProtoMessage() {}

func (x *ProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileRequest.ProtoReflect.Descriptor instead.
func (*ProfileRequest) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{1}
}

func (x *ProfileRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type FollowUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     *FollowUserRequest_User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Username string                  `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *FollowUserRequest) Reset() {
	*x = FollowUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowUserRequest) ProtoMessage() {}

func (x *FollowUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowUserRequest.ProtoReflect.Descriptor instead.
func (*FollowUserRequest) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{2}
}

func (x *FollowUserRequest) GetUser() *FollowUserRequest_User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *FollowUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ProfileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *ProfileReply_Data `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *ProfileReply) Reset() {
	*x = ProfileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileReply) ProtoMessage() {}

func (x *ProfileReply) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileReply.ProtoReflect.Descriptor instead.
func (*ProfileReply) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{3}
}

func (x *ProfileReply) GetProfile() *ProfileReply_Data {
	if x != nil {
		return x.Profile
	}
	return nil
}

type RegisterCelebRequest_User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *RegisterCelebRequest_User) Reset() {
	*x = RegisterCelebRequest_User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterCelebRequest_User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterCelebRequest_User) ProtoMessage() {}

func (x *RegisterCelebRequest_User) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterCelebRequest_User.ProtoReflect.Descriptor instead.
func (*RegisterCelebRequest_User) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RegisterCelebRequest_User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterCelebRequest_User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterCelebRequest_User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type FollowUserRequest_User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *FollowUserRequest_User) Reset() {
	*x = FollowUserRequest_User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowUserRequest_User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowUserRequest_User) ProtoMessage() {}

func (x *FollowUserRequest_User) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowUserRequest_User.ProtoReflect.Descriptor instead.
func (*FollowUserRequest_User) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{2, 0}
}

func (x *FollowUserRequest_User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ProfileReply_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Bio       string `protobuf:"bytes,4,opt,name=bio,proto3" json:"bio,omitempty"`
	Image     string `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Following bool   `protobuf:"varint,6,opt,name=following,proto3" json:"following,omitempty"`
}

func (x *ProfileReply_Data) Reset() {
	*x = ProfileReply_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileReply_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileReply_Data) ProtoMessage() {}

func (x *ProfileReply_Data) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileReply_Data.ProtoReflect.Descriptor instead.
func (*ProfileReply_Data) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ProfileReply_Data) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ProfileReply_Data) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *ProfileReply_Data) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *ProfileReply_Data) GetFollowing() bool {
	if x != nil {
		return x.Following
	}
	return false
}

var File_profile_v1_profile_proto protoreflect.FileDescriptor

var file_profile_v1_profile_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x65, 0x6c, 0x65, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x65, 0x6c, 0x65, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x1a, 0x54, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x11, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x1a, 0x1c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x22, 0xb5, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x3b, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x68,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x62, 0x69, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x32, 0xd5, 0x03, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x6a, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x43, 0x65, 0x6c, 0x65, 0x62, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43,
	0x65, 0x6c, 0x65, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0f, 0x22, 0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a,
	0x12, 0x6c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x79,
	0x0a, 0x0a, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x24, 0x22, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x3a, 0x01, 0x2a, 0x12, 0x75, 0x0a, 0x0c, 0x55, 0x6e, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x2a,
	0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x7b,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x42, 0x2f, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x1b, 0x72, 0x65, 0x61, 0x6c, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_profile_v1_profile_proto_rawDescOnce sync.Once
	file_profile_v1_profile_proto_rawDescData = file_profile_v1_profile_proto_rawDesc
)

func file_profile_v1_profile_proto_rawDescGZIP() []byte {
	file_profile_v1_profile_proto_rawDescOnce.Do(func() {
		file_profile_v1_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_profile_v1_profile_proto_rawDescData)
	})
	return file_profile_v1_profile_proto_rawDescData
}

var file_profile_v1_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_profile_v1_profile_proto_goTypes = []interface{}{
	(*RegisterCelebRequest)(nil),      // 0: api.profile.v1.RegisterCelebRequest
	(*ProfileRequest)(nil),            // 1: api.profile.v1.ProfileRequest
	(*FollowUserRequest)(nil),         // 2: api.profile.v1.FollowUserRequest
	(*ProfileReply)(nil),              // 3: api.profile.v1.ProfileReply
	(*RegisterCelebRequest_User)(nil), // 4: api.profile.v1.RegisterCelebRequest.User
	(*FollowUserRequest_User)(nil),    // 5: api.profile.v1.FollowUserRequest.User
	(*ProfileReply_Data)(nil),         // 6: api.profile.v1.ProfileReply.Data
}
var file_profile_v1_profile_proto_depIdxs = []int32{
	4, // 0: api.profile.v1.RegisterCelebRequest.user:type_name -> api.profile.v1.RegisterCelebRequest.User
	5, // 1: api.profile.v1.FollowUserRequest.user:type_name -> api.profile.v1.FollowUserRequest.User
	6, // 2: api.profile.v1.ProfileReply.profile:type_name -> api.profile.v1.ProfileReply.Data
	0, // 3: api.profile.v1.Profile.RegisterCeleb:input_type -> api.profile.v1.RegisterCelebRequest
	1, // 4: api.profile.v1.Profile.GetProfile:input_type -> api.profile.v1.ProfileRequest
	2, // 5: api.profile.v1.Profile.FollowUser:input_type -> api.profile.v1.FollowUserRequest
	1, // 6: api.profile.v1.Profile.UnfollowUser:input_type -> api.profile.v1.ProfileRequest
	3, // 7: api.profile.v1.Profile.RegisterCeleb:output_type -> api.profile.v1.ProfileReply
	3, // 8: api.profile.v1.Profile.GetProfile:output_type -> api.profile.v1.ProfileReply
	3, // 9: api.profile.v1.Profile.FollowUser:output_type -> api.profile.v1.ProfileReply
	3, // 10: api.profile.v1.Profile.UnfollowUser:output_type -> api.profile.v1.ProfileReply
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_profile_v1_profile_proto_init() }
func file_profile_v1_profile_proto_init() {
	if File_profile_v1_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_profile_v1_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterCelebRequest); i {
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
		file_profile_v1_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileRequest); i {
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
		file_profile_v1_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowUserRequest); i {
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
		file_profile_v1_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileReply); i {
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
		file_profile_v1_profile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterCelebRequest_User); i {
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
		file_profile_v1_profile_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowUserRequest_User); i {
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
		file_profile_v1_profile_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileReply_Data); i {
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
			RawDescriptor: file_profile_v1_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profile_v1_profile_proto_goTypes,
		DependencyIndexes: file_profile_v1_profile_proto_depIdxs,
		MessageInfos:      file_profile_v1_profile_proto_msgTypes,
	}.Build()
	File_profile_v1_profile_proto = out.File
	file_profile_v1_profile_proto_rawDesc = nil
	file_profile_v1_profile_proto_goTypes = nil
	file_profile_v1_profile_proto_depIdxs = nil
}
