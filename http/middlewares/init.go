package middlewares

import "github.com/xylong/bingo/iface"

var (
	Log     iface.Middleware
	Cors    iface.Middleware
	Auth    iface.Middleware
	Request iface.Middleware
)

func init() {
	Log = newLog()
	Cors = newCors()
	Auth = newAuthentication()
	Request = newRequest()
}
