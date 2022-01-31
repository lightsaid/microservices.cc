package main

import (
	"net/rpc"
	"fmt"
	"net"
)

func registerServer(msName, ip string, object interface{}) {
	// 1. 注册 RPC 服务，绑定对象方法
	err := rpc.RegisterName(msName, object)
	if err != nil {
		fmt.Println("注册 rpc 服务失败！",err)
		return
	}

	// 2. 设置监听
	listener, err := net.Listen("tcp", ip)
	if err != nil {
		fmt.Println("设置监听失败！",err)
		return
	}

	defer listener.Close()
	fmt.Println(msName,"服务监听成功！")

	// 建立链接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Accept 失败！",err)
		return
	}
	defer conn.Close()
	fmt.Println("链接成功！")

	// 4. 绑定服务
	rpc.ServeConn(conn)
}

func main(){
	// 注册两个微服务
	go func() {
		registerServer("MS_Product", "127.0.0.1:8001",new(Product))
	}()
	go func() {
		registerServer("MS_User", "127.0.0.1:8002",new(User))
	}()

	for{

	}
}	