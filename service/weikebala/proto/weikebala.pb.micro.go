// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: weikebala.proto

package go_micro_service_waiter

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for WaiterService service

type WaiterService interface {
	// 获取上传入口地址
	DoMD5(ctx context.Context, in *Req, opts ...client.CallOption) (*Res, error)
}

type waiterService struct {
	c    client.Client
	name string
}

func NewWaiterService(name string, c client.Client) WaiterService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.service.waiter"
	}
	return &waiterService{
		c:    c,
		name: name,
	}
}

func (c *waiterService) DoMD5(ctx context.Context, in *Req, opts ...client.CallOption) (*Res, error) {
	req := c.c.NewRequest(c.name, "WaiterService.DoMD5", in)
	out := new(Res)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WaiterService service

type WaiterServiceHandler interface {
	// 获取上传入口地址
	DoMD5(context.Context, *Req, *Res) error
}

func RegisterWaiterServiceHandler(s server.Server, hdlr WaiterServiceHandler, opts ...server.HandlerOption) error {
	type waiterService interface {
		DoMD5(ctx context.Context, in *Req, out *Res) error
	}
	type WaiterService struct {
		waiterService
	}
	h := &waiterServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&WaiterService{h}, opts...))
}

type waiterServiceHandler struct {
	WaiterServiceHandler
}

func (h *waiterServiceHandler) DoMD5(ctx context.Context, in *Req, out *Res) error {
	return h.WaiterServiceHandler.DoMD5(ctx, in, out)
}
