package handler

import (
	"kook-bot-chatgpt/constants"
	"kook-bot-chatgpt/context"
	"kook-bot-chatgpt/types"
	"kook-bot-chatgpt/utils"
	"encoding/json"
	"log"
	"strings"
)

type Chat struct{}

// chat指令实现
func (h *Chat) Handler(body map[string]interface{}) {
	log.Println("Command chat Run")

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

	_, commandContent := utils.ParseCommand(receivedMessage.D.Content)
	if commandContent == "" {
		log.Println("failed content IsEmpty")
		return
	}

	// 这里有些没有定义常量，等官网有详细的文档在定义常量避免错误
	cards := []types.Card{
		{Type: "card", Theme: constants.ThemeSecondary, Size: constants.SizeLG,
			Modules: []types.Module{
				{Type: "section", Text: types.ModuleText{Type: "kmarkdown", Content: "让我先走一根思考一下......\n"}},
			},
		},
	}

	cardMessage, err := json.Marshal(cards)
	if err != nil {
		log.Println("Parse Cards to JSON String fail")
	}

	kookClient := utils.NewKookClient()
	var msgID string
	msgID, err = kookClient.CreateMessage(constants.MSGCARD, receivedMessage.D.TargetID, string(cardMessage))
	if err != nil {
		log.Println("Failed to create message:", err)
		return
	}

	chatGPTClient := utils.NewChatGptClient()
	ch := make(chan string)
	go chatGPTClient.CreateChat(ch, commandContent, receivedMessage.D.AuthorID)

	for msg := range ch {
    	// 处理chan接收到的消息
		cards[0].Modules[0].Text.Content += msg
		cardMessage, err = json.Marshal(cards)
		if err != nil {
			log.Println("Error marshalling JSON:", err)
			return
		}
		kookClient.UpdateMessage(string(cardMessage), msgID)
	}
	
	// 由于kook不支持流只能调用更新接口，所有这里要吧默认发送消息的字符串给去掉，免得影响结果
	cards[0].Modules[0].Text.Content = strings.Replace(cards[0].Modules[0].Text.Content, "让我先走一根思考一下......\n", "", 1)
	// 记录响应上下文
	context.StoreContext(receivedMessage.D.AuthorID, types.ChatMessage{ Role: "assistant", Content: cards[0].Modules[0].Text.Content })
}
