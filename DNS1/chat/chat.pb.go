// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: chat.proto

package chat

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje          string   `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
	ConsistenciaList []string `protobuf:"bytes,2,rep,name=consistenciaList,proto3" json:"consistenciaList,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

func (x *Message) GetConsistenciaList() []string {
	if x != nil {
		return x.ConsistenciaList
	}
	return nil
}

type Archivos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Registro    []byte `protobuf:"bytes,1,opt,name=registro,proto3" json:"registro,omitempty"`
	Otroarchivo []byte `protobuf:"bytes,2,opt,name=otroarchivo,proto3" json:"otroarchivo,omitempty"`
	Reloj       []byte `protobuf:"bytes,3,opt,name=reloj,proto3" json:"reloj,omitempty"`
	Nombre      string `protobuf:"bytes,4,opt,name=nombre,proto3" json:"nombre,omitempty"`
}

func (x *Archivos) Reset() {
	*x = Archivos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Archivos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Archivos) ProtoMessage() {}

func (x *Archivos) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Archivos.ProtoReflect.Descriptor instead.
func (*Archivos) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *Archivos) GetRegistro() []byte {
	if x != nil {
		return x.Registro
	}
	return nil
}

func (x *Archivos) GetOtroarchivo() []byte {
	if x != nil {
		return x.Otroarchivo
	}
	return nil
}

func (x *Archivos) GetReloj() []byte {
	if x != nil {
		return x.Reloj
	}
	return nil
}

func (x *Archivos) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68,
	0x61, 0x74, 0x22, 0x4f, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x73, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x76, 0x0a, 0x08, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x6f, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x6f,
	0x74, 0x72, 0x6f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0b, 0x6f, 0x74, 0x72, 0x6f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x6f, 0x12, 0x14, 0x0a,
	0x05, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x72, 0x65,
	0x6c, 0x6f, 0x6a, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x32, 0xa2, 0x01, 0x0a, 0x0b,
	0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x0e, 0x52,
	0x65, 0x63, 0x69, 0x62, 0x69, 0x72, 0x44, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x0d, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a,
	0x0c, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x12, 0x0d, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a,
	0x0e, 0x56, 0x75, 0x65, 0x6c, 0x74, 0x61, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x6f, 0x73, 0x12,
	0x0e, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x6f, 0x73, 0x1a,
	0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chat_proto_goTypes = []interface{}{
	(*Message)(nil),  // 0: chat.Message
	(*Archivos)(nil), // 1: chat.Archivos
}
var file_chat_proto_depIdxs = []int32{
	0, // 0: chat.ChatService.RecibirDeAdmin:input_type -> chat.Message
	0, // 1: chat.ChatService.RecibirDeCliente:input_type -> chat.Message
	0, // 2: chat.ChatService.RecibirDeBroker:input_type -> chat.Message
	0, // 3: chat.ChatService.RecibirDeAdmin:output_type -> chat.Message
	0, // 4: chat.ChatService.RecibirDeCliente:output_type -> chat.Message
	0, // 5: chat.ChatService.RecibirDeBroker:output_type -> chat.Message
	0, // 1: chat.ChatService.Consistencia:input_type -> chat.Message
	1, // 2: chat.ChatService.VueltaArchivos:input_type -> chat.Archivos
	0, // 3: chat.ChatService.RecibirDeAdmin:output_type -> chat.Message
	0, // 4: chat.ChatService.Consistencia:output_type -> chat.Message
	0, // 5: chat.ChatService.VueltaArchivos:output_type -> chat.Message
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Archivos); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	RecibirDeAdmin(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	RecibirDeCliente(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	RecibirDeBroker(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Consistencia(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	VueltaArchivos(ctx context.Context, in *Archivos, opts ...grpc.CallOption) (*Message, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) RecibirDeAdmin(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/RecibirDeAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) RecibirDeCliente(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/RecibirDeCliente", in, out, opts...)
func (c *chatServiceClient) Consistencia(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/Consistencia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) RecibirDeBroker(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/RecibirDeBroker", in, out, opts...)
func (c *chatServiceClient) VueltaArchivos(ctx context.Context, in *Archivos, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/VueltaArchivos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	RecibirDeAdmin(context.Context, *Message) (*Message, error)
	RecibirDeCliente(context.Context, *Message) (*Message, error)
	RecibirDeBroker(context.Context, *Message) (*Message, error)
	Consistencia(context.Context, *Message) (*Message, error)
	VueltaArchivos(context.Context, *Archivos) (*Message, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) RecibirDeAdmin(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecibirDeAdmin not implemented")
}
func (*UnimplementedChatServiceServer) RecibirDeCliente(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecibirDeCliente not implemented")
}
func (*UnimplementedChatServiceServer) RecibirDeBroker(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecibirDeBroker not implemented")
func (*UnimplementedChatServiceServer) Consistencia(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Consistencia not implemented")
}
func (*UnimplementedChatServiceServer) VueltaArchivos(context.Context, *Archivos) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VueltaArchivos not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_RecibirDeAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).RecibirDeAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/RecibirDeAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).RecibirDeAdmin(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}


func _ChatService_RecibirDeCliente_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {

func _ChatService_Consistencia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {

	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).RecibirDeCliente(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/RecibirDeCliente",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).RecibirDeCliente(ctx, req.(*Message))
		return srv.(ChatServiceServer).Consistencia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/Consistencia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Consistencia(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_RecibirDeBroker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
func _ChatService_VueltaArchivos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Archivos)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).RecibirDeBroker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/RecibirDeBroker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).RecibirDeBroker(ctx, req.(*Message))
		return srv.(ChatServiceServer).VueltaArchivos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/VueltaArchivos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).VueltaArchivos(ctx, req.(*Archivos))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecibirDeAdmin",
			Handler:    _ChatService_RecibirDeAdmin_Handler,
		},
		{
			MethodName: "RecibirDeCliente",
			Handler:    _ChatService_RecibirDeCliente_Handler,
		},
		{
			MethodName: "RecibirDeBroker",
			Handler:    _ChatService_RecibirDeBroker_Handler,
			MethodName: "Consistencia",
			Handler:    _ChatService_Consistencia_Handler,
		},
		{
			MethodName: "VueltaArchivos",
			Handler:    _ChatService_VueltaArchivos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
