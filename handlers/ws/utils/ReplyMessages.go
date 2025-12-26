package utils

import (
	"fmt"
)

func ReplyMessages() {
	for msg := range Broadcast {
		receiverID, senderMsg := msg[0], msg[1]

		// Deliver to the receiver if connected
		if client, ok := ActiveUsers[receiverID]; ok {
			fmt.Println("Sender Message: ", senderMsg, "Receiver ID: ", receiverID)
			_ = client.Conn.WriteMessage(1, []byte(senderMsg))
		} else {
			fmt.Printf("User %s not connected\n", receiverID)
		}
	}
}
