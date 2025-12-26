package utils

import (
	"fmt"
)

func ReplyMessages() {
	for msg := range Broadcast {
		receiverID, senderMsg := msg[0], msg[1]

		// Deliver to the receiver if connected
		if client, ok := ActiveUsers[receiverID]; ok {
			_ = client.Conn.WriteMessage(1, []byte(senderMsg))
		}

		fmt.Println("Sender Message: ", senderMsg, "Receiver ID: ", receiverID)
	}
}
