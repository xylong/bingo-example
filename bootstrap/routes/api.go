package routes

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xylong/bingo"
)

// Swagger api文档
func Swagger(group *bingo.Group) {
	group.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
