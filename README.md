<h1 align="center" >Go Web</h1>

<div align="center">
由gin + jwt + casbin + gorm技术栈实现Golang版的RBAC权限管理脚手架, 其主要目的是使Golang初学者通过实践进一步掌握Golang相关的开发技能
</div>
<p align="center">
<img src="https://img.shields.io/badge/Go-v1.16-blue" alt="Go version"/>
<img src="https://img.shields.io/badge/Gin-v1.7.2-brightgreen" alt="Gin version"/>
<img src="https://img.shields.io/badge/Gorm-v1.21.11-brightgreen" alt="Gorm version"/>
<img src="https://img.shields.io/github/license/vagaryer/go-web" alt="License"/>
</p>

## 初衷
在学习golang一段时间后，深感迫切需要动手开发一个项目来进一步掌握和巩固知识，由于之前有过开发web后端的经历，所以决定开发go-web

## 特性
- `RESTful API` 设计风格
- `MySQL` 数据库存储
- `gin` golang web 的微框架
- `gorm` 数据库的ORM管理框架
- `gin-jwt` gin封装的jwt中间件，用户认证
- `casbin` 轻量级开源访问控制框架，RBAC
- `go-redis` redis客户端开发工具
- `viper` 轻便的golang配置管理工具
- `zap` 高性能日志库，提供多种级别的日志打印
- `lumberjack` 日志文件切割归档工具

## 项目结构
```
├── cmd
│    └── admin # admin项目主程序入口
├── configs # 配置目录
├── internal # 内部目录，不对外公开
│    ├── admin # admin项目目录
│    │     ├── api # api目录
│    │     │    └── v1 # v1版本接口目录(类似于Java中的controller), 如果有新版本可以继续添加v2/v3
│    │     ├── router # 路由目录
│    │     ├── service # 业务逻辑目录
│    │     │    └── v1 # v1版本业务目录, 如果有新版本可以继续添加v2/v3
│    │     ├── store # 数据操作目录
│    │     ├── data.go # 初始化数据
│    │     └── router.go # 定义路由规则
│    ├── pkg # 内部公共模块目录
│    │     ├── cache # redis操作目录
│    │     ├── global # 全局公用模型目录
│    │     ├── initialize # 工具初始化目录
│    │     ├── middleware # 中间件目录
│    │     ├── model # 传输模型目录
│    │     ├── response # 响应模型目录
│    │     └── util # 工具包目录
├── pkg # 外部公共模块目录
│   └── model # 存储层模型定义目录
├── .gitignore # git忽略
├── go.mod # go依赖列表
├── go.sum # go依赖下载历史
├── LICENSE # 开源证书
├── README.md # 说明文档
```

## 快速开始
```
# 开始前请使用go mod，它可以为你减少很多麻烦

# 下载项目
git clone https://github.com/vagaryer/go-web.git

# 进入admin主程序入口
cd cmd/admin

# 运行程序
# 使用开发环境默认配置运行
go run main.go 
# 使用生产环境指定配置运行
go run main.go -web_config=xxxx -web_mode=prod

```
> 启动成功后，携带参数username:admin password:123456，发送POST请求到 `http://127.0.0.1:9999/api/v1/base/login`获取token

## 感想
花足够多的时间，做足够多的练习才是学习一门技艺的康庄大道。  
学习过程中一直深受[如何快速高效率地学习Go语言](https://www.cnblogs.com/code-craftsman/p/12515802.html)文章影响，在此与大家分享。

## 特别鸣谢
> 本项目开发过程学习了以下大神的思路和代码风格，感谢大神！！！

<br/>
[gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin): Gin-vue-admin is a full-stack (frontend and backend separation) framework designed for management system.
<br/>
[gin-web](https://github.com/piupuer/gin-web)
<br/>
[iam](https://github.com/marmotedu/iam)


## MIT License

    Copyright (c) 2021 vagaryer