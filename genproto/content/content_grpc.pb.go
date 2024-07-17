// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: content.proto

package content

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Content_CreateStories_FullMethodName       = "/content.Content/CreateStories"
	Content_UpdateStories_FullMethodName       = "/content.Content/UpdateStories"
	Content_DeleteStories_FullMethodName       = "/content.Content/DeleteStories"
	Content_GetAllStories_FullMethodName       = "/content.Content/GetAllStories"
	Content_GetStory_FullMethodName            = "/content.Content/GetStory"
	Content_CommentStory_FullMethodName        = "/content.Content/CommentStory"
	Content_GetCommentsOfStory_FullMethodName  = "/content.Content/GetCommentsOfStory"
	Content_Like_FullMethodName                = "/content.Content/Like"
	Content_Itineraries_FullMethodName         = "/content.Content/Itineraries"
	Content_UpdateItineraries_FullMethodName   = "/content.Content/UpdateItineraries"
	Content_DeleteItineraries_FullMethodName   = "/content.Content/DeleteItineraries"
	Content_GetItineraries_FullMethodName      = "/content.Content/GetItineraries"
	Content_GetItinerariesById_FullMethodName  = "/content.Content/GetItinerariesById"
	Content_CommentItineraries_FullMethodName  = "/content.Content/CommentItineraries"
	Content_GetDestinations_FullMethodName     = "/content.Content/GetDestinations"
	Content_GetDestinationsById_FullMethodName = "/content.Content/GetDestinationsById"
	Content_SendMessage_FullMethodName         = "/content.Content/SendMessage"
	Content_GetMessages_FullMethodName         = "/content.Content/GetMessages"
	Content_CreateTips_FullMethodName          = "/content.Content/CreateTips"
	Content_GetTips_FullMethodName             = "/content.Content/GetTips"
	Content_GetUserStat_FullMethodName         = "/content.Content/GetUserStat"
	Content_TopDestinations_FullMethodName     = "/content.Content/TopDestinations"
)

// ContentClient is the client API for Content service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentClient interface {
	CreateStories(ctx context.Context, in *CreateStoriesRequest, opts ...grpc.CallOption) (*CreateStoriesResponse, error)
	UpdateStories(ctx context.Context, in *UpdateStoriesReq, opts ...grpc.CallOption) (*UpdateStoriesRes, error)
	DeleteStories(ctx context.Context, in *StoryId, opts ...grpc.CallOption) (*Void, error)
	GetAllStories(ctx context.Context, in *GetAllStoriesReq, opts ...grpc.CallOption) (*GetAllStoriesRes, error)
	GetStory(ctx context.Context, in *StoryId, opts ...grpc.CallOption) (*GetStoryRes, error)
	CommentStory(ctx context.Context, in *CommentStoryReq, opts ...grpc.CallOption) (*CommentStoryRes, error)
	GetCommentsOfStory(ctx context.Context, in *GetCommentsOfStoryReq, opts ...grpc.CallOption) (*GetCommentsOfStoryRes, error)
	Like(ctx context.Context, in *LikeReq, opts ...grpc.CallOption) (*LikeRes, error)
	Itineraries(ctx context.Context, in *ItinerariesReq, opts ...grpc.CallOption) (*ItinerariesRes, error)
	UpdateItineraries(ctx context.Context, in *UpdateItinerariesReq, opts ...grpc.CallOption) (*ItinerariesRes, error)
	DeleteItineraries(ctx context.Context, in *StoryId, opts ...grpc.CallOption) (*Void, error)
	GetItineraries(ctx context.Context, in *GetItinerariesReq, opts ...grpc.CallOption) (*GetItinerariesRes, error)
	GetItinerariesById(ctx context.Context, in *StoryId, opts ...grpc.CallOption) (*GetItinerariesByIdRes, error)
	CommentItineraries(ctx context.Context, in *CommentItinerariesReq, opts ...grpc.CallOption) (*CommentItinerariesRes, error)
	GetDestinations(ctx context.Context, in *GetDestinationsReq, opts ...grpc.CallOption) (*GetDestinationsRes, error)
	GetDestinationsById(ctx context.Context, in *GetDestinationsByIdReq, opts ...grpc.CallOption) (*GetDestinationsByIdRes, error)
	SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageRes, error)
	GetMessages(ctx context.Context, in *GetMessagesReq, opts ...grpc.CallOption) (*GetMessagesRes, error)
	CreateTips(ctx context.Context, in *CreateTipsReq, opts ...grpc.CallOption) (*CreateTipsRes, error)
	GetTips(ctx context.Context, in *GetTipsReq, opts ...grpc.CallOption) (*GetTipsRes, error)
	GetUserStat(ctx context.Context, in *GetUserStatReq, opts ...grpc.CallOption) (*GetUserStatRes, error)
	TopDestinations(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Answer, error)
}

type contentClient struct {
	cc grpc.ClientConnInterface
}

func NewContentClient(cc grpc.ClientConnInterface) ContentClient {
	return &contentClient{cc}
}

func (c *contentClient) CreateStories(ctx context.Context, in *CreateStoriesRequest, opts ...grpc.CallOption) (*CreateStoriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateStoriesResponse)
	err := c.cc.Invoke(ctx, Content_CreateStories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) UpdateStories(ctx context.Context, in *UpdateStoriesReq, opts ...grpc.CallOption) (*UpdateStoriesRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateStoriesRes)
	err := c.cc.Invoke(ctx, Content_UpdateStories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) DeleteStories(ctx context.Context, in *StoryId, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, Content_DeleteStories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetAllStories(ctx context.Context, in *GetAllStoriesReq, opts ...grpc.CallOption) (*GetAllStoriesRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllStoriesRes)
	err := c.cc.Invoke(ctx, Content_GetAllStories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetStory(ctx context.Context, in *StoryId, opts ...grpc.CallOption) (*GetStoryRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStoryRes)
	err := c.cc.Invoke(ctx, Content_GetStory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) CommentStory(ctx context.Context, in *CommentStoryReq, opts ...grpc.CallOption) (*CommentStoryRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommentStoryRes)
	err := c.cc.Invoke(ctx, Content_CommentStory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetCommentsOfStory(ctx context.Context, in *GetCommentsOfStoryReq, opts ...grpc.CallOption) (*GetCommentsOfStoryRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommentsOfStoryRes)
	err := c.cc.Invoke(ctx, Content_GetCommentsOfStory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) Like(ctx context.Context, in *LikeReq, opts ...grpc.CallOption) (*LikeRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LikeRes)
	err := c.cc.Invoke(ctx, Content_Like_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) Itineraries(ctx context.Context, in *ItinerariesReq, opts ...grpc.CallOption) (*ItinerariesRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ItinerariesRes)
	err := c.cc.Invoke(ctx, Content_Itineraries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) UpdateItineraries(ctx context.Context, in *UpdateItinerariesReq, opts ...grpc.CallOption) (*ItinerariesRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ItinerariesRes)
	err := c.cc.Invoke(ctx, Content_UpdateItineraries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) DeleteItineraries(ctx context.Context, in *StoryId, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, Content_DeleteItineraries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetItineraries(ctx context.Context, in *GetItinerariesReq, opts ...grpc.CallOption) (*GetItinerariesRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetItinerariesRes)
	err := c.cc.Invoke(ctx, Content_GetItineraries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetItinerariesById(ctx context.Context, in *StoryId, opts ...grpc.CallOption) (*GetItinerariesByIdRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetItinerariesByIdRes)
	err := c.cc.Invoke(ctx, Content_GetItinerariesById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) CommentItineraries(ctx context.Context, in *CommentItinerariesReq, opts ...grpc.CallOption) (*CommentItinerariesRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommentItinerariesRes)
	err := c.cc.Invoke(ctx, Content_CommentItineraries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetDestinations(ctx context.Context, in *GetDestinationsReq, opts ...grpc.CallOption) (*GetDestinationsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDestinationsRes)
	err := c.cc.Invoke(ctx, Content_GetDestinations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetDestinationsById(ctx context.Context, in *GetDestinationsByIdReq, opts ...grpc.CallOption) (*GetDestinationsByIdRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDestinationsByIdRes)
	err := c.cc.Invoke(ctx, Content_GetDestinationsById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendMessageRes)
	err := c.cc.Invoke(ctx, Content_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetMessages(ctx context.Context, in *GetMessagesReq, opts ...grpc.CallOption) (*GetMessagesRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMessagesRes)
	err := c.cc.Invoke(ctx, Content_GetMessages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) CreateTips(ctx context.Context, in *CreateTipsReq, opts ...grpc.CallOption) (*CreateTipsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTipsRes)
	err := c.cc.Invoke(ctx, Content_CreateTips_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetTips(ctx context.Context, in *GetTipsReq, opts ...grpc.CallOption) (*GetTipsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTipsRes)
	err := c.cc.Invoke(ctx, Content_GetTips_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetUserStat(ctx context.Context, in *GetUserStatReq, opts ...grpc.CallOption) (*GetUserStatRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserStatRes)
	err := c.cc.Invoke(ctx, Content_GetUserStat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) TopDestinations(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Answer, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Answer)
	err := c.cc.Invoke(ctx, Content_TopDestinations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServer is the server API for Content service.
// All implementations must embed UnimplementedContentServer
// for forward compatibility
type ContentServer interface {
	CreateStories(context.Context, *CreateStoriesRequest) (*CreateStoriesResponse, error)
	UpdateStories(context.Context, *UpdateStoriesReq) (*UpdateStoriesRes, error)
	DeleteStories(context.Context, *StoryId) (*Void, error)
	GetAllStories(context.Context, *GetAllStoriesReq) (*GetAllStoriesRes, error)
	GetStory(context.Context, *StoryId) (*GetStoryRes, error)
	CommentStory(context.Context, *CommentStoryReq) (*CommentStoryRes, error)
	GetCommentsOfStory(context.Context, *GetCommentsOfStoryReq) (*GetCommentsOfStoryRes, error)
	Like(context.Context, *LikeReq) (*LikeRes, error)
	Itineraries(context.Context, *ItinerariesReq) (*ItinerariesRes, error)
	UpdateItineraries(context.Context, *UpdateItinerariesReq) (*ItinerariesRes, error)
	DeleteItineraries(context.Context, *StoryId) (*Void, error)
	GetItineraries(context.Context, *GetItinerariesReq) (*GetItinerariesRes, error)
	GetItinerariesById(context.Context, *StoryId) (*GetItinerariesByIdRes, error)
	CommentItineraries(context.Context, *CommentItinerariesReq) (*CommentItinerariesRes, error)
	GetDestinations(context.Context, *GetDestinationsReq) (*GetDestinationsRes, error)
	GetDestinationsById(context.Context, *GetDestinationsByIdReq) (*GetDestinationsByIdRes, error)
	SendMessage(context.Context, *SendMessageReq) (*SendMessageRes, error)
	GetMessages(context.Context, *GetMessagesReq) (*GetMessagesRes, error)
	CreateTips(context.Context, *CreateTipsReq) (*CreateTipsRes, error)
	GetTips(context.Context, *GetTipsReq) (*GetTipsRes, error)
	GetUserStat(context.Context, *GetUserStatReq) (*GetUserStatRes, error)
	TopDestinations(context.Context, *Void) (*Answer, error)
	mustEmbedUnimplementedContentServer()
}

// UnimplementedContentServer must be embedded to have forward compatible implementations.
type UnimplementedContentServer struct {
}

func (UnimplementedContentServer) CreateStories(context.Context, *CreateStoriesRequest) (*CreateStoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStories not implemented")
}
func (UnimplementedContentServer) UpdateStories(context.Context, *UpdateStoriesReq) (*UpdateStoriesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStories not implemented")
}
func (UnimplementedContentServer) DeleteStories(context.Context, *StoryId) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStories not implemented")
}
func (UnimplementedContentServer) GetAllStories(context.Context, *GetAllStoriesReq) (*GetAllStoriesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllStories not implemented")
}
func (UnimplementedContentServer) GetStory(context.Context, *StoryId) (*GetStoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStory not implemented")
}
func (UnimplementedContentServer) CommentStory(context.Context, *CommentStoryReq) (*CommentStoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentStory not implemented")
}
func (UnimplementedContentServer) GetCommentsOfStory(context.Context, *GetCommentsOfStoryReq) (*GetCommentsOfStoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentsOfStory not implemented")
}
func (UnimplementedContentServer) Like(context.Context, *LikeReq) (*LikeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Like not implemented")
}
func (UnimplementedContentServer) Itineraries(context.Context, *ItinerariesReq) (*ItinerariesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Itineraries not implemented")
}
func (UnimplementedContentServer) UpdateItineraries(context.Context, *UpdateItinerariesReq) (*ItinerariesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItineraries not implemented")
}
func (UnimplementedContentServer) DeleteItineraries(context.Context, *StoryId) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItineraries not implemented")
}
func (UnimplementedContentServer) GetItineraries(context.Context, *GetItinerariesReq) (*GetItinerariesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItineraries not implemented")
}
func (UnimplementedContentServer) GetItinerariesById(context.Context, *StoryId) (*GetItinerariesByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItinerariesById not implemented")
}
func (UnimplementedContentServer) CommentItineraries(context.Context, *CommentItinerariesReq) (*CommentItinerariesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentItineraries not implemented")
}
func (UnimplementedContentServer) GetDestinations(context.Context, *GetDestinationsReq) (*GetDestinationsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDestinations not implemented")
}
func (UnimplementedContentServer) GetDestinationsById(context.Context, *GetDestinationsByIdReq) (*GetDestinationsByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDestinationsById not implemented")
}
func (UnimplementedContentServer) SendMessage(context.Context, *SendMessageReq) (*SendMessageRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedContentServer) GetMessages(context.Context, *GetMessagesReq) (*GetMessagesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedContentServer) CreateTips(context.Context, *CreateTipsReq) (*CreateTipsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTips not implemented")
}
func (UnimplementedContentServer) GetTips(context.Context, *GetTipsReq) (*GetTipsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTips not implemented")
}
func (UnimplementedContentServer) GetUserStat(context.Context, *GetUserStatReq) (*GetUserStatRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserStat not implemented")
}
func (UnimplementedContentServer) TopDestinations(context.Context, *Void) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopDestinations not implemented")
}
func (UnimplementedContentServer) mustEmbedUnimplementedContentServer() {}

// UnsafeContentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentServer will
// result in compilation errors.
type UnsafeContentServer interface {
	mustEmbedUnimplementedContentServer()
}

func RegisterContentServer(s grpc.ServiceRegistrar, srv ContentServer) {
	s.RegisterService(&Content_ServiceDesc, srv)
}

func _Content_CreateStories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).CreateStories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_CreateStories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).CreateStories(ctx, req.(*CreateStoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_UpdateStories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStoriesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).UpdateStories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_UpdateStories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).UpdateStories(ctx, req.(*UpdateStoriesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_DeleteStories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoryId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).DeleteStories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_DeleteStories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).DeleteStories(ctx, req.(*StoryId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetAllStories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllStoriesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetAllStories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_GetAllStories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetAllStories(ctx, req.(*GetAllStoriesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoryId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_GetStory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetStory(ctx, req.(*StoryId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_CommentStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentStoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).CommentStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_CommentStory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).CommentStory(ctx, req.(*CommentStoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetCommentsOfStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsOfStoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetCommentsOfStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_GetCommentsOfStory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetCommentsOfStory(ctx, req.(*GetCommentsOfStoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_Like_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).Like(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_Like_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).Like(ctx, req.(*LikeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_Itineraries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItinerariesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).Itineraries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_Itineraries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).Itineraries(ctx, req.(*ItinerariesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_UpdateItineraries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateItinerariesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).UpdateItineraries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_UpdateItineraries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).UpdateItineraries(ctx, req.(*UpdateItinerariesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_DeleteItineraries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoryId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).DeleteItineraries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_DeleteItineraries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).DeleteItineraries(ctx, req.(*StoryId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetItineraries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItinerariesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetItineraries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_GetItineraries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetItineraries(ctx, req.(*GetItinerariesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetItinerariesById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoryId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetItinerariesById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_GetItinerariesById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetItinerariesById(ctx, req.(*StoryId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_CommentItineraries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentItinerariesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).CommentItineraries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_CommentItineraries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).CommentItineraries(ctx, req.(*CommentItinerariesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetDestinations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDestinationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetDestinations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_GetDestinations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetDestinations(ctx, req.(*GetDestinationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetDestinationsById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDestinationsByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetDestinationsById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_GetDestinationsById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetDestinationsById(ctx, req.(*GetDestinationsByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).SendMessage(ctx, req.(*SendMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_GetMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetMessages(ctx, req.(*GetMessagesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_CreateTips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTipsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).CreateTips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_CreateTips_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).CreateTips(ctx, req.(*CreateTipsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetTips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTipsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetTips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_GetTips_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetTips(ctx, req.(*GetTipsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetUserStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserStatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetUserStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_GetUserStat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetUserStat(ctx, req.(*GetUserStatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_TopDestinations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).TopDestinations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Content_TopDestinations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).TopDestinations(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

// Content_ServiceDesc is the grpc.ServiceDesc for Content service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Content_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content.Content",
	HandlerType: (*ContentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStories",
			Handler:    _Content_CreateStories_Handler,
		},
		{
			MethodName: "UpdateStories",
			Handler:    _Content_UpdateStories_Handler,
		},
		{
			MethodName: "DeleteStories",
			Handler:    _Content_DeleteStories_Handler,
		},
		{
			MethodName: "GetAllStories",
			Handler:    _Content_GetAllStories_Handler,
		},
		{
			MethodName: "GetStory",
			Handler:    _Content_GetStory_Handler,
		},
		{
			MethodName: "CommentStory",
			Handler:    _Content_CommentStory_Handler,
		},
		{
			MethodName: "GetCommentsOfStory",
			Handler:    _Content_GetCommentsOfStory_Handler,
		},
		{
			MethodName: "Like",
			Handler:    _Content_Like_Handler,
		},
		{
			MethodName: "Itineraries",
			Handler:    _Content_Itineraries_Handler,
		},
		{
			MethodName: "UpdateItineraries",
			Handler:    _Content_UpdateItineraries_Handler,
		},
		{
			MethodName: "DeleteItineraries",
			Handler:    _Content_DeleteItineraries_Handler,
		},
		{
			MethodName: "GetItineraries",
			Handler:    _Content_GetItineraries_Handler,
		},
		{
			MethodName: "GetItinerariesById",
			Handler:    _Content_GetItinerariesById_Handler,
		},
		{
			MethodName: "CommentItineraries",
			Handler:    _Content_CommentItineraries_Handler,
		},
		{
			MethodName: "GetDestinations",
			Handler:    _Content_GetDestinations_Handler,
		},
		{
			MethodName: "GetDestinationsById",
			Handler:    _Content_GetDestinationsById_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _Content_SendMessage_Handler,
		},
		{
			MethodName: "GetMessages",
			Handler:    _Content_GetMessages_Handler,
		},
		{
			MethodName: "CreateTips",
			Handler:    _Content_CreateTips_Handler,
		},
		{
			MethodName: "GetTips",
			Handler:    _Content_GetTips_Handler,
		},
		{
			MethodName: "GetUserStat",
			Handler:    _Content_GetUserStat_Handler,
		},
		{
			MethodName: "TopDestinations",
			Handler:    _Content_TopDestinations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content.proto",
}
