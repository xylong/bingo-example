package middleware

import "github.com/xylong/bingo/iface"

var (
	Cors    iface.Middleware
	Auth    iface.Middleware
	Request iface.Middleware
)

func init() {
	Cors = newCors()
	Auth = newAuthentication()
	Request = newRequest()
}
