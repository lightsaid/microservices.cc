## 练习
由 01-microservices 基础初窥 rpc，功能却很简陋，为了加深记忆，于是乎动手实现一个相对 01 复杂一点的。
- 项目假设
    -  假如有个web服务器，其中有三个请求
        - api/user/register   // 用户注册
        - api/user/login 		// 用户登陆
        - api/product/list	// 产品列表
        - api/product/one // 一个产品
    - 服务端有两个微服务其中一个用户微服务，另一个是商品微服务

### run 一个
- 启动服务端：`go run ./service/*`
- 启动客户端：`go run ./client/*`

访问上面4个接口，例如 `http://localhost:3000/api/user/login`