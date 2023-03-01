package structs

import (
	"github.com/gorilla/websocket"
	"sync"
)

type ClientMsg struct {
	MsgType string      `json:"msgType"`
	MsgData interface{} `json:"msgData,omitempty"`
}

type Player struct {
	Ws   *websocket.Conn
	mu   sync.Mutex
	Name string
}

func (p *Player) Send(v interface{}) error {
    p.mu.Lock()
    defer p.mu.Unlock()
    return p.Ws.WriteJSON(v)
}

type Cell struct {
	Revealed  bool `json:"revealed"`
	TrueCount int  `json:"-"`
	Mine      bool `json:"-"`
	Flagged   bool `json:"flagged"`
	Count     int  `json:"count"`
}

type CellUpdate struct {
	Coords Coords `json:"coords"`
	Cell Cell `json:"cell"`
}

type Coords struct {
	X int `json:"x"`
	Y int `json:"y"`
}
