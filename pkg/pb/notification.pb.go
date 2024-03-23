// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/pb/notification.proto

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

type CommentedNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId      string `protobuf:"bytes,1,opt,name=postId,proto3" json:"postId,omitempty"`
	CommenterId string `protobuf:"bytes,2,opt,name=commenterId,proto3" json:"commenterId,omitempty"`
	Comment     string `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CommentedNotification) Reset() {
	*x = CommentedNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentedNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentedNotification) ProtoMessage() {}

func (x *CommentedNotification) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentedNotification.ProtoReflect.Descriptor instead.
func (*CommentedNotification) Descriptor() ([]byte, []int) {
	return file_pkg_pb_notification_proto_rawDescGZIP(), []int{0}
}

func (x *CommentedNotification) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *CommentedNotification) GetCommenterId() string {
	if x != nil {
		return x.CommenterId
	}
	return ""
}

func (x *CommentedNotification) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type FollowedNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowerId     string `protobuf:"bytes,1,opt,name=followerId,proto3" json:"followerId,omitempty"`
	FollowedUserId string `protobuf:"bytes,2,opt,name=followedUserId,proto3" json:"followedUserId,omitempty"`
}

func (x *FollowedNotification) Reset() {
	*x = FollowedNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowedNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowedNotification) ProtoMessage() {}

func (x *FollowedNotification) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowedNotification.ProtoReflect.Descriptor instead.
func (*FollowedNotification) Descriptor() ([]byte, []int) {
	return file_pkg_pb_notification_proto_rawDescGZIP(), []int{1}
}

func (x *FollowedNotification) GetFollowerId() string {
	if x != nil {
		return x.FollowerId
	}
	return ""
}

func (x *FollowedNotification) GetFollowedUserId() string {
	if x != nil {
		return x.FollowedUserId
	}
	return ""
}

// Kafka message for notifications
type KafkaNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KafkaNotification) Reset() {
	*x = KafkaNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KafkaNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KafkaNotification) ProtoMessage() {}

func (x *KafkaNotification) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KafkaNotification.ProtoReflect.Descriptor instead.
func (*KafkaNotification) Descriptor() ([]byte, []int) {
	return file_pkg_pb_notification_proto_rawDescGZIP(), []int{2}
}

func (x *KafkaNotification) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KafkaNotification) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// LikeNotification Request
type LikeNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId int64 `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	UserId int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *LikeNotification) Reset() {
	*x = LikeNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_notification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeNotification) ProtoMessage() {}

func (x *LikeNotification) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_notification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeNotification.ProtoReflect.Descriptor instead.
func (*LikeNotification) Descriptor() ([]byte, []int) {
	return file_pkg_pb_notification_proto_rawDescGZIP(), []int{3}
}

func (x *LikeNotification) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *LikeNotification) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type NotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *NotificationResponse) Reset() {
	*x = NotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_notification_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationResponse) ProtoMessage() {}

func (x *NotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_notification_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationResponse.ProtoReflect.Descriptor instead.
func (*NotificationResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_notification_proto_rawDescGZIP(), []int{4}
}

func (x *NotificationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pkg_pb_notification_proto protoreflect.FileDescriptor

var file_pkg_pb_notification_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6b, 0x0a, 0x15, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x5e, 0x0a, 0x14, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x11, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x44, 0x0a, 0x10, 0x4c, 0x69, 0x6b, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x14, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa1, 0x03, 0x0a, 0x13,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x19, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x23, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x22, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x18, 0x53,
	0x65, 0x6e, 0x64, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x22, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5e, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x22, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5c, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6b, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x22, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x67,
	0x6f, 0x6b, 0x75, 0x6c, 0x39, 0x36, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2d, 0x68, 0x75, 0x62,
	0x2d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x76,
	0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_notification_proto_rawDescOnce sync.Once
	file_pkg_pb_notification_proto_rawDescData = file_pkg_pb_notification_proto_rawDesc
)

func file_pkg_pb_notification_proto_rawDescGZIP() []byte {
	file_pkg_pb_notification_proto_rawDescOnce.Do(func() {
		file_pkg_pb_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_notification_proto_rawDescData)
	})
	return file_pkg_pb_notification_proto_rawDescData
}

var file_pkg_pb_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_pb_notification_proto_goTypes = []interface{}{
	(*CommentedNotification)(nil), // 0: notification.CommentedNotification
	(*FollowedNotification)(nil),  // 1: notification.FollowedNotification
	(*KafkaNotification)(nil),     // 2: notification.KafkaNotification
	(*LikeNotification)(nil),      // 3: notification.LikeNotification
	(*NotificationResponse)(nil),  // 4: notification.NotificationResponse
}
var file_pkg_pb_notification_proto_depIdxs = []int32{
	0, // 0: notification.NotificationService.SendCommentedNotification:input_type -> notification.CommentedNotification
	1, // 1: notification.NotificationService.SendFollowedNotification:input_type -> notification.FollowedNotification
	2, // 2: notification.NotificationService.SendKafkaNotification:input_type -> notification.KafkaNotification
	3, // 3: notification.NotificationService.SendLikeNotification:input_type -> notification.LikeNotification
	4, // 4: notification.NotificationService.SendCommentedNotification:output_type -> notification.NotificationResponse
	4, // 5: notification.NotificationService.SendFollowedNotification:output_type -> notification.NotificationResponse
	4, // 6: notification.NotificationService.SendKafkaNotification:output_type -> notification.NotificationResponse
	4, // 7: notification.NotificationService.SendLikeNotification:output_type -> notification.NotificationResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_pb_notification_proto_init() }
func file_pkg_pb_notification_proto_init() {
	if File_pkg_pb_notification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentedNotification); i {
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
		file_pkg_pb_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowedNotification); i {
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
		file_pkg_pb_notification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KafkaNotification); i {
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
		file_pkg_pb_notification_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeNotification); i {
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
		file_pkg_pb_notification_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationResponse); i {
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
			RawDescriptor: file_pkg_pb_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_notification_proto_goTypes,
		DependencyIndexes: file_pkg_pb_notification_proto_depIdxs,
		MessageInfos:      file_pkg_pb_notification_proto_msgTypes,
	}.Build()
	File_pkg_pb_notification_proto = out.File
	file_pkg_pb_notification_proto_rawDesc = nil
	file_pkg_pb_notification_proto_goTypes = nil
	file_pkg_pb_notification_proto_depIdxs = nil
}
