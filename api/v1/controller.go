package v1

import "github.com/xylong/bingo"

// Controller v1版本控制器
var Controller = make([]bingo.Controller, 0)

// registerCtrl 注册控制器
func registerCtrl(ctrl bingo.Controller) {
	Controller = append(Controller, ctrl)
}
