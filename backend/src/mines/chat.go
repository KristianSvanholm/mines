package mines

import(
	"mines/structs"
)

func SendChat(username string, data interface{}) {
	chat := data.(map[string]interface{}) // Convert interface to map

	// Prevent empty messages in chat
	if chat["message"] == "" {
		return
	}

	chatData := map[string]interface{}{
		"name":    username,
		"message": chat["message"],
	}

	for _, player := range Players {
		msg := structs.ClientMsg{MsgType: "chat", MsgData: chatData}
		err := player.Ws.WriteJSON(msg)
		if err != nil {
			continue
		}
	}
}

