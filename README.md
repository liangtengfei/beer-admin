# beer-admin

项目目录结构大概（参考了DDD和简洁架构设计的项目结构）：

```
application
|____api
| |____helloworld
| | |____v1
| | |____errors
|____cmd
| |____helloworld
|____configs
|____internal
| |____conf
| |____data
| |____biz
| |____service
| |____server
|____test
|____pkg
|____go.mod
|____go.sum
|____LICENSE
|____README.md
```

`目录说明详情可以查看：[Go工程化 - Project Layout 最佳实践](https://go-kratos.dev/blog/go-project-layout/)`

## 技术栈

1. `ent` [ent](https://entgo.io/) 是一个简单但功能强大的 Go 实体框架，可以轻松构建和维护具有大型数据模型的应用程序。
2. `Wire` [Wire](https://github.com/google/wire) 是一种代码生成工具，可使用依赖注入自动连接组件。
3.  `Gin` [Gin](https://gin-gonic.com/) 是一个用 Go (Golang) 编写的 Web 框架。 它具有类似 martini 的
  API，性能要好得多，多亏了 [httprouter](https://github.com/julienschmidt/httprouter)，速度提高了 40 倍。
4. `Asynq` [Asynq](https://github.com/hibiken/asynq) Asynq 是一个 Go 库，用于排队任务并与 worker 异步处理它们。它由 Redis
  提供支持，旨在实现可扩展且易于上手。