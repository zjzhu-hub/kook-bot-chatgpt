package handler

import (
	"encoding/json"
	"kook-bot-chatgpt/constants"
	"kook-bot-chatgpt/types"
	"kook-bot-chatgpt/utils"
	"log"
)

type Help struct{}

// help指令实现
func (h *Help) Handler(body map[string]interface{}) {
	log.Println("Command help Run")

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

	// 这里有些没有定义常量，等官网有详细的文档在定义常量避免错误
	cards := []types.Card{
		{Type: "card", Theme: constants.ThemeSecondary, Size: constants.SizeLG,
			Modules: []types.Module{
				{Type: "section", Text: types.ModuleText{Type: "kmarkdown", Content: "(font)/help(font)[success] 帮助信息"}},
				{Type: "section", Text: types.ModuleText{Type: "kmarkdown", Content: "(font)/chat {content}(font)[success] 创建一个新的主题"}},
				{Type: "section", Text: types.ModuleText{Type: "kmarkdown", Content: "(font)/reset(font)[success] 重置会话主题"}},
			},
		},
	}
	content, err := json.Marshal(cards)
	if err != nil {
		log.Println("Parse Cards to JSON String fail")
	}

	client := utils.NewKookClient()
	_, err = client.CreateMessage(constants.MSGCARD, receivedMessage.D.TargetID, string(content))
	if err != nil {
		log.Println("Failed to create message:", err)
		return
	}
}
