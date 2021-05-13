// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cpaper_asservice/service/service.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	schema "github.com/s7techlab/cckit/examples/cpaper_asservice/schema"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("cpaper_asservice/service/service.proto", fileDescriptor_de505efdc9697bce)
}

var fileDescriptor_de505efdc9697bce = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xd1, 0xaa, 0xd3, 0x30,
	0x18, 0xc7, 0x99, 0x47, 0x2b, 0x46, 0xf4, 0xc8, 0xa7, 0x1c, 0x8e, 0x55, 0x10, 0xab, 0xa8, 0x4c,
	0x4d, 0x40, 0x2f, 0x04, 0x11, 0x2f, 0x7a, 0x1c, 0x63, 0xe0, 0x85, 0x78, 0x29, 0x88, 0xa4, 0xe9,
	0xe7, 0x16, 0x6c, 0x9a, 0x92, 0xa4, 0xb2, 0x3a, 0x7b, 0x23, 0x82, 0x0f, 0xe0, 0xa3, 0xf9, 0x0a,
	0x82, 0xaf, 0x21, 0x4d, 0x5b, 0x84, 0x8d, 0xe2, 0x36, 0x38, 0x57, 0x1f, 0xe5, 0xfb, 0xe7, 0xf7,
	0xfb, 0x53, 0xd2, 0x92, 0x7b, 0xa2, 0xe0, 0x05, 0x9a, 0x0f, 0xdc, 0x5a, 0x34, 0x9f, 0xa5, 0x40,
	0xb6, 0x36, 0x69, 0x61, 0xb4, 0xd3, 0x70, 0xbc, 0x9e, 0xa3, 0xdd, 0x0c, 0x6f, 0xce, 0xb5, 0x9e,
	0x67, 0xc8, 0x78, 0x21, 0x19, 0xcf, 0x73, 0xed, 0xb8, 0x93, 0x3a, 0xb7, 0xed, 0xb9, 0xf0, 0x46,
	0xb7, 0xf5, 0x4f, 0x49, 0xf9, 0x91, 0xa1, 0x2a, 0x5c, 0xd5, 0x2d, 0xef, 0x6e, 0xca, 0xc5, 0x02,
	0x15, 0xef, 0x46, 0x9b, 0x7a, 0xf2, 0x27, 0x20, 0xc1, 0xc9, 0x9b, 0x26, 0x08, 0xef, 0xc9, 0xd9,
	0xd7, 0xd2, 0x3a, 0x38, 0xa2, 0x2d, 0x96, 0xf6, 0x58, 0x3a, 0x69, 0xb0, 0xe1, 0x23, 0xba, 0x59,
	0xb3, 0x45, 0x9d, 0x68, 0xa5, 0xd0, 0x08, 0xc9, 0x33, 0x4f, 0x6a, 0x28, 0xd1, 0xe1, 0xb7, 0x5f,
	0xbf, 0x7f, 0x9e, 0xb9, 0x00, 0xe7, 0x59, 0x7b, 0x0a, 0xbe, 0x8f, 0xc8, 0xc1, 0x14, 0x1d, 0x8c,
	0xb7, 0xc5, 0xcc, 0xd2, 0xf0, 0xc1, 0xb6, 0xd9, 0xe8, 0xbe, 0xd7, 0xdd, 0x86, 0x5b, 0x9d, 0x8e,
	0xad, 0xa4, 0xb5, 0x25, 0x9a, 0x9a, 0xad, 0x5a, 0x42, 0x5e, 0xaa, 0x04, 0x4d, 0x0d, 0x5f, 0xc9,
	0xe1, 0x14, 0x5d, 0x5c, 0x4d, 0x96, 0x0e, 0x4d, 0xce, 0xb3, 0x59, 0x0a, 0x77, 0x06, 0x2d, 0xff,
	0x42, 0x3b, 0x54, 0x09, 0x7d, 0x95, 0x6b, 0x00, 0x7d, 0x15, 0x5c, 0x3a, 0x99, 0xb2, 0x95, 0x4c,
	0x6b, 0xf8, 0x42, 0xce, 0xcd, 0x9a, 0x5e, 0xf0, 0x78, 0x10, 0xe7, 0xf7, 0x6b, 0xcc, 0x1d, 0xec,
	0xc7, 0xde, 0x0e, 0xd1, 0xa5, 0xde, 0xee, 0xdf, 0xc3, 0xf3, 0xd1, 0x18, 0x1c, 0x39, 0x88, 0xcb,
	0x0a, 0x1e, 0x0e, 0xa2, 0xe2, 0xb2, 0xda, 0xdf, 0x7b, 0xe4, 0xbd, 0x57, 0xa2, 0x8b, 0xbd, 0x37,
	0x29, 0xab, 0xc6, 0x5a, 0x93, 0xe0, 0x2d, 0xa6, 0x88, 0x0a, 0xe8, 0x20, 0xab, 0x0d, 0xec, 0xef,
	0xbe, 0xee, 0xdd, 0x57, 0xa3, 0xcb, 0xbd, 0xdb, 0x78, 0x60, 0xa3, 0xff, 0x31, 0x22, 0xc1, 0x2b,
	0xcc, 0xd0, 0xe1, 0xe9, 0x5e, 0xbc, 0xf1, 0xff, 0x2e, 0x5e, 0xfc, 0xf2, 0xdd, 0x8b, 0xb9, 0x74,
	0x8b, 0x32, 0xa1, 0x42, 0x2b, 0x66, 0x9f, 0x39, 0x14, 0x8b, 0x8c, 0x27, 0x4c, 0x88, 0x4f, 0xd2,
	0x31, 0x5c, 0x72, 0x55, 0x64, 0x68, 0xd9, 0xd0, 0x2f, 0x23, 0x09, 0xfc, 0xe7, 0xf8, 0xf4, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xee, 0x94, 0x7e, 0xe4, 0x55, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CPaperClient is the client API for CPaper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CPaperClient interface {
	// List method returns all registered commercial papers
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*schema.CommercialPaperList, error)
	// Get method returns commercial paper data by id
	Get(ctx context.Context, in *schema.CommercialPaperId, opts ...grpc.CallOption) (*schema.CommercialPaper, error)
	// GetByExternalId
	GetByExternalId(ctx context.Context, in *schema.ExternalId, opts ...grpc.CallOption) (*schema.CommercialPaper, error)
	// Issue commercial paper
	Issue(ctx context.Context, in *schema.IssueCommercialPaper, opts ...grpc.CallOption) (*schema.CommercialPaper, error)
	// Buy commercial paper
	Buy(ctx context.Context, in *schema.BuyCommercialPaper, opts ...grpc.CallOption) (*schema.CommercialPaper, error)
	// Redeem commercial paper
	Redeem(ctx context.Context, in *schema.RedeemCommercialPaper, opts ...grpc.CallOption) (*schema.CommercialPaper, error)
	// Delete commercial paper
	Delete(ctx context.Context, in *schema.CommercialPaperId, opts ...grpc.CallOption) (*schema.CommercialPaper, error)
}

