package main

import (
	v1 "bingo-example/api/v1"
	"bingo-example/application/middleware"
	"bingo-example/bootstrap/routes"
	_ "bingo-example/docs"
	"bingo-example/lib/core"
	"github.com/xylong/bingo"
)

// @title           Bingo Example API
// @version         1.0
// @description     bingo案例
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth
func main() {
	bingo.Init("conf", "app").
		Inject(core.NewClient(), core.NewService()).
		Route("swagger", routes.Swagger).
		Mount("v1", v1.Controller...)(middleware.Cors).
		Lunch()
}
