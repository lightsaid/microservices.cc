package main

import (
	"net/http"
	"net/rpc"
	"fmt"
	"log"
	"encoding/json"
)

type User struct{
	ID int
	Name string
	Password string
}

type Product struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}


func getConn(ip string) (*rpc.Client, error){
	// 1. 用 rpc 链接服务器 --Dial()
	conn, err := rpc.Dial("tcp", ip)
	if err != nil {
		fmt.Println("Dial error ", err)
		return nil, err
	}
	return conn, nil
	// defer conn.Close()

	// // 2.调用远程函数
	// var reply string // 接受函数返回值 --- 传出参数
	// // err = conn.Call("helloms.HelloWorld", "新年", &reply)
	// err = conn.Call("helloms.SayHi", "新年", &reply)

	// if err != nil {
	// 	fmt.Println("Call error ", err)
	// 	return
	// }
	// fmt.Println(reply)
}

var userConn, prodConn *rpc.Client
func init (){
	var err error
	prodConn, err = getConn("127.0.0.1:8001")
	if err != nil {
		log.Panic("链接微服务失败 prodConn: ", err)
	}

	userConn, err = getConn("127.0.0.1:8002")
	if err != nil {
		log.Panic("链接微服务失败 userConn", err)
	}
}

func register(w http.ResponseWriter, r *http.Request){
	var user User
	u := &User{
		Name: "lisi",
		Password: "123456",
	}
	err := userConn.Call("MS_User.Register", u, &user)
	if err != nil{
		fmt.Println("注册失败:", err)
		w.Write([]byte("注册失败"))
		return
	}
	enc  := json.NewEncoder(w)
	err = enc.Encode(user)
	if err != nil{
		w.Write([]byte("NewEncoder失败"))
	}
}

func login(w http.ResponseWriter, r *http.Request){
	var user User
	u := User{
		Name: "张三",
		Password: "abc123",
	}
	err := userConn.Call("MS_User.Login", u, &user)
	if err != nil{
		fmt.Println("登陆失败:", err)
		w.Write([]byte("登陆失败"))
		return
	}
	enc  := json.NewEncoder(w)
	err = enc.Encode(user)
	if err != nil{
		w.Write([]byte("NewEncoder失败"))
	}
}

func getProducts(w http.ResponseWriter, r *http.Request){
	fmt.Println(11111111)
	var prods []Product
	err := prodConn.Call("MS_Product.GetProducts", "test", &prods)
	if err != nil{
		w.Write([]byte(err.Error()))
	}
	enc := json.NewEncoder(w)
	err = enc.Encode(prods)
	if err != nil{
		w.Write([]byte("NewEncoder失败"))
	}
}

func getProductById(w http.ResponseWriter, r *http.Request){
	var prod Product
	err := prodConn.Call("MS_Product.GetProductById", 1, &prod)
	if err != nil{
		w.Write([]byte(err.Error()))
	}
	enc  := json.NewEncoder(w)
	err = enc.Encode(prod)
	if err != nil{
		w.Write([]byte("NewEncoder失败"))
	}
}