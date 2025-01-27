// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.9.0
// source: rpc/file.proto

package file_rpc

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

type FileInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
}

func (x *FileInfoRequest) Reset() {
	*x = FileInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfoRequest) ProtoMessage() {}

func (x *FileInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfoRequest.ProtoReflect.Descriptor instead.
func (*FileInfoRequest) Descriptor() ([]byte, []int) {
	return file_rpc_file_proto_rawDescGZIP(), []int{0}
}

func (x *FileInfoRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type FileInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Size int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"` // 单位为MB
	Ext  string `protobuf:"bytes,4,opt,name=ext,proto3" json:"ext,omitempty"`
}

func (x *FileInfoResponse) Reset() {
	*x = FileInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfoResponse) ProtoMessage() {}

func (x *FileInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfoResponse.ProtoReflect.Descriptor instead.
func (*FileInfoResponse) Descriptor() ([]byte, []int) {
	return file_rpc_file_proto_rawDescGZIP(), []int{1}
}

func (x *FileInfoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileInfoResponse) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *FileInfoResponse) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileInfoResponse) GetExt() string {
	if x != nil {
		return x.Ext
	}
	return ""
}

var File_rpc_file_proto protoreflect.FileDescriptor

var file_rpc_file_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x22, 0x2a, 0x0a, 0x0f, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x78, 0x74, 0x32, 0x49, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x41, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x72,
	0x70, 0x63, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_file_proto_rawDescOnce sync.Once
	file_rpc_file_proto_rawDescData = file_rpc_file_proto_rawDesc
)

func file_rpc_file_proto_rawDescGZIP() []byte {
	file_rpc_file_proto_rawDescOnce.Do(func() {
		file_rpc_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_file_proto_rawDescData)
	})
	return file_rpc_file_proto_rawDescData
}

var file_rpc_file_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_file_proto_goTypes = []interface{}{
	(*FileInfoRequest)(nil),  // 0: file_rpc.FileInfoRequest
	(*FileInfoResponse)(nil), // 1: file_rpc.FileInfoResponse
}
var file_rpc_file_proto_depIdxs = []int32{
	0, // 0: file_rpc.File.FileInfo:input_type -> file_rpc.FileInfoRequest
	1, // 1: file_rpc.File.FileInfo:output_type -> file_rpc.FileInfoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_file_proto_init() }
func file_rpc_file_proto_init() {
	if File_rpc_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfoRequest); i {
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
		file_rpc_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfoResponse); i {
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
			RawDescriptor: file_rpc_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_file_proto_goTypes,
		DependencyIndexes: file_rpc_file_proto_depIdxs,
		MessageInfos:      file_rpc_file_proto_msgTypes,
	}.Build()
	File_rpc_file_proto = out.File
	file_rpc_file_proto_rawDesc = nil
	file_rpc_file_proto_goTypes = nil
	file_rpc_file_proto_depIdxs = nil
}
