package cmd

import (
	"bingo-example/bootstrap"
	"bingo-example/bootstrap/routes"
	v1 "bingo-example/http/controllers/api/v1"
	"bingo-example/http/middlewares"
	"bingo-example/lib/core"
	"bingo-example/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/xylong/bingo"
)

// CmdServe represents the available web sub-command.
var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	// 运行服务器
	bingo.Init().
		Inject(bootstrap.NewClient(), core.NewService()).
		Route("swagger", routes.Swagger).
		Mount("v1", v1.Controller...)(middlewares.Cors, middlewares.Log).
		Lunch(config.GetInt("app.port"))
}
