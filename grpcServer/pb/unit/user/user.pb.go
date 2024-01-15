// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: pb/unit/user/user.proto

package user

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserCredential *UserCredential `protobuf:"bytes,1,opt,name=userCredential,proto3" json:"userCredential,omitempty"`
	Name           string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email          string          `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_unit_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_pb_unit_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_pb_unit_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserCredential() *UserCredential {
	if x != nil {
		return x.UserCredential
	}
	return nil
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UserCredential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *UserCredential) Reset() {
	*x = UserCredential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_unit_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCredential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCredential) ProtoMessage() {}

func (x *UserCredential) ProtoReflect() protoreflect.Message {
	mi := &file_pb_unit_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCredential.ProtoReflect.Descriptor instead.
func (*UserCredential) Descriptor() ([]byte, []int) {
	return file_pb_unit_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserCredential) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

var File_pb_unit_user_user_proto protoreflect.FileDescriptor

var file_pb_unit_user_user_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x62, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x62, 0x2e, 0x75, 0x6e,
	0x69, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x22, 0x76, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x44, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x75, 0x6e, 0x69,
	0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0x22, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x42, 0x42, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x50,
	0x01, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x6e,
	0x69, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_unit_user_user_proto_rawDescOnce sync.Once
	file_pb_unit_user_user_proto_rawDescData = file_pb_unit_user_user_proto_rawDesc
)

func file_pb_unit_user_user_proto_rawDescGZIP() []byte {
	file_pb_unit_user_user_proto_rawDescOnce.Do(func() {
		file_pb_unit_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_unit_user_user_proto_rawDescData)
	})
	return file_pb_unit_user_user_proto_rawDescData
}

var file_pb_unit_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_unit_user_user_proto_goTypes = []interface{}{
	(*User)(nil),           // 0: pb.unit.user.User
	(*UserCredential)(nil), // 1: pb.unit.user.UserCredential
}
var file_pb_unit_user_user_proto_depIdxs = []int32{
	1, // 0: pb.unit.user.User.userCredential:type_name -> pb.unit.user.UserCredential
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_unit_user_user_proto_init() }
func file_pb_unit_user_user_proto_init() {
	if File_pb_unit_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_unit_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_pb_unit_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCredential); i {
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
			RawDescriptor: file_pb_unit_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_unit_user_user_proto_goTypes,
		DependencyIndexes: file_pb_unit_user_user_proto_depIdxs,
		MessageInfos:      file_pb_unit_user_user_proto_msgTypes,
	}.Build()
	File_pb_unit_user_user_proto = out.File
	file_pb_unit_user_user_proto_rawDesc = nil
	file_pb_unit_user_user_proto_goTypes = nil
	file_pb_unit_user_user_proto_depIdxs = nil
}
