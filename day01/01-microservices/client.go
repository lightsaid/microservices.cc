package main

import (
	"net/rpc"
	"fmt"
)

func main(){
	// 1. 用 rpc 链接服务器 --Dial()
	conn, err := rpc.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Dial error ", err)
		return
	}
	defer conn.Close()

	// 2.调用远程函数
	var reply string // 接受函数返回值 --- 传出参数
	// err = conn.Call("helloms.HelloWorld", "新年", &reply)
	err = conn.Call("helloms.SayHi", "新年", &reply)

	if err != nil {
		fmt.Println("Call error ", err)
		return
	}
	fmt.Println(reply)
}
