
## 目录结构
```tree
app
├── Dockfile
├── README.md
├── cmd
│   ├── admin_server.go                                 # admin 子服务
│   ├── root.go
│   ├── server.go                                       # 默认服务
│   └── tool.go                                         # 相关辅助工具
├── go.mod
├── go.sum
├── internal
│   ├── app-admin-server
│   │   ├── api                                         # API 接口定义
│   │   ├── data
│   │   ├── domain
│   │   └── service
│   ├── app-server
│   │   ├── api                                         # API 接口定义
│   │   │   └── library
│   │   │       └── v1
│   │   │           └── api.go
│   │   ├── data                                        # 领域模型实现
│   │   ├── domain                                      # 领域模型定义
│   │   ├── server                                      # 服务启动
│   │   │   ├── http                                    # http 服务
│   │   │   │   └── server.go
│   │   │   └── rpc                                     # rpc 服务
│   │   │       └── server.go
│   │   └── service                                     # API 实现，内部与领域模型交互
│   │       └── library
│   │           └── v1
│   │               └── service.go
│   ├── app-tool
│   │   └── service                                     # tool 功能实现
│   │       └── service.go
│   └── pkg                                             # 共用组件
│       ├── appcontext                                  # app 上下文
│       │   └── app_context.go
│       ├── cache
│       │   └── redis
│       ├── conf
│       │   └── config.go
│       ├── database
│       │   └── mongo
│       ├── log
│       └── model
│           ├── group.go
│           └── user.go
└── main.go # 采用 cobra 作为启动基座
```


## 使用说明
```bash
 % go run . help
app system

Usage:
  app [command]

Available Commands:
  adminserver app admin-server
  help        Help about any command
  server      app server
  tool        app tool

Flags:
  -h, --help   help for app

Use "app [command] --help" for more information about a command.


% go run . help server
app server

Usage:
  app server [flags]

Flags:
      --config string   config file (default is $HOME/.app.yaml)
  -h, --help            help for server
      --http            run http server
      --rpc             run rpc server
      
      
% go run . help tool
app tool

Usage:
  app tool [command]

Available Commands:
  addGroup    add app group
  addUser     add app user

Flags:
  -h, --help   help for tool

Use "app tool [command] --help" for more information about a command.
```

## 运行 http server
```bash
% go run . server --http
```

## 运行 rpc server
```bash
% go run . server --rpc
```

## 新增后台用户
```bash
% go run . tool addUser
```

## 新增后台分组
```bash
% go run . tool addGroup
```