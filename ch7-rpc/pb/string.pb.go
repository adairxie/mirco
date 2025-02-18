// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.7.1
// source: pb/string.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type StringRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
	B string `protobuf:"bytes,2,opt,name=B,proto3" json:"B,omitempty"`
}

func (x *StringRequest) Reset() {
	*x = StringRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_string_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringRequest) ProtoMessage() {}

func (x *StringRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_string_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringRequest.ProtoReflect.Descriptor instead.
func (*StringRequest) Descriptor() ([]byte, []int) {
	return file_pb_string_proto_rawDescGZIP(), []int{0}
}

func (x *StringRequest) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *StringRequest) GetB() string {
	if x != nil {
		return x.B
	}
	return ""
}

type StringResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret string `protobuf:"bytes,1,opt,name=Ret,proto3" json:"Ret,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *StringResponse) Reset() {
	*x = StringResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_string_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringResponse) ProtoMessage() {}

func (x *StringResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_string_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringResponse.ProtoReflect.Descriptor instead.
func (*StringResponse) Descriptor() ([]byte, []int) {
	return file_pb_string_proto_rawDescGZIP(), []int{1}
}

func (x *StringResponse) GetRet() string {
	if x != nil {
		return x.Ret
	}
	return ""
}

func (x *StringResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_pb_string_proto protoreflect.FileDescriptor

var file_pb_string_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x2b, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x41, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x01, 0x41, 0x12, 0x0c, 0x0a, 0x01, 0x42, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x01, 0x42, 0x22, 0x34, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x52, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x32, 0x73, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x43, 0x6f, 0x6e,
	0x63, 0x61, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x04,
	0x44, 0x69, 0x66, 0x66, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_string_proto_rawDescOnce sync.Once
	file_pb_string_proto_rawDescData = file_pb_string_proto_rawDesc
)

func file_pb_string_proto_rawDescGZIP() []byte {
	file_pb_string_proto_rawDescOnce.Do(func() {
		file_pb_string_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_string_proto_rawDescData)
	})
	return file_pb_string_proto_rawDescData
}

var file_pb_string_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_string_proto_goTypes = []interface{}{
	(*StringRequest)(nil),  // 0: pb.StringRequest
	(*StringResponse)(nil), // 1: pb.StringResponse
}
var file_pb_string_proto_depIdxs = []int32{
	0, // 0: pb.StringService.Concat:input_type -> pb.StringRequest
	0, // 1: pb.StringService.Diff:input_type -> pb.StringRequest
	1, // 2: pb.StringService.Concat:output_type -> pb.StringResponse
	1, // 3: pb.StringService.Diff:output_type -> pb.StringResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_string_proto_init() }
func file_pb_string_proto_init() {
	if File_pb_string_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_string_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringRequest); i {
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
		file_pb_string_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringResponse); i {
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
			RawDescriptor: file_pb_string_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_string_proto_goTypes,
		DependencyIndexes: file_pb_string_proto_depIdxs,
		MessageInfos:      file_pb_string_proto_msgTypes,
	}.Build()
	File_pb_string_proto = out.File
	file_pb_string_proto_rawDesc = nil
	file_pb_string_proto_goTypes = nil
	file_pb_string_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StringServiceClient is the client API for StringService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StringServiceClient interface {
	Concat(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResponse, error)
	Diff(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResponse, error)
}

type stringServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStringServiceClient(cc grpc.ClientConnInterface) StringServiceClient {
	return &stringServiceClient{cc}
}

func (c *stringServiceClient) Concat(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResponse, error) {
	out := new(StringResponse)
	err := c.cc.Invoke(ctx, "/pb.StringService/Concat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stringServiceClient) Diff(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResponse, error) {
	out := new(StringResponse)
	err := c.cc.Invoke(ctx, "/pb.StringService/Diff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StringServiceServer is the server API for StringService service.
type StringServiceServer interface {
	Concat(context.Context, *StringRequest) (*StringResponse, error)
	Diff(context.Context, *StringRequest) (*StringResponse, error)
}

// UnimplementedStringServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStringServiceServer struct {
}

func (*UnimplementedStringServiceServer) Concat(context.Context, *StringRequest) (*StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Concat not implemented")
}
func (*UnimplementedStringServiceServer) Diff(context.Context, *StringRequest) (*StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Diff not implemented")
}

func RegisterStringServiceServer(s *grpc.Server, srv StringServiceServer) {
	s.RegisterService(&_StringService_serviceDesc, srv)
}

func _StringService_Concat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServiceServer).Concat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StringService/Concat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServiceServer).Concat(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StringService_Diff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServiceServer).Diff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StringService/Diff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServiceServer).Diff(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StringService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StringService",
	HandlerType: (*StringServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Concat",
			Handler:    _StringService_Concat_Handler,
		},
		{
			MethodName: "Diff",
			Handler:    _StringService_Diff_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/string.proto",
}
