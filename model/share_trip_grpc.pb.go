// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: model/share_trip.proto

package model

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ShareTrip_GetDestination_FullMethodName = "/model.ShareTrip/GetDestination"
	ShareTrip_StreamLocation_FullMethodName = "/model.ShareTrip/StreamLocation"
)

// ShareTripClient is the client API for ShareTrip service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShareTripClient interface {
	GetDestination(ctx context.Context, opts ...grpc.CallOption) (ShareTrip_GetDestinationClient, error)
	StreamLocation(ctx context.Context, opts ...grpc.CallOption) (ShareTrip_StreamLocationClient, error)
}

type shareTripClient struct {
	cc grpc.ClientConnInterface
}

func NewShareTripClient(cc grpc.ClientConnInterface) ShareTripClient {
	return &shareTripClient{cc}
}

func (c *shareTripClient) GetDestination(ctx context.Context, opts ...grpc.CallOption) (ShareTrip_GetDestinationClient, error) {
	stream, err := c.cc.NewStream(ctx, &ShareTrip_ServiceDesc.Streams[0], ShareTrip_GetDestination_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &shareTripGetDestinationClient{stream}
	return x, nil
}

type ShareTrip_GetDestinationClient interface {
	Send(*RoutePolyline) error
	Recv() (*RoutePolyline, error)
	grpc.ClientStream
}

type shareTripGetDestinationClient struct {
	grpc.ClientStream
}

func (x *shareTripGetDestinationClient) Send(m *RoutePolyline) error {
	return x.ClientStream.SendMsg(m)
}

func (x *shareTripGetDestinationClient) Recv() (*RoutePolyline, error) {
	m := new(RoutePolyline)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *shareTripClient) StreamLocation(ctx context.Context, opts ...grpc.CallOption) (ShareTrip_StreamLocationClient, error) {
	stream, err := c.cc.NewStream(ctx, &ShareTrip_ServiceDesc.Streams[1], ShareTrip_StreamLocation_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &shareTripStreamLocationClient{stream}
	return x, nil
}

type ShareTrip_StreamLocationClient interface {
	Send(*Point) error
	Recv() (*RouteTrip, error)
	grpc.ClientStream
}

type shareTripStreamLocationClient struct {
	grpc.ClientStream
}

func (x *shareTripStreamLocationClient) Send(m *Point) error {
	return x.ClientStream.SendMsg(m)
}

func (x *shareTripStreamLocationClient) Recv() (*RouteTrip, error) {
	m := new(RouteTrip)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ShareTripServer is the server API for ShareTrip service.
// All implementations must embed UnimplementedShareTripServer
// for forward compatibility
type ShareTripServer interface {
	GetDestination(ShareTrip_GetDestinationServer) error
	StreamLocation(ShareTrip_StreamLocationServer) error
	mustEmbedUnimplementedShareTripServer()
}

// UnimplementedShareTripServer must be embedded to have forward compatible implementations.
type UnimplementedShareTripServer struct {
}

func (UnimplementedShareTripServer) GetDestination(ShareTrip_GetDestinationServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDestination not implemented")
}
func (UnimplementedShareTripServer) StreamLocation(ShareTrip_StreamLocationServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamLocation not implemented")
}
func (UnimplementedShareTripServer) mustEmbedUnimplementedShareTripServer() {}

// UnsafeShareTripServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShareTripServer will
// result in compilation errors.
type UnsafeShareTripServer interface {
	mustEmbedUnimplementedShareTripServer()
}

func RegisterShareTripServer(s grpc.ServiceRegistrar, srv ShareTripServer) {
	s.RegisterService(&ShareTrip_ServiceDesc, srv)
}

func _ShareTrip_GetDestination_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ShareTripServer).GetDestination(&shareTripGetDestinationServer{stream})
}

type ShareTrip_GetDestinationServer interface {
	Send(*RoutePolyline) error
	Recv() (*RoutePolyline, error)
	grpc.ServerStream
}

type shareTripGetDestinationServer struct {
	grpc.ServerStream
}

func (x *shareTripGetDestinationServer) Send(m *RoutePolyline) error {
	return x.ServerStream.SendMsg(m)
}

func (x *shareTripGetDestinationServer) Recv() (*RoutePolyline, error) {
	m := new(RoutePolyline)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ShareTrip_StreamLocation_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ShareTripServer).StreamLocation(&shareTripStreamLocationServer{stream})
}

type ShareTrip_StreamLocationServer interface {
	Send(*RouteTrip) error
	Recv() (*Point, error)
	grpc.ServerStream
}

type shareTripStreamLocationServer struct {
	grpc.ServerStream
}

func (x *shareTripStreamLocationServer) Send(m *RouteTrip) error {
	return x.ServerStream.SendMsg(m)
}

func (x *shareTripStreamLocationServer) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ShareTrip_ServiceDesc is the grpc.ServiceDesc for ShareTrip service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShareTrip_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model.ShareTrip",
	HandlerType: (*ShareTripServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDestination",
			Handler:       _ShareTrip_GetDestination_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamLocation",
			Handler:       _ShareTrip_StreamLocation_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "model/share_trip.proto",
}