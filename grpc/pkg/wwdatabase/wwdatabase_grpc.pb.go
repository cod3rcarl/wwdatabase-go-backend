// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: pkg/wwdatabase/wwdatabase.proto

package wwdatabase_go_backend

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

// WwdatabaseClient is the client API for Wwdatabase service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WwdatabaseClient interface {
	GetChampions(ctx context.Context, in *GetChampionsRequest, opts ...grpc.CallOption) (*ChampionsList, error)
	// rpc GetChampionByName (GetChampionByNameRequest) returns (ChampionList) {}
	// rpc GetCurrentChampion (GetCurrentChampionRequest) returns (ChampionResponse) {}
	// rpc GetChampionByDate (GetChampionByDateRequest) returns (ChampionResponse) {}
	GetPreviousChampion(ctx context.Context, in *PreviousChampionNumber, opts ...grpc.CallOption) (*ChampionResponse, error)
	AddChampion(ctx context.Context, in *NewChampionData, opts ...grpc.CallOption) (*CreateChampionPayload, error)
	UpdateChampion(ctx context.Context, in *UpdateChampionData, opts ...grpc.CallOption) (*UpdateChampionPayload, error)
	DeleteChampion(ctx context.Context, in *DeleteChampionRequest, opts ...grpc.CallOption) (*DeleteChampionResponse, error)
}

type wwdatabaseClient struct {
	cc grpc.ClientConnInterface
}

func NewWwdatabaseClient(cc grpc.ClientConnInterface) WwdatabaseClient {
	return &wwdatabaseClient{cc}
}

func (c *wwdatabaseClient) GetChampions(ctx context.Context, in *GetChampionsRequest, opts ...grpc.CallOption) (*ChampionsList, error) {
	out := new(ChampionsList)
	err := c.cc.Invoke(ctx, "/wwdatabase.Wwdatabase/GetChampions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wwdatabaseClient) GetPreviousChampion(ctx context.Context, in *PreviousChampionNumber, opts ...grpc.CallOption) (*ChampionResponse, error) {
	out := new(ChampionResponse)
	err := c.cc.Invoke(ctx, "/wwdatabase.Wwdatabase/GetPreviousChampion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wwdatabaseClient) AddChampion(ctx context.Context, in *NewChampionData, opts ...grpc.CallOption) (*CreateChampionPayload, error) {
	out := new(CreateChampionPayload)
	err := c.cc.Invoke(ctx, "/wwdatabase.Wwdatabase/AddChampion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wwdatabaseClient) UpdateChampion(ctx context.Context, in *UpdateChampionData, opts ...grpc.CallOption) (*UpdateChampionPayload, error) {
	out := new(UpdateChampionPayload)
	err := c.cc.Invoke(ctx, "/wwdatabase.Wwdatabase/UpdateChampion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wwdatabaseClient) DeleteChampion(ctx context.Context, in *DeleteChampionRequest, opts ...grpc.CallOption) (*DeleteChampionResponse, error) {
	out := new(DeleteChampionResponse)
	err := c.cc.Invoke(ctx, "/wwdatabase.Wwdatabase/DeleteChampion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WwdatabaseServer is the server API for Wwdatabase service.
// All implementations must embed UnimplementedWwdatabaseServer
// for forward compatibility
type WwdatabaseServer interface {
	GetChampions(context.Context, *GetChampionsRequest) (*ChampionsList, error)
	// rpc GetChampionByName (GetChampionByNameRequest) returns (ChampionList) {}
	// rpc GetCurrentChampion (GetCurrentChampionRequest) returns (ChampionResponse) {}
	// rpc GetChampionByDate (GetChampionByDateRequest) returns (ChampionResponse) {}
	GetPreviousChampion(context.Context, *PreviousChampionNumber) (*ChampionResponse, error)
	AddChampion(context.Context, *NewChampionData) (*CreateChampionPayload, error)
	UpdateChampion(context.Context, *UpdateChampionData) (*UpdateChampionPayload, error)
	DeleteChampion(context.Context, *DeleteChampionRequest) (*DeleteChampionResponse, error)
	mustEmbedUnimplementedWwdatabaseServer()
}

// UnimplementedWwdatabaseServer must be embedded to have forward compatible implementations.
type UnimplementedWwdatabaseServer struct {
}

func (UnimplementedWwdatabaseServer) GetChampions(context.Context, *GetChampionsRequest) (*ChampionsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChampions not implemented")
}
func (UnimplementedWwdatabaseServer) GetPreviousChampion(context.Context, *PreviousChampionNumber) (*ChampionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPreviousChampion not implemented")
}
func (UnimplementedWwdatabaseServer) AddChampion(context.Context, *NewChampionData) (*CreateChampionPayload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddChampion not implemented")
}
func (UnimplementedWwdatabaseServer) UpdateChampion(context.Context, *UpdateChampionData) (*UpdateChampionPayload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChampion not implemented")
}
func (UnimplementedWwdatabaseServer) DeleteChampion(context.Context, *DeleteChampionRequest) (*DeleteChampionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChampion not implemented")
}
func (UnimplementedWwdatabaseServer) mustEmbedUnimplementedWwdatabaseServer() {}

// UnsafeWwdatabaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WwdatabaseServer will
// result in compilation errors.
type UnsafeWwdatabaseServer interface {
	mustEmbedUnimplementedWwdatabaseServer()
}

func RegisterWwdatabaseServer(s grpc.ServiceRegistrar, srv WwdatabaseServer) {
	s.RegisterService(&Wwdatabase_ServiceDesc, srv)
}

func _Wwdatabase_GetChampions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChampionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WwdatabaseServer).GetChampions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wwdatabase.Wwdatabase/GetChampions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WwdatabaseServer).GetChampions(ctx, req.(*GetChampionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wwdatabase_GetPreviousChampion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreviousChampionNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WwdatabaseServer).GetPreviousChampion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wwdatabase.Wwdatabase/GetPreviousChampion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WwdatabaseServer).GetPreviousChampion(ctx, req.(*PreviousChampionNumber))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wwdatabase_AddChampion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewChampionData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WwdatabaseServer).AddChampion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wwdatabase.Wwdatabase/AddChampion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WwdatabaseServer).AddChampion(ctx, req.(*NewChampionData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wwdatabase_UpdateChampion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChampionData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WwdatabaseServer).UpdateChampion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wwdatabase.Wwdatabase/UpdateChampion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WwdatabaseServer).UpdateChampion(ctx, req.(*UpdateChampionData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wwdatabase_DeleteChampion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChampionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WwdatabaseServer).DeleteChampion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wwdatabase.Wwdatabase/DeleteChampion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WwdatabaseServer).DeleteChampion(ctx, req.(*DeleteChampionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Wwdatabase_ServiceDesc is the grpc.ServiceDesc for Wwdatabase service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wwdatabase_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wwdatabase.Wwdatabase",
	HandlerType: (*WwdatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChampions",
			Handler:    _Wwdatabase_GetChampions_Handler,
		},
		{
			MethodName: "GetPreviousChampion",
			Handler:    _Wwdatabase_GetPreviousChampion_Handler,
		},
		{
			MethodName: "AddChampion",
			Handler:    _Wwdatabase_AddChampion_Handler,
		},
		{
			MethodName: "UpdateChampion",
			Handler:    _Wwdatabase_UpdateChampion_Handler,
		},
		{
			MethodName: "DeleteChampion",
			Handler:    _Wwdatabase_DeleteChampion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/wwdatabase/wwdatabase.proto",
}
