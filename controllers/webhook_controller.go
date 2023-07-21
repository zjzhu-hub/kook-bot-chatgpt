package controllers

import (
	"encoding/json"
	"io/ioutil"
	"kook-bot-chatgpt/events"
	"kook-bot-chatgpt/queues"
	"kook-bot-chatgpt/utils"
	"log"
	"net/http"
)

type WebhookResponse struct {
	Challenge string `json:"challenge"`
}

// 处理 Kook 平台 Webhook 请求的函数
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling webhook request...")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ErrorLogger(w, "Error reading request body")
		return
	}

	receiveData := make(map[string]interface{})
	err = json.Unmarshal(body, &receiveData)
	if err != nil {
		utils.ErrorLogger(w, "Error parse json")
		return
	}

	log.Println(string(body))

	content, ok := receiveData["d"].(map[string]interface{})["content"].(string)
	if !ok {
		content = ""
	}
	// 指令解析
	command, _ := utils.ParseCommand(content)
	// 放入指令队列
	events.Dispatcher.Push(queues.RequestWithCommand{Command: command, Body: receiveData})

	// 这个字段可能为空需要安全的取出
	challenge, ok := receiveData["d"].(map[string]interface{})["challenge"].(string)
	if !ok {
		challenge = ""
	}
	// 构造回复消息, 并且返回给Kook
	kookResp := WebhookResponse{
		Challenge: challenge,
	}
	respData, err := json.Marshal(kookResp)
	if err != nil {
		log.Printf("failed to marshal webhook response: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respData)
}
