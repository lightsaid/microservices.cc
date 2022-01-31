# micro services

## 微服务的介绍
1. 单体服务，如果有一个模块崩溃，整个系统就会芭比Q
1. 微服务就是将模块拆分为一个小单体服务
1. 思想就是解耦

### 单体式架构服务
-- 过往大家熟悉的服务器。
1. 特性：
    1. 复杂性随着开发越来高，遇到问题解决困难
    1. 技术债务逐渐上升。（人员走动）
    1. 耦合度高，维护成本大
        1. 出现bug，不容易排查
        1. 解决旧bug，会出现新bug
    1. 持续交付时间较长
    1. 技术选型成功高，风险高。
    1. 扩展性较差
        1. 垂直扩展：通过增加单个系统负荷来实现扩展。
        1. 水平扩展：通过增加更多系统成员来扩展。
    1. 运维成本低
### 微服务
1. -- 优点
    1. 指责单一
    1. 轻量级通信
    1. 独立性 （单独开发、测试）
    1. 迭代开发
1. 缺点
    1. 运维成本较高（远远高于单体，比如开发迭代，每一个服务都需要build发布测试）
    1. 分布式复杂度
    1. 接口实现成本高 （微服务进行通信，需要提高更多接口，服务调用服务接口）
    1. 重复性劳动 （比如配置，每一个服务都需要读取配置文件）
    1. 业务分离困难 （按理论微服务越小越好，那么业务就不好分离）

## 微服务的通信传输协议 RPPC
RPC (Remote Procedure Call Protocol)，意思就是远程调用，通俗来讲就是调用远程函数。<br>
理解RPC
    - 像调用本地函数一样，去调用远程函数。
    - 通过rpc协议，传递：函数名、函数参数。达到在本地，调用远端函数，得到结果返回给本地。
1. 为什么微服务会使用 rpc?
    1. 每个微服务都被封装成进程。彼此“独立”。
    1. 进程和进程之间，可以使用不同的语言实现。
    1. 所有需要一套协议规范，不管每个微服务用什么语言编写（go、python、java...）,只要遵循rpc协议，都能相互调用。

### RPC入门使用
回顾：Go语言一般性网络socket通信
1. server 端：
    1. net.Listen() --listener 创建监听器
    1. listener.Accpet() --conn 启动监听，建立连接
    1. conn.Read()
    1. conn.Write()
    1. defer conn.CLose() / listener.Close()
1. client 端：
    1. net.Dial() --conn
    1. conn.Write()
    1. conn.Read()
    1. defer conn.Close()
#### RPC 使用步骤
1. -- 服务端
    1. 注册 rpc 服务对象。给对象绑定方法（1. 定义类 2. 绑定类方法）
        > rpc.RegisterName("服务名"，回调对象)
    1. 创建监听器
        > listener, err := net.Listen()
    1. 建立连接
        > conn, err := listener.Accept()
    1. 将连接绑定 rpc 服务
        > rpc.ServeConn(conn)
1. -- 客户端
    1. 用 rpc 连接服务器 rpc.Dial()
        > conn, err := rpc.Dial()
    1. 调用远程函数
        > conn.Call("服务名.方法名", 传入参数，传出参数)

### RPC相关函数
1. 注册 rpc 服务 （服务端）
    1. func (server *Server) RegisterName(name string, rcvr interface{}) error
        1. name - 服务名， string 类型
        1. rcvr - 对应 rpc 对象。该对象绑定方法要满足如下条件；
            1. 方法必须是导出的 -- 包外可见。首字母大写。
            1. 方法必须有两个参数，都是导出类型、内建类型。
            1. 方法的第二参数必须是指针（传出参数）
            1. 方法只有一个 error 接口类型的返回值，
            1. 如：
            ``` go
                type World struct{}
                func (this *World) HelloWorld(name string, resp *string) error {}
                rpc.RegisterName("服务名，随意起", new(World))
            ```
1. 绑定 rpc 服务
    1. func (server *Server) ServeConn(conn io.ReadWriteClose)
        1. conn - 成功建立好连接的 socket -- conn
1. 调用远程函数 （客户端）
    ``` go
        func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error
        // serviceMethod: “服务名.方法名”
        // args： 传入参数
        // reply：传出参数，定义一个变量，&变量名，完成传参。
    ```
### rpc 基础，代码实现
- 服务端
``` go
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
```

- 客户端
``` go
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
```

