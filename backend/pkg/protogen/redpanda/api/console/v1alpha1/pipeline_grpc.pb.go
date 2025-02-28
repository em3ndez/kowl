// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: redpanda/api/console/v1alpha1/pipeline.proto

package consolev1alpha1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PipelineService_CreatePipeline_FullMethodName                 = "/redpanda.api.console.v1alpha1.PipelineService/CreatePipeline"
	PipelineService_GetPipeline_FullMethodName                    = "/redpanda.api.console.v1alpha1.PipelineService/GetPipeline"
	PipelineService_DeletePipeline_FullMethodName                 = "/redpanda.api.console.v1alpha1.PipelineService/DeletePipeline"
	PipelineService_ListPipelines_FullMethodName                  = "/redpanda.api.console.v1alpha1.PipelineService/ListPipelines"
	PipelineService_UpdatePipeline_FullMethodName                 = "/redpanda.api.console.v1alpha1.PipelineService/UpdatePipeline"
	PipelineService_StopPipeline_FullMethodName                   = "/redpanda.api.console.v1alpha1.PipelineService/StopPipeline"
	PipelineService_StartPipeline_FullMethodName                  = "/redpanda.api.console.v1alpha1.PipelineService/StartPipeline"
	PipelineService_GetPipelineServiceConfigSchema_FullMethodName = "/redpanda.api.console.v1alpha1.PipelineService/GetPipelineServiceConfigSchema"
	PipelineService_GetPipelinesForSecret_FullMethodName          = "/redpanda.api.console.v1alpha1.PipelineService/GetPipelinesForSecret"
	PipelineService_GetPipelinesBySecrets_FullMethodName          = "/redpanda.api.console.v1alpha1.PipelineService/GetPipelinesBySecrets"
)

// PipelineServiceClient is the client API for PipelineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PipelineServiceClient interface {
	CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*CreatePipelineResponse, error)
	GetPipeline(ctx context.Context, in *GetPipelineRequest, opts ...grpc.CallOption) (*GetPipelineResponse, error)
	DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*DeletePipelineResponse, error)
	ListPipelines(ctx context.Context, in *ListPipelinesRequest, opts ...grpc.CallOption) (*ListPipelinesResponse, error)
	UpdatePipeline(ctx context.Context, in *UpdatePipelineRequest, opts ...grpc.CallOption) (*UpdatePipelineResponse, error)
	StopPipeline(ctx context.Context, in *StopPipelineRequest, opts ...grpc.CallOption) (*StopPipelineResponse, error)
	StartPipeline(ctx context.Context, in *StartPipelineRequest, opts ...grpc.CallOption) (*StartPipelineResponse, error)
	GetPipelineServiceConfigSchema(ctx context.Context, in *GetPipelineServiceConfigSchemaRequest, opts ...grpc.CallOption) (*GetPipelineServiceConfigSchemaResponse, error)
	GetPipelinesForSecret(ctx context.Context, in *GetPipelinesForSecretRequest, opts ...grpc.CallOption) (*GetPipelinesForSecretResponse, error)
	GetPipelinesBySecrets(ctx context.Context, in *GetPipelinesBySecretsRequest, opts ...grpc.CallOption) (*GetPipelinesBySecretsResponse, error)
}

type pipelineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPipelineServiceClient(cc grpc.ClientConnInterface) PipelineServiceClient {
	return &pipelineServiceClient{cc}
}