type cPaperClient struct {
	cc grpc.ClientConnInterface
}

func NewCPaperClient(cc grpc.ClientConnInterface) CPaperClient {
	return &cPaperClient{cc}
}

func (c *cPaperClient) List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*schema.CommercialPaperList, error) {
	out := new(schema.CommercialPaperList)
	err := c.cc.Invoke(ctx, "/cpaper_asservice.service.CPaper/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPaperClient) Get(ctx context.Context, in *schema.CommercialPaperId, opts ...grpc.CallOption) (*schema.CommercialPaper, error) {
	out := new(schema.CommercialPaper)
	err := c.cc.Invoke(ctx, "/cpaper_asservice.service.CPaper/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPaperClient) GetByExternalId(ctx context.Context, in *schema.ExternalId, opts ...grpc.CallOption) (*schema.CommercialPaper, error) {
	out := new(schema.CommercialPaper)
	err := c.cc.Invoke(ctx, "/cpaper_asservice.service.CPaper/GetByExternalId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPaperClient) Issue(ctx context.Context, in *schema.IssueCommercialPaper, opts ...grpc.CallOption) (*schema.CommercialPaper, error) {
	out := new(schema.CommercialPaper)
	err := c.cc.Invoke(ctx, "/cpaper_asservice.service.CPaper/Issue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPaperClient) Buy(ctx context.Context, in *schema.BuyCommercialPaper, opts ...grpc.CallOption) (*schema.CommercialPaper, error) {
	out := new(schema.CommercialPaper)
	err := c.cc.Invoke(ctx, "/cpaper_asservice.service.CPaper/Buy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPaperClient) Redeem(ctx context.Context, in *schema.RedeemCommercialPaper, opts ...grpc.CallOption) (*schema.CommercialPaper, error) {
	out := new(schema.CommercialPaper)
	err := c.cc.Invoke(ctx, "/cpaper_asservice.service.CPaper/Redeem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPaperClient) Delete(ctx context.Context, in *schema.CommercialPaperId, opts ...grpc.CallOption) (*schema.CommercialPaper, error) {
	out := new(schema.CommercialPaper)
	err := c.cc.Invoke(ctx, "/cpaper_asservice.service.CPaper/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CPaperServer is the server API for CPaper service.
type CPaperServer interface {
	// List method returns all registered commercial papers
	List(context.Context, *empty.Empty) (*schema.CommercialPaperList, error)
	// Get method returns commercial paper data by id
	Get(context.Context, *schema.CommercialPaperId) (*schema.CommercialPaper, error)
	// GetByExternalId
	GetByExternalId(context.Context, *schema.ExternalId) (*schema.CommercialPaper, error)
	// Issue commercial paper
	Issue(context.Context, *schema.IssueCommercialPaper) (*schema.CommercialPaper, error)
	// Buy commercial paper
	Buy(context.Context, *schema.BuyCommercialPaper) (*schema.CommercialPaper, error)
	// Redeem commercial paper
	Redeem(context.Context, *schema.RedeemCommercialPaper) (*schema.CommercialPaper, error)
	// Delete commercial paper
	Delete(context.Context, *schema.CommercialPaperId) (*schema.CommercialPaper, error)
}

// UnimplementedCPaperServer can be embedded to have forward compatible implementations.
type UnimplementedCPaperServer struct {
}

func (*UnimplementedCPaperServer) List(ctx context.Context, req *empty.Empty) (*schema.CommercialPaperList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedCPaperServer) Get(ctx context.Context, req *schema.CommercialPaperId) (*schema.CommercialPaper, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCPaperServer) GetByExternalId(ctx context.Context, req *schema.ExternalId) (*schema.CommercialPaper, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByExternalId not implemented")
}
func (*UnimplementedCPaperServer) Issue(ctx context.Context, req *schema.IssueCommercialPaper) (*schema.CommercialPaper, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Issue not implemented")
}
func (*UnimplementedCPaperServer) Buy(ctx context.Context, req *schema.BuyCommercialPaper) (*schema.CommercialPaper, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Buy not implemented")
}
func (*UnimplementedCPaperServer) Redeem(ctx context.Context, req *schema.RedeemCommercialPaper) (*schema.CommercialPaper, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Redeem not implemented")
}
func (*UnimplementedCPaperServer) Delete(ctx context.Context, req *schema.CommercialPaperId) (*schema.CommercialPaper, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterCPaperServer(s *grpc.Server, srv CPaperServer) {
	s.RegisterService(&_CPaper_serviceDesc, srv)
}

func _CPaper_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPaperServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpaper_asservice.service.CPaper/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPaperServer).List(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPaper_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.CommercialPaperId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPaperServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpaper_asservice.service.CPaper/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPaperServer).Get(ctx, req.(*schema.CommercialPaperId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPaper_GetByExternalId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.ExternalId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPaperServer).GetByExternalId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpaper_asservice.service.CPaper/GetByExternalId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPaperServer).GetByExternalId(ctx, req.(*schema.ExternalId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPaper_Issue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.IssueCommercialPaper)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPaperServer).Issue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpaper_asservice.service.CPaper/Issue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPaperServer).Issue(ctx, req.(*schema.IssueCommercialPaper))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPaper_Buy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.BuyCommercialPaper)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPaperServer).Buy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpaper_asservice.service.CPaper/Buy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPaperServer).Buy(ctx, req.(*schema.BuyCommercialPaper))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPaper_Redeem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.RedeemCommercialPaper)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPaperServer).Redeem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpaper_asservice.service.CPaper/Redeem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPaperServer).Redeem(ctx, req.(*schema.RedeemCommercialPaper))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPaper_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.CommercialPaperId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPaperServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpaper_asservice.service.CPaper/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPaperServer).Delete(ctx, req.(*schema.CommercialPaperId))
	}
	return interceptor(ctx, in, info, handler)
}

var _CPaper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cpaper_asservice.service.CPaper",
	HandlerType: (*CPaperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _CPaper_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CPaper_Get_Handler,
		},
		{
			MethodName: "GetByExternalId",
			Handler:    _CPaper_GetByExternalId_Handler,
		},
		{
			MethodName: "Issue",
			Handler:    _CPaper_Issue_Handler,
		},
		{
			MethodName: "Buy",
			Handler:    _CPaper_Buy_Handler,
		},
		{
			MethodName: "Redeem",
			Handler:    _CPaper_Redeem_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CPaper_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cpaper_asservice/service/service.proto",
}
