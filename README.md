bingo脚手架使用demo

### 相关库
- [gin](https://github.com/gin-gonic/gin) http框架
- [zap](https://github.com/uber-go/zap) 日志
- [gorm](gorm.io/gorm) orm
- [squirrel](github.com/Masterminds/squirrel) sql拼装工具
- [elastic](https://github.com/olivere/elastic) es
- [graphql](https://github.com/graphql-go/graphql)
- [mongo](https://go.mongodb.org/mongo-driver) mongodb
- [swagger](https://github.com/swaggo/gin-swagger) 文档
- [govalidator](https://github.com/thedevsaddam/govalidator) 验证器
- [amqp](https://github.com/streadway/amqp) rabbitmq

### 结构
```azure
.
├── application             
│   ├── assembler               // 装配层(领域对象和响应dto之间的类型转换、数据填充)
│   ├── dto
│   ├── service
│   ├── utils
│   └── validation
├── bootstrap                   // 程序模块初始化目录
│   └── routes
├── cmd                         // 命令
├── conf
├── config                      // 配置信息目录
├── constants
│   └── errors
├── docs
├── domain                      //  领域层
│   ├── aggregate               // 聚合层
│   ├── entity                  // 实体、模型
│   └── repository              // 仓储层，定义数据库操作方法
├── gen
├── http                        // http 请求处理逻辑
├── controllers                 // 控制器，存放 API 和视图控制器
│   └── api                     // API 控制器，支持多版本的 API 控制器
│       └── v1                  // v1 版本的 API 控制器
│           ├── auth
│           │   └── signup.go
│           ├── user.go
│           └── ...
│   ├── middlewares             // 中间件
│   └── requests                // 请求验证目录
├── infrastructure              // 基础实施层
│   ├── dao                     // 实现仓储，数据库交互具体代码
│   ├── es
│   └── util
├── initializers
├── interface
├── lib
│   ├── config
│   ├── core
│   └── factory
├── logs
├── pkg                         // 内置辅助包
│   ├── app
│   ├── cache
│   ├── config
│   ├── database
│   ├── es
│   ├── helpers
│   ├── logger
│   └── response
├── storage                     // 内部存储目录
│   └── logs
└── utils
    └── cache
```

### 命令
- swag init -g cmd/main.go
- swag init --parseDependency --parseInternal -g cmd/main.go