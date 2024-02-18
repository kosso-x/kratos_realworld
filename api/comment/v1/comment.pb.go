// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: comment/v1/comment.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddCommentsToAnArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *AddCommentsToAnArticleRequest_CommentRequest `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
	Slug    string                                        `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (x *AddCommentsToAnArticleRequest) Reset() {
	*x = AddCommentsToAnArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_v1_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCommentsToAnArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCommentsToAnArticleRequest) ProtoMessage() {}

func (x *AddCommentsToAnArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_v1_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCommentsToAnArticleRequest.ProtoReflect.Descriptor instead.
func (*AddCommentsToAnArticleRequest) Descriptor() ([]byte, []int) {
	return file_comment_v1_comment_proto_rawDescGZIP(), []int{0}
}

func (x *AddCommentsToAnArticleRequest) GetComment() *AddCommentsToAnArticleRequest_CommentRequest {
	if x != nil {
		return x.Comment
	}
	return nil
}

func (x *AddCommentsToAnArticleRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type CommentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *CommentReply_Comment `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CommentReply) Reset() {
	*x = CommentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_v1_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentReply) ProtoMessage() {}

func (x *CommentReply) ProtoReflect() protoreflect.Message {
	mi := &file_comment_v1_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentReply.ProtoReflect.Descriptor instead.
func (*CommentReply) Descriptor() ([]byte, []int) {
	return file_comment_v1_comment_proto_rawDescGZIP(), []int{1}
}

func (x *CommentReply) GetComment() *CommentReply_Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type GetCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (x *GetCommentsRequest) Reset() {
	*x = GetCommentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_v1_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsRequest) ProtoMessage() {}

func (x *GetCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_v1_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsRequest.ProtoReflect.Descriptor instead.
func (*GetCommentsRequest) Descriptor() ([]byte, []int) {
	return file_comment_v1_comment_proto_rawDescGZIP(), []int{2}
}

func (x *GetCommentsRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type MultipleComments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*MultipleComments_Comments `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *MultipleComments) Reset() {
	*x = MultipleComments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_v1_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultipleComments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleComments) ProtoMessage() {}

func (x *MultipleComments) ProtoReflect() protoreflect.Message {
	mi := &file_comment_v1_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipleComments.ProtoReflect.Descriptor instead.
func (*MultipleComments) Descriptor() ([]byte, []int) {
	return file_comment_v1_comment_proto_rawDescGZIP(), []int{3}
}

func (x *MultipleComments) GetComments() []*MultipleComments_Comments {
	if x != nil {
		return x.Comments
	}
	return nil
}

type DeleteCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	Id   int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCommentRequest) Reset() {
	*x = DeleteCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_v1_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentRequest) ProtoMessage() {}

func (x *DeleteCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_v1_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentRequest.ProtoReflect.Descriptor instead.
func (*DeleteCommentRequest) Descriptor() ([]byte, []int) {
	return file_comment_v1_comment_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteCommentRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *DeleteCommentRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteCommentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCommentReply) Reset() {
	*x = DeleteCommentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_v1_comment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentReply) ProtoMessage() {}

func (x *DeleteCommentReply) ProtoReflect() protoreflect.Message {
	mi := &file_comment_v1_comment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentReply.ProtoReflect.Descriptor instead.
func (*DeleteCommentReply) Descriptor() ([]byte, []int) {
	return file_comment_v1_comment_proto_rawDescGZIP(), []int{5}
}

type AddCommentsToAnArticleRequest_CommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *AddCommentsToAnArticleRequest_CommentRequest) Reset() {
	*x = AddCommentsToAnArticleRequest_CommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_v1_comment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCommentsToAnArticleRequest_CommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCommentsToAnArticleRequest_CommentRequest) ProtoMessage() {}

func (x *AddCommentsToAnArticleRequest_CommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_v1_comment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCommentsToAnArticleRequest_CommentRequest.ProtoReflect.Descriptor instead.
func (*AddCommentsToAnArticleRequest_CommentRequest) Descriptor() ([]byte, []int) {
	return file_comment_v1_comment_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AddCommentsToAnArticleRequest_CommentRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type CommentReply_Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Bio       string `protobuf:"bytes,2,opt,name=bio,proto3" json:"bio,omitempty"`
	Image     string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Following bool   `protobuf:"varint,4,opt,name=following,proto3" json:"following,omitempty"`
}

func (x *CommentReply_Author) Reset() {
	*x = CommentReply_Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_v1_comment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentReply_Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentReply_Author) ProtoMessage() {}

func (x *CommentReply_Author) ProtoReflect() protoreflect.Message {
	mi := &file_comment_v1_comment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentReply_Author.ProtoReflect.Descriptor instead.
func (*CommentReply_Author) Descriptor() ([]byte, []int) {
	return file_comment_v1_comment_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CommentReply_Author) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CommentReply_Author) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *CommentReply_Author) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CommentReply_Author) GetFollowing() bool {
	if x != nil {
		return x.Following
	}
	return false
}

type CommentReply_Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Body      string                 `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Author    *CommentReply_Author   `protobuf:"bytes,5,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *CommentReply_Comment) Reset() {
	*x = CommentReply_Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_v1_comment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentReply_Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentReply_Comment) ProtoMessage() {}

func (x *CommentReply_Comment) ProtoReflect() protoreflect.Message {
	mi := &file_comment_v1_comment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentReply_Comment.ProtoReflect.Descriptor instead.
func (*CommentReply_Comment) Descriptor() ([]byte, []int) {
	return file_comment_v1_comment_proto_rawDescGZIP(), []int{1, 1}
}

func (x *CommentReply_Comment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommentReply_Comment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CommentReply_Comment) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *CommentReply_Comment) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *CommentReply_Comment) GetAuthor() *CommentReply_Author {
	if x != nil {
		return x.Author
	}
	return nil
}

type MultipleComments_Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Bio       string `protobuf:"bytes,2,opt,name=bio,proto3" json:"bio,omitempty"`
	Image     string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Following bool   `protobuf:"varint,4,opt,name=following,proto3" json:"following,omitempty"`
}

func (x *MultipleComments_Author) Reset() {
	*x = MultipleComments_Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_v1_comment_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultipleComments_Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleComments_Author) ProtoMessage() {}

func (x *MultipleComments_Author) ProtoReflect() protoreflect.Message {
	mi := &file_comment_v1_comment_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipleComments_Author.ProtoReflect.Descriptor instead.
func (*MultipleComments_Author) Descriptor() ([]byte, []int) {
	return file_comment_v1_comment_proto_rawDescGZIP(), []int{3, 0}
}

func (x *MultipleComments_Author) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *MultipleComments_Author) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *MultipleComments_Author) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *MultipleComments_Author) GetFollowing() bool {
	if x != nil {
		return x.Following
	}
	return false
}

type MultipleComments_Comments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt *timestamppb.Timestamp   `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp   `protobuf:"bytes,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Body      string                   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Author    *MultipleComments_Author `protobuf:"bytes,5,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *MultipleComments_Comments) Reset() {
	*x = MultipleComments_Comments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_v1_comment_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultipleComments_Comments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleComments_Comments) ProtoMessage() {}

func (x *MultipleComments_Comments) ProtoReflect() protoreflect.Message {
	mi := &file_comment_v1_comment_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipleComments_Comments.ProtoReflect.Descriptor instead.
func (*MultipleComments_Comments) Descriptor() ([]byte, []int) {
	return file_comment_v1_comment_proto_rawDescGZIP(), []int{3, 1}
}

func (x *MultipleComments_Comments) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MultipleComments_Comments) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *MultipleComments_Comments) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *MultipleComments_Comments) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *MultipleComments_Comments) GetAuthor() *MultipleComments_Author {
	if x != nil {
		return x.Author
	}
	return nil
}

var File_comment_v1_comment_proto protoreflect.FileDescriptor

var file_comment_v1_comment_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x1d, 0x41, 0x64,
	0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x54, 0x6f, 0x41, 0x6e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x56, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x54, 0x6f, 0x41, 0x6e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x1a, 0x24, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x9b, 0x03,
	0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3e,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x6a,
	0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x1a, 0xde, 0x01, 0x0a, 0x07, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x3b,
	0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x28, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x6c, 0x75, 0x67, 0x22, 0xab, 0x03, 0x0a, 0x10, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x45, 0x0a, 0x08, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x1a, 0x6a, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x1a, 0xe3, 0x01,
	0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x3f, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x22, 0x3a, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xad, 0x03, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x8f, 0x01, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x54, 0x6f, 0x41, 0x6e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x2d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x54, 0x6f, 0x41, 0x6e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x22, 0x22, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x2f, 0x7b, 0x73, 0x6c, 0x75, 0x67, 0x7d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x3a, 0x01, 0x2a, 0x12, 0x87, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x6e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x73,
	0x6c, 0x75, 0x67, 0x7d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x85, 0x01,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x24, 0x2a, 0x22, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x2f, 0x7b, 0x73, 0x6c, 0x75, 0x67, 0x7d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x2f, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1b, 0x72, 0x65, 0x61, 0x6c, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comment_v1_comment_proto_rawDescOnce sync.Once
	file_comment_v1_comment_proto_rawDescData = file_comment_v1_comment_proto_rawDesc
)

func file_comment_v1_comment_proto_rawDescGZIP() []byte {
	file_comment_v1_comment_proto_rawDescOnce.Do(func() {
		file_comment_v1_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_comment_v1_comment_proto_rawDescData)
	})
	return file_comment_v1_comment_proto_rawDescData
}

var file_comment_v1_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_comment_v1_comment_proto_goTypes = []interface{}{
	(*AddCommentsToAnArticleRequest)(nil),                // 0: api.comment.v1.AddCommentsToAnArticleRequest
	(*CommentReply)(nil),                                 // 1: api.comment.v1.CommentReply
	(*GetCommentsRequest)(nil),                           // 2: api.comment.v1.GetCommentsRequest
	(*MultipleComments)(nil),                             // 3: api.comment.v1.MultipleComments
	(*DeleteCommentRequest)(nil),                         // 4: api.comment.v1.DeleteCommentRequest
	(*DeleteCommentReply)(nil),                           // 5: api.comment.v1.DeleteCommentReply
	(*AddCommentsToAnArticleRequest_CommentRequest)(nil), // 6: api.comment.v1.AddCommentsToAnArticleRequest.CommentRequest
	(*CommentReply_Author)(nil),                          // 7: api.comment.v1.CommentReply.Author
	(*CommentReply_Comment)(nil),                         // 8: api.comment.v1.CommentReply.Comment
	(*MultipleComments_Author)(nil),                      // 9: api.comment.v1.MultipleComments.Author
	(*MultipleComments_Comments)(nil),                    // 10: api.comment.v1.MultipleComments.Comments
	(*timestamppb.Timestamp)(nil),                        // 11: google.protobuf.Timestamp
}
var file_comment_v1_comment_proto_depIdxs = []int32{
	6,  // 0: api.comment.v1.AddCommentsToAnArticleRequest.comment:type_name -> api.comment.v1.AddCommentsToAnArticleRequest.CommentRequest
	8,  // 1: api.comment.v1.CommentReply.comment:type_name -> api.comment.v1.CommentReply.Comment
	10, // 2: api.comment.v1.MultipleComments.comments:type_name -> api.comment.v1.MultipleComments.Comments
	11, // 3: api.comment.v1.CommentReply.Comment.createdAt:type_name -> google.protobuf.Timestamp
	11, // 4: api.comment.v1.CommentReply.Comment.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 5: api.comment.v1.CommentReply.Comment.author:type_name -> api.comment.v1.CommentReply.Author
	11, // 6: api.comment.v1.MultipleComments.Comments.createdAt:type_name -> google.protobuf.Timestamp
	11, // 7: api.comment.v1.MultipleComments.Comments.updatedAt:type_name -> google.protobuf.Timestamp
	9,  // 8: api.comment.v1.MultipleComments.Comments.author:type_name -> api.comment.v1.MultipleComments.Author
	0,  // 9: api.comment.v1.Comment.AddCommentsToAnArticle:input_type -> api.comment.v1.AddCommentsToAnArticleRequest
	2,  // 10: api.comment.v1.Comment.GetCommentsFromAnArticle:input_type -> api.comment.v1.GetCommentsRequest
	4,  // 11: api.comment.v1.Comment.DeleteComment:input_type -> api.comment.v1.DeleteCommentRequest
	1,  // 12: api.comment.v1.Comment.AddCommentsToAnArticle:output_type -> api.comment.v1.CommentReply
	3,  // 13: api.comment.v1.Comment.GetCommentsFromAnArticle:output_type -> api.comment.v1.MultipleComments
	5,  // 14: api.comment.v1.Comment.DeleteComment:output_type -> api.comment.v1.DeleteCommentReply
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_comment_v1_comment_proto_init() }
func file_comment_v1_comment_proto_init() {
	if File_comment_v1_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comment_v1_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCommentsToAnArticleRequest); i {
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
		file_comment_v1_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentReply); i {
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
		file_comment_v1_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentsRequest); i {
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
		file_comment_v1_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultipleComments); i {
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
		file_comment_v1_comment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentRequest); i {
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
		file_comment_v1_comment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentReply); i {
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
		file_comment_v1_comment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCommentsToAnArticleRequest_CommentRequest); i {
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
		file_comment_v1_comment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentReply_Author); i {
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
		file_comment_v1_comment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentReply_Comment); i {
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
		file_comment_v1_comment_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultipleComments_Author); i {
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
		file_comment_v1_comment_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultipleComments_Comments); i {
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
			RawDescriptor: file_comment_v1_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comment_v1_comment_proto_goTypes,
		DependencyIndexes: file_comment_v1_comment_proto_depIdxs,
		MessageInfos:      file_comment_v1_comment_proto_msgTypes,
	}.Build()
	File_comment_v1_comment_proto = out.File
	file_comment_v1_comment_proto_rawDesc = nil
	file_comment_v1_comment_proto_goTypes = nil
	file_comment_v1_comment_proto_depIdxs = nil
}
