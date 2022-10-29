package v1

import (
	"bingo-example/application/service/ws"
	"github.com/gorilla/websocket"
	"github.com/xylong/bingo"
	"net/http"
)

func init() {
	registerCtrl(NewWsController())
}

type WsController struct{}

func NewWsController() *WsController {
	return &WsController{}
}

func (c *WsController) Name() string {
	return "WsController"
}

func (c *WsController) Route(group *bingo.Group) {
	group.GET("ws", c.upgradeWs)
}

// upgradeWs 升级ws协议
func (c *WsController) upgradeWs(ctx *bingo.Context) {
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		http.NotFound(ctx.Writer, ctx.Request)
		return
	}

	ws.Clients.Store(conn)
}
