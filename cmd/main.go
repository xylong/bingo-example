package main

import (
	"bingo-example/bootstrap"
	"bingo-example/bootstrap/routes"
	_ "bingo-example/config"
	_ "bingo-example/docs"
	"bingo-example/http/controllers/api/v1"
	_ "bingo-example/http/controllers/api/v1/auth"
	"bingo-example/http/middlewares"
	"bingo-example/lib/core"
	"bingo-example/pkg/config"
	"flag"
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
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	bootstrap.SetupLogger()

	bingo.Init("conf", "app").
		Inject(bootstrap.NewClient(), core.NewService()).
		Route("swagger", routes.Swagger).
		Mount("v1", v1.Controller...)(middlewares.Cors, middlewares.Log).
		Lunch()
}
