// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	PublishRequest
	PublishResponse
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// NOTE: Protocol buffers used variable-length encoding, so even though uint64
// is arguably much bigger than we'd need, it's simpler to just use that. We
// do, however, use signed ints for disk usage and limit because those fields
// are optional (and should not be reported if they are < 0).
type PublishRequest struct {
	UnixTime      int64  `protobuf:"varint,1,opt,name=unix_time,json=unixTime" json:"unix_time,omitempty"`
	Running       bool   `protobuf:"varint,2,opt,name=running" json:"running,omitempty"`
	MilliCpuUsage uint64 `protobuf:"varint,3,opt,name=milli_cpu_usage,json=milliCpuUsage" json:"milli_cpu_usage,omitempty"`
	MemoryTotalMb uint64 `protobuf:"varint,4,opt,name=memory_total_mb,json=memoryTotalMb" json:"memory_total_mb,omitempty"`
	MemoryRssMb   uint64 `protobuf:"varint,5,opt,name=memory_rss_mb,json=memoryRssMb" json:"memory_rss_mb,omitempty"`
	MemoryLimitMb uint64 `protobuf:"varint,6,opt,name=memory_limit_mb,json=memoryLimitMb" json:"memory_limit_mb,omitempty"`
	DiskUsageMb   int64  `protobuf:"zigzag64,7,opt,name=disk_usage_mb,json=diskUsageMb" json:"disk_usage_mb,omitempty"`
	DiskLimitMb   int64  `protobuf:"zigzag64,8,opt,name=disk_limit_mb,json=diskLimitMb" json:"disk_limit_mb,omitempty"`
	DiskReadKbps  uint64 `protobuf:"varint,9,opt,name=disk_read_kbps,json=diskReadKbps" json:"disk_read_kbps,omitempty"`
	DiskWriteKbps uint64 `protobuf:"varint,10,opt,name=disk_write_kbps,json=diskWriteKbps" json:"disk_write_kbps,omitempty"`
	DiskReadIops  uint64 `protobuf:"varint,11,opt,name=disk_read_iops,json=diskReadIops" json:"disk_read_iops,omitempty"`
	DiskWriteIops uint64 `protobuf:"varint,12,opt,name=disk_write_iops,json=diskWriteIops" json:"disk_write_iops,omitempty"`
}

