package ws

import (
	"bingo-example/application/service"
	"bingo-example/infrastructure/util"
	"encoding/json"
	"errors"
	"reflect"
)

const (
	Ping = iota // 心跳包
)

var (
	cmd = map[int]Commander{}
)

func init() {
	cmd[Ping] = (*service.PingService)(nil)
}

// Commander 指令
type Commander interface {
	Execute(string) (*util.Response, error)
}

// Command 命令
type Command struct {
	Type   int
	Action string
	Data   map[string]interface{}
}

// Parse 解析命令
func (c *Command) Parse() (*util.Response, error) {
	if v, ok := cmd[c.Type]; ok {
		obj := reflect.New(reflect.TypeOf(v).Elem()).Interface()
		bytes, err := json.Marshal(c.Data)

		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal(bytes, obj); err != nil {
			return nil, err
		}

		return obj.(Commander).Execute(c.Action)
	}

	return nil, errors.New("invalid command")
}
