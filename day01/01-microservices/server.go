package main
import (
	"fmt"
	"net/rpc"
	"net"
)

type World struct{
	
}

func (this *World) HelloWorld(name string, resp *string) error{
	*resp = name + " " + "你好！"
	return nil
}

func (this *World) SayHi(name string, resp *string) error{
	*resp = "你好，虎年大吉！"
	return nil
}

func main(){
	// 1. 注册RPC服务，绑定对象方法
	err := rpc.RegisterName("helloms", new(World))
	if err != nil {
		fmt.Println("注册 rpc 服务失败！",err)
		return
	}

	// 2. 设置监听
	listenter, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listenter.Close()
	fmt.Println("开始监听...")

	// 3. 建立链接
	conn, err := listenter.Accept()
	if err != nil {
		fmt.Println("listenter.Accept err:", err)
		return
	}
	fmt.Println("链接成功...")
	defer conn.Close()

	// 4. 绑定服务
	rpc.ServeConn(conn)
}