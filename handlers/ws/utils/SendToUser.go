package utils

import (
	"fmt"
	"strings"
)

func SendToUser(data []byte) {
	parts := strings.SplitN(string(data), ":", 2)
	var receiverID, senderMsg string

	if len(parts) == 2 {
		receiverID = parts[0]
		senderMsg = parts[1]
	} else {
		return // Skip message if the ID is not provided
	}

	fmt.Println("Sender Message: ", senderMsg, "Receiver ID: ", receiverID)

}