func (m *PublishRequest) Reset()                    { *m = PublishRequest{} }
func (m *PublishRequest) String() string            { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()               {}
func (*PublishRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PublishRequest) GetUnixTime() int64 {
	if m != nil {
		return m.UnixTime
	}
	return 0
}

func (m *PublishRequest) GetRunning() bool {
	if m != nil {
		return m.Running
	}
	return false
}

func (m *PublishRequest) GetMilliCpuUsage() uint64 {
	if m != nil {
		return m.MilliCpuUsage
	}
	return 0
}

func (m *PublishRequest) GetMemoryTotalMb() uint64 {
	if m != nil {
		return m.MemoryTotalMb
	}
	return 0
}

func (m *PublishRequest) GetMemoryRssMb() uint64 {
	if m != nil {
		return m.MemoryRssMb
	}
	return 0
}

func (m *PublishRequest) GetMemoryLimitMb() uint64 {
	if m != nil {
		return m.MemoryLimitMb
	}
	return 0
}

func (m *PublishRequest) GetDiskUsageMb() int64 {
	if m != nil {
		return m.DiskUsageMb
	}
	return 0
}

func (m *PublishRequest) GetDiskLimitMb() int64 {
	if m != nil {
		return m.DiskLimitMb
	}
	return 0
}

func (m *PublishRequest) GetDiskReadKbps() uint64 {
	if m != nil {
		return m.DiskReadKbps
	}
	return 0
}

func (m *PublishRequest) GetDiskWriteKbps() uint64 {
	if m != nil {
		return m.DiskWriteKbps
	}
	return 0
}

func (m *PublishRequest) GetDiskReadIops() uint64 {
	if m != nil {
		return m.DiskReadIops
	}
	return 0
}

func (m *PublishRequest) GetDiskWriteIops() uint64 {
	if m != nil {
		return m.DiskWriteIops
	}
	return 0
}

type PublishResponse struct {
}

func (m *PublishResponse) Reset()                    { *m = PublishResponse{} }
func (m *PublishResponse) String() string            { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()               {}
func (*PublishResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*PublishRequest)(nil), "PublishRequest")
	proto.RegisterType((*PublishResponse)(nil), "PublishResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Aggregator service

type AggregatorClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
}

type aggregatorClient struct {
	cc *grpc.ClientConn
}

func NewAggregatorClient(cc *grpc.ClientConn) AggregatorClient {
	return &aggregatorClient{cc}
}

func (c *aggregatorClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := grpc.Invoke(ctx, "/Aggregator/Publish", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Aggregator service

type AggregatorServer interface {
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
}

func RegisterAggregatorServer(s *grpc.Server, srv AggregatorServer) {
	s.RegisterService(&_Aggregator_serviceDesc, srv)
}

func _Aggregator_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Aggregator/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Aggregator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Aggregator",
	HandlerType: (*AggregatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Aggregator_Publish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcd, 0x4e, 0xe3, 0x30,
	0x14, 0x46, 0x27, 0xd3, 0x4e, 0x7f, 0xdc, 0x4e, 0x3b, 0xe3, 0x95, 0x05, 0x9b, 0x28, 0x42, 0x28,
	0xab, 0x2c, 0x60, 0xcb, 0x06, 0xb1, 0x42, 0x50, 0x09, 0x59, 0x45, 0x2c, 0xad, 0x98, 0x5a, 0xc1,
	0x6a, 0x12, 0x1b, 0x5f, 0x5b, 0xc0, 0x13, 0xf2, 0x5a, 0xe8, 0x3a, 0x6d, 0x69, 0xcb, 0x32, 0x47,
	0x27, 0x27, 0x9f, 0x74, 0x43, 0xc6, 0xa5, 0xd5, 0x85, 0x75, 0xc6, 0x9b, 0xec, 0xb3, 0x47, 0x66,
	0x0f, 0x41, 0xd6, 0x1a, 0x5e, 0xb8, 0x7a, 0x0d, 0x0a, 0x3c, 0x3d, 0x25, 0xe3, 0xd0, 0xea, 0x77,
	0xe1, 0x75, 0xa3, 0x58, 0x92, 0x26, 0x79, 0x8f, 0x8f, 0x10, 0x2c, 0x75, 0xa3, 0x28, 0x23, 0x43,
	0x17, 0xda, 0x56, 0xb7, 0x15, 0xfb, 0x9d, 0x26, 0xf9, 0x88, 0x6f, 0x1f, 0xe9, 0x39, 0x99, 0x37,
	0xba, 0xae, 0xb5, 0x78, 0xb6, 0x41, 0x04, 0x28, 0x2b, 0xc5, 0x7a, 0x69, 0x92, 0xf7, 0xf9, 0xdf,
	0x88, 0x6f, 0x6c, 0x78, 0x44, 0x18, 0x3d, 0xd5, 0x18, 0xf7, 0x21, 0xbc, 0xf1, 0x65, 0x2d, 0x1a,
	0xc9, 0xfa, 0x1b, 0x2f, 0xe2, 0x25, 0xd2, 0x85, 0xa4, 0x19, 0xd9, 0x00, 0xe1, 0x00, 0xd0, 0xfa,
	0x13, 0xad, 0x49, 0x07, 0x39, 0xc0, 0x42, 0xee, 0xb5, 0x6a, 0xdd, 0x68, 0x8f, 0xd6, 0x60, 0xbf,
	0x75, 0x8f, 0xb4, 0x6b, 0xad, 0x34, 0xac, 0xbb, 0x59, 0x68, 0x0d, 0xd3, 0x24, 0xa7, 0x7c, 0x82,
	0x30, 0xae, 0xda, 0x73, 0x76, 0xa5, 0xd1, 0xb7, 0xb3, 0xed, 0x9c, 0x91, 0x59, 0x74, 0x9c, 0x2a,
	0x57, 0x62, 0x2d, 0x2d, 0xb0, 0x71, 0xfc, 0xdc, 0x14, 0x29, 0x57, 0xe5, 0xea, 0x4e, 0x5a, 0xc0,
	0x55, 0xd1, 0x7a, 0x73, 0xda, 0xab, 0x4e, 0x23, 0xdd, 0x2a, 0xc4, 0x4f, 0x48, 0xa3, 0x77, 0x50,
	0xd3, 0xc6, 0x02, 0x9b, 0x1c, 0xd6, 0x6e, 0xcd, 0x8f, 0x5a, 0xd4, 0xa6, 0x47, 0x35, 0xf4, 0xb2,
	0xff, 0x64, 0xbe, 0x3b, 0x24, 0x58, 0xd3, 0x82, 0xba, 0xb8, 0x22, 0xe4, 0xba, 0xaa, 0x9c, 0xaa,
	0x4a, 0x6f, 0x1c, 0x2d, 0xc8, 0x70, 0x23, 0xd0, 0x79, 0x71, 0x78, 0xf3, 0x93, 0x7f, 0xc5, 0xd1,
	0xbb, 0xd9, 0x2f, 0x39, 0x88, 0x7f, 0xc8, 0xe5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x30,
	0xae, 0xe8, 0x2e, 0x02, 0x00, 0x00,
}
