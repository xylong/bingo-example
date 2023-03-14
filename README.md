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

### 命令
- swag init -g cmd/main.go
- swag init --parseDependency --parseInternal -g cmd/main.go