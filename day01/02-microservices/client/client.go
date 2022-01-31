package main
import (
	"net/http"
	"time"
	"fmt"
)

func main(){
	serve := http.Server{
		Addr:           ":3000",
		Handler:        nil,
		ReadTimeout:   time.Second * 30,
		WriteTimeout:   time.Second * 30,
	}
	http.HandleFunc("/api/user/register", register)
	http.HandleFunc("/api/user/login", login)
	http.HandleFunc("/api/product/list", getProducts)
	http.HandleFunc("/api/product/one", getProductById)

	fmt.Println("web 服务启动成功 ：3000")
	serve.ListenAndServe()
}

