package mines

import (
	"mines/structs"
)

func sendToAll(msg *structs.ClientMsg){
	for _, p := range Players {
		err := p.Send(*msg)
		if err != nil {
			continue
		}
	}
}