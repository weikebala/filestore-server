package main

import (
	"fmt"
	"time"

	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"

	"filestore-server/common"

	weiekbalaProto "filestore-server/service/weikebala/proto"
	weikebalaRpc "filestore-server/service/weikebala/rpc"
)

func startRPCService() {
	service := micro.NewService(
		micro.Name("go.micro.service.weikebala"), // 服务名称
		micro.RegisterTTL(time.Second*10),        // TTL指定从上一次心跳间隔起，超过这个时间服务会被服务发现移除
		micro.RegisterInterval(time.Second*5),    // 让服务在指定时间内重新注册，保持TTL获取的注册时间有效
		micro.Flags(common.CustomFlags...),
	)
	service.Init()

	weiekbalaProto.RegisterWaiterServiceHandler(service.Server(), new(weikebalaRpc.Weikebala))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	// rpc 服务
	startRPCService()
}
