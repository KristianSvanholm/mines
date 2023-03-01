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

	msg := structs.ClientMsg{MsgType: "chat", MsgData: chatData}
	sendToAll(&msg)
}

func SystemMessage(txt string){
	chatData := map[string]interface{}{
		"name":    "System",
		"message": txt,
	}

	msg := structs.ClientMsg{MsgType: "chat", MsgData: chatData}
	sendToAll(&msg)
}

