package mines

import (
	"mines/structs"
)

func SocketListener(player *structs.Player) {
	for {
		msg := structs.ClientMsg{}
		err := player.Ws.ReadJSON(&msg)
		if err != nil { // Close connection
			playerLeave(player)
			return
		}

		messageHandler(player, &msg)
	}
}

func playerLeave(player *structs.Player) {
	index := pos(player)
	_ = player.Ws.Close()

	if index == -1 {
		return
	}

	Players[index] = Players[len(Players)-1]
	Players = Players[:len(Players)-1]
}

func pos(player *structs.Player) int {
	for i, v := range Players {
		if v == player {
			return i
		}
	}
	return -1
}