func (c *pipelineServiceClient) CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*CreatePipelineResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePipelineResponse)
	err := c.cc.Invoke(ctx, PipelineService_CreatePipeline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) GetPipeline(ctx context.Context, in *GetPipelineRequest, opts ...grpc.CallOption) (*GetPipelineResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPipelineResponse)
	err := c.cc.Invoke(ctx, PipelineService_GetPipeline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*DeletePipelineResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePipelineResponse)
	err := c.cc.Invoke(ctx, PipelineService_DeletePipeline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) ListPipelines(ctx context.Context, in *ListPipelinesRequest, opts ...grpc.CallOption) (*ListPipelinesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPipelinesResponse)
	err := c.cc.Invoke(ctx, PipelineService_ListPipelines_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) UpdatePipeline(ctx context.Context, in *UpdatePipelineRequest, opts ...grpc.CallOption) (*UpdatePipelineResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePipelineResponse)
	err := c.cc.Invoke(ctx, PipelineService_UpdatePipeline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) StopPipeline(ctx context.Context, in *StopPipelineRequest, opts ...grpc.CallOption) (*StopPipelineResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StopPipelineResponse)
	err := c.cc.Invoke(ctx, PipelineService_StopPipeline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) StartPipeline(ctx context.Context, in *StartPipelineRequest, opts ...grpc.CallOption) (*StartPipelineResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartPipelineResponse)
	err := c.cc.Invoke(ctx, PipelineService_StartPipeline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) GetPipelineServiceConfigSchema(ctx context.Context, in *GetPipelineServiceConfigSchemaRequest, opts ...grpc.CallOption) (*GetPipelineServiceConfigSchemaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPipelineServiceConfigSchemaResponse)
	err := c.cc.Invoke(ctx, PipelineService_GetPipelineServiceConfigSchema_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) GetPipelinesForSecret(ctx context.Context, in *GetPipelinesForSecretRequest, opts ...grpc.CallOption) (*GetPipelinesForSecretResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPipelinesForSecretResponse)
	err := c.cc.Invoke(ctx, PipelineService_GetPipelinesForSecret_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) GetPipelinesBySecrets(ctx context.Context, in *GetPipelinesBySecretsRequest, opts ...grpc.CallOption) (*GetPipelinesBySecretsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPipelinesBySecretsResponse)
	err := c.cc.Invoke(ctx, PipelineService_GetPipelinesBySecrets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PipelineServiceServer is the server API for PipelineService service.
// All implementations must embed UnimplementedPipelineServiceServer
// for forward compatibility.
type PipelineServiceServer interface {
	CreatePipeline(context.Context, *CreatePipelineRequest) (*CreatePipelineResponse, error)
	GetPipeline(context.Context, *GetPipelineRequest) (*GetPipelineResponse, error)
	DeletePipeline(context.Context, *DeletePipelineRequest) (*DeletePipelineResponse, error)
	ListPipelines(context.Context, *ListPipelinesRequest) (*ListPipelinesResponse, error)
	UpdatePipeline(context.Context, *UpdatePipelineRequest) (*UpdatePipelineResponse, error)
	StopPipeline(context.Context, *StopPipelineRequest) (*StopPipelineResponse, error)
	StartPipeline(context.Context, *StartPipelineRequest) (*StartPipelineResponse, error)
	GetPipelineServiceConfigSchema(context.Context, *GetPipelineServiceConfigSchemaRequest) (*GetPipelineServiceConfigSchemaResponse, error)
	GetPipelinesForSecret(context.Context, *GetPipelinesForSecretRequest) (*GetPipelinesForSecretResponse, error)
	GetPipelinesBySecrets(context.Context, *GetPipelinesBySecretsRequest) (*GetPipelinesBySecretsResponse, error)
	mustEmbedUnimplementedPipelineServiceServer()
}

// UnimplementedPipelineServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPipelineServiceServer struct{}

func (UnimplementedPipelineServiceServer) CreatePipeline(context.Context, *CreatePipelineRequest) (*CreatePipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePipeline not implemented")
}
func (UnimplementedPipelineServiceServer) GetPipeline(context.Context, *GetPipelineRequest) (*GetPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPipeline not implemented")
}
func (UnimplementedPipelineServiceServer) DeletePipeline(context.Context, *DeletePipelineRequest) (*DeletePipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePipeline not implemented")
}
func (UnimplementedPipelineServiceServer) ListPipelines(context.Context, *ListPipelinesRequest) (*ListPipelinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPipelines not implemented")
}
func (UnimplementedPipelineServiceServer) UpdatePipeline(context.Context, *UpdatePipelineRequest) (*UpdatePipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePipeline not implemented")
}
func (UnimplementedPipelineServiceServer) StopPipeline(context.Context, *StopPipelineRequest) (*StopPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopPipeline not implemented")
}
func (UnimplementedPipelineServiceServer) StartPipeline(context.Context, *StartPipelineRequest) (*StartPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPipeline not implemented")
}
func (UnimplementedPipelineServiceServer) GetPipelineServiceConfigSchema(context.Context, *GetPipelineServiceConfigSchemaRequest) (*GetPipelineServiceConfigSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPipelineServiceConfigSchema not implemented")
}
func (UnimplementedPipelineServiceServer) GetPipelinesForSecret(context.Context, *GetPipelinesForSecretRequest) (*GetPipelinesForSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPipelinesForSecret not implemented")
}
func (UnimplementedPipelineServiceServer) GetPipelinesBySecrets(context.Context, *GetPipelinesBySecretsRequest) (*GetPipelinesBySecretsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPipelinesBySecrets not implemented")
}
func (UnimplementedPipelineServiceServer) mustEmbedUnimplementedPipelineServiceServer() {}
func (UnimplementedPipelineServiceServer) testEmbeddedByValue()                         {}

// UnsafePipelineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PipelineServiceServer will
// result in compilation errors.
type UnsafePipelineServiceServer interface {
	mustEmbedUnimplementedPipelineServiceServer()
}

func RegisterPipelineServiceServer(s grpc.ServiceRegistrar, srv PipelineServiceServer) {
	// If the following call pancis, it indicates UnimplementedPipelineServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PipelineService_ServiceDesc, srv)
}

func _PipelineService_CreatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).CreatePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineService_CreatePipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).CreatePipeline(ctx, req.(*CreatePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_GetPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).GetPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineService_GetPipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).GetPipeline(ctx, req.(*GetPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_DeletePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).DeletePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineService_DeletePipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).DeletePipeline(ctx, req.(*DeletePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_ListPipelines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPipelinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).ListPipelines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineService_ListPipelines_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).ListPipelines(ctx, req.(*ListPipelinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_UpdatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).UpdatePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineService_UpdatePipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).UpdatePipeline(ctx, req.(*UpdatePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_StopPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).StopPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineService_StopPipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).StopPipeline(ctx, req.(*StopPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_StartPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).StartPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineService_StartPipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).StartPipeline(ctx, req.(*StartPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_GetPipelineServiceConfigSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPipelineServiceConfigSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).GetPipelineServiceConfigSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineService_GetPipelineServiceConfigSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).GetPipelineServiceConfigSchema(ctx, req.(*GetPipelineServiceConfigSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_GetPipelinesForSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPipelinesForSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).GetPipelinesForSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineService_GetPipelinesForSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).GetPipelinesForSecret(ctx, req.(*GetPipelinesForSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_GetPipelinesBySecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPipelinesBySecretsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).GetPipelinesBySecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineService_GetPipelinesBySecrets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).GetPipelinesBySecrets(ctx, req.(*GetPipelinesBySecretsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PipelineService_ServiceDesc is the grpc.ServiceDesc for PipelineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PipelineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "redpanda.api.console.v1alpha1.PipelineService",
	HandlerType: (*PipelineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePipeline",
			Handler:    _PipelineService_CreatePipeline_Handler,
		},
		{
			MethodName: "GetPipeline",
			Handler:    _PipelineService_GetPipeline_Handler,
		},
		{
			MethodName: "DeletePipeline",
			Handler:    _PipelineService_DeletePipeline_Handler,
		},
		{
			MethodName: "ListPipelines",
			Handler:    _PipelineService_ListPipelines_Handler,
		},
		{
			MethodName: "UpdatePipeline",
			Handler:    _PipelineService_UpdatePipeline_Handler,
		},
		{
			MethodName: "StopPipeline",
			Handler:    _PipelineService_StopPipeline_Handler,
		},
		{
			MethodName: "StartPipeline",
			Handler:    _PipelineService_StartPipeline_Handler,
		},
		{
			MethodName: "GetPipelineServiceConfigSchema",
			Handler:    _PipelineService_GetPipelineServiceConfigSchema_Handler,
		},
		{
			MethodName: "GetPipelinesForSecret",
			Handler:    _PipelineService_GetPipelinesForSecret_Handler,
		},
		{
			MethodName: "GetPipelinesBySecrets",
			Handler:    _PipelineService_GetPipelinesBySecrets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "redpanda/api/console/v1alpha1/pipeline.proto",
}
