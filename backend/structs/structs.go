package structs

import "github.com/gorilla/websocket"

type ClientMsg struct {
	MsgType string      `json:"msgType"`
	MsgData interface{} `json:"msgData,omitempty"`
}

type Player struct {
	Ws   *websocket.Conn
	Name string
}

type Cell struct {
	Revealed  bool `json:"revealed"`
	TrueCount int  `json:"-"`
	Mine      bool `json:"-"`
	Flagged   bool `json:"flagged"`
	Count     int  `json:"count"`
}

type Coords struct {
	X int `json:"x"`
	Y int `json:"y"`
}
