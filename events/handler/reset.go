package handler

import (
	"kook-bot-chatgpt/context"
	"kook-bot-chatgpt/types"
	"encoding/json"
	"log"
)

type Reset struct{}

// chat指令实现
func (h *Reset) Handler(body map[string]interface{}) {
	log.Println("Command reset Run")

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Println("failed to marshal")
		return
	}

	receivedMessage := types.ReceivedMessage{}
	err = json.Unmarshal(bodyBytes, &receivedMessage)
	if err != nil {
		return
	}

	// 清空用户上下文
	context.ResetContext(receivedMessage.D.AuthorID)
}