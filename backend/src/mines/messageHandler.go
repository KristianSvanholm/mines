package mines

import (
	"mines/structs"
)

func messageHandler(player *structs.Player, msg *structs.ClientMsg) {

	switch msg.MsgType {
	case "leftClick":
		openCell(msg.MsgData, player.Name)
		break
	case "rightClick":
		setFlag(msg.MsgData)
		break
	case "chat":
		SendChat(player.Name, msg.MsgData)
		break
	default:
		return
	}
}
