// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: shared/v1/pockets.proto

package sharedv1

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

type GetContainersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetContainersRequest) Reset() {
	*x = GetContainersRequest{}
	mi := &file_shared_v1_pockets_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetContainersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContainersRequest) ProtoMessage() {}

func (x *GetContainersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shared_v1_pockets_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContainersRequest.ProtoReflect.Descriptor instead.
func (*GetContainersRequest) Descriptor() ([]byte, []int) {
	return file_shared_v1_pockets_proto_rawDescGZIP(), []int{0}
}

type GetContainersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Containers    []*Container           `protobuf:"bytes,1,rep,name=containers,proto3" json:"containers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetContainersResponse) Reset() {
	*x = GetContainersResponse{}
	mi := &file_shared_v1_pockets_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetContainersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContainersResponse) ProtoMessage() {}

func (x *GetContainersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shared_v1_pockets_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContainersResponse.ProtoReflect.Descriptor instead.
func (*GetContainersResponse) Descriptor() ([]byte, []int) {
	return file_shared_v1_pockets_proto_rawDescGZIP(), []int{1}
}

func (x *GetContainersResponse) GetContainers() []*Container {
	if x != nil {
		return x.Containers
	}
	return nil
}

type Container struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status        string                 `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Container) Reset() {
	*x = Container{}
	mi := &file_shared_v1_pockets_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Container) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container) ProtoMessage() {}

func (x *Container) ProtoReflect() protoreflect.Message {
	mi := &file_shared_v1_pockets_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Container.ProtoReflect.Descriptor instead.
func (*Container) Descriptor() ([]byte, []int) {
	return file_shared_v1_pockets_proto_rawDescGZIP(), []int{2}
}

func (x *Container) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Container) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Container) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_shared_v1_pockets_proto protoreflect.FileDescriptor

var file_shared_v1_pockets_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x76, 0x31, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x22, 0x47, 0x0a, 0x09, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x32, 0x66, 0x0a, 0x0e, 0x50, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x74, 0x66, 0x69,
	0x72, 0x65, 0x2f, 0x70, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shared_v1_pockets_proto_rawDescOnce sync.Once
	file_shared_v1_pockets_proto_rawDescData = file_shared_v1_pockets_proto_rawDesc
)

func file_shared_v1_pockets_proto_rawDescGZIP() []byte {
	file_shared_v1_pockets_proto_rawDescOnce.Do(func() {
		file_shared_v1_pockets_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_v1_pockets_proto_rawDescData)
	})
	return file_shared_v1_pockets_proto_rawDescData
}

var file_shared_v1_pockets_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_shared_v1_pockets_proto_goTypes = []any{
	(*GetContainersRequest)(nil),  // 0: shared.v1.GetContainersRequest
	(*GetContainersResponse)(nil), // 1: shared.v1.GetContainersResponse
	(*Container)(nil),             // 2: shared.v1.Container
}
var file_shared_v1_pockets_proto_depIdxs = []int32{
	2, // 0: shared.v1.GetContainersResponse.containers:type_name -> shared.v1.Container
	0, // 1: shared.v1.PocketsService.GetContainers:input_type -> shared.v1.GetContainersRequest
	1, // 2: shared.v1.PocketsService.GetContainers:output_type -> shared.v1.GetContainersResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_shared_v1_pockets_proto_init() }
func file_shared_v1_pockets_proto_init() {
	if File_shared_v1_pockets_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shared_v1_pockets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shared_v1_pockets_proto_goTypes,
		DependencyIndexes: file_shared_v1_pockets_proto_depIdxs,
		MessageInfos:      file_shared_v1_pockets_proto_msgTypes,
	}.Build()
	File_shared_v1_pockets_proto = out.File
	file_shared_v1_pockets_proto_rawDesc = nil
	file_shared_v1_pockets_proto_goTypes = nil
	file_shared_v1_pockets_proto_depIdxs = nil
}
