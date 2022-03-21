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
	GetChampionByName(ctx context.Context, in *GetChampionByNameRequest, opts ...grpc.CallOption) (*ChampionsList, error)
	GetCurrentChampion(ctx context.Context, in *GetCurrentChampionRequest, opts ...grpc.CallOption) (*ChampionResponse, error)
	GetChampionByDate(ctx context.Context, in *GetChampionByDateRequest, opts ...grpc.CallOption) (*ChampionResponse, error)
	GetChampionByOrderNumber(ctx context.Context, in *ChampionNumber, opts ...grpc.CallOption) (*ChampionResponse, error)
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

func (c *wwdatabaseClient) GetChampionByName(ctx context.Context, in *GetChampionByNameRequest, opts ...grpc.CallOption) (*ChampionsList, error) {
	out := new(ChampionsList)
	err := c.cc.Invoke(ctx, "/wwdatabase.Wwdatabase/GetChampionByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wwdatabaseClient) GetCurrentChampion(ctx context.Context, in *GetCurrentChampionRequest, opts ...grpc.CallOption) (*ChampionResponse, error) {
	out := new(ChampionResponse)
	err := c.cc.Invoke(ctx, "/wwdatabase.Wwdatabase/GetCurrentChampion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wwdatabaseClient) GetChampionByDate(ctx context.Context, in *GetChampionByDateRequest, opts ...grpc.CallOption) (*ChampionResponse, error) {
	out := new(ChampionResponse)
	err := c.cc.Invoke(ctx, "/wwdatabase.Wwdatabase/GetChampionByDate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wwdatabaseClient) GetChampionByOrderNumber(ctx context.Context, in *ChampionNumber, opts ...grpc.CallOption) (*ChampionResponse, error) {
	out := new(ChampionResponse)
	err := c.cc.Invoke(ctx, "/wwdatabase.Wwdatabase/GetChampionByOrderNumber", in, out, opts...)
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
	GetChampionByName(context.Context, *GetChampionByNameRequest) (*ChampionsList, error)
	GetCurrentChampion(context.Context, *GetCurrentChampionRequest) (*ChampionResponse, error)
	GetChampionByDate(context.Context, *GetChampionByDateRequest) (*ChampionResponse, error)
	GetChampionByOrderNumber(context.Context, *ChampionNumber) (*ChampionResponse, error)
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
func (UnimplementedWwdatabaseServer) GetChampionByName(context.Context, *GetChampionByNameRequest) (*ChampionsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChampionByName not implemented")
}
func (UnimplementedWwdatabaseServer) GetCurrentChampion(context.Context, *GetCurrentChampionRequest) (*ChampionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentChampion not implemented")
}
func (UnimplementedWwdatabaseServer) GetChampionByDate(context.Context, *GetChampionByDateRequest) (*ChampionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChampionByDate not implemented")
}
func (UnimplementedWwdatabaseServer) GetChampionByOrderNumber(context.Context, *ChampionNumber) (*ChampionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChampionByOrderNumber not implemented")
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

func _Wwdatabase_GetChampionByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChampionByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WwdatabaseServer).GetChampionByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wwdatabase.Wwdatabase/GetChampionByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WwdatabaseServer).GetChampionByName(ctx, req.(*GetChampionByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wwdatabase_GetCurrentChampion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentChampionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WwdatabaseServer).GetCurrentChampion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wwdatabase.Wwdatabase/GetCurrentChampion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WwdatabaseServer).GetCurrentChampion(ctx, req.(*GetCurrentChampionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wwdatabase_GetChampionByDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChampionByDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WwdatabaseServer).GetChampionByDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wwdatabase.Wwdatabase/GetChampionByDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WwdatabaseServer).GetChampionByDate(ctx, req.(*GetChampionByDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wwdatabase_GetChampionByOrderNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChampionNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WwdatabaseServer).GetChampionByOrderNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wwdatabase.Wwdatabase/GetChampionByOrderNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WwdatabaseServer).GetChampionByOrderNumber(ctx, req.(*ChampionNumber))
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
			MethodName: "GetChampionByName",
			Handler:    _Wwdatabase_GetChampionByName_Handler,
		},
		{
			MethodName: "GetCurrentChampion",
			Handler:    _Wwdatabase_GetCurrentChampion_Handler,
		},
		{
			MethodName: "GetChampionByDate",
			Handler:    _Wwdatabase_GetChampionByDate_Handler,
		},
		{
			MethodName: "GetChampionByOrderNumber",
			Handler:    _Wwdatabase_GetChampionByOrderNumber_Handler,
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
