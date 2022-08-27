package ws

import (
	"bingo-example/infrastructure/util"
	"github.com/gorilla/websocket"
	"sync"
)

func init() {
	Clients = &clientMap{}
}

var (
	Clients *clientMap
)

// clientMap 客户端集合
type clientMap struct {
	data sync.Map
}

func (m *clientMap) Store(conn *websocket.Conn) {
	client := NewClient(conn)
	m.data.Store(conn.RemoteAddr().String(), client)
}

// Client 客户端
type Client struct {
	conn   *websocket.Conn
	reader chan *Message       // 读消息
	writer chan *util.Response // 写消息
	closer chan struct{}       // 关闭
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		conn:   conn,
		reader: make(chan *Message, 1),
		writer: make(chan *util.Response, 1),
		closer: make(chan struct{}),
	}
}
