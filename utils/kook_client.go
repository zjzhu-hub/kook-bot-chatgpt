package utils

import (
	"bytes"
	"kook-bot-chatgpt/config"
	"kook-bot-chatgpt/constants"
	"kook-bot-chatgpt/types"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	contentType    = "application/json"
)

type KookClient struct{}

type Response struct {
    Code    int             `json:"code"`
    Message string          `json:"message"`
    Data    ResponseData    `json:"data"`
}

type ResponseData struct {
    MsgID           string  `json:"msg_id"`
    MsgTimestamp    int64   `json:"msg_timestamp"`
    Nonce           string  `json:"nonce"`
}

// 发送消息返回msg_id
func (kc KookClient) CreateMessage(msgType constants.Type, channelID, content string) (string, error) {
	message := types.SendMessage{
        Type: msgType,
		TargetID: channelID,
		Content: content,
	}

	payload, err := json.Marshal(message); if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, config.GlobalConfig.Kook.BaseURL + "/message/create", bytes.NewBuffer(payload)); if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", config.GlobalConfig.Kook.Authorization)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req); if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", ErrSendMessageFailed
	}

	body, err := ioutil.ReadAll(res.Body); if err != nil {
		return "", err
    }

	response := Response{}
	err = json.Unmarshal(body, &response); if err != nil {
		return "", err
	}

	return response.Data.MsgID, nil
}

// 根据msg_id更新消息
func (kc KookClient) UpdateMessage(content, msgID string) error {
	message := types.SendMessage{
        MsgID: msgID,
		Content: content,
	}

	payload, err := json.Marshal(message); if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, config.GlobalConfig.Kook.BaseURL + "/message/update", bytes.NewBuffer(payload)); if err != nil {
		return err
	}

	req.Header.Set("Authorization", config.GlobalConfig.Kook.Authorization)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req); if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return ErrSendMessageFailed
	}

	return nil
}


func NewKookClient() *KookClient {
    return &KookClient{}
}

type ClientError struct {
    StatusCode int
    Message    string
}


func (ce *ClientError) Error() string {
    return ce.Message
}


func NewClientError(statusCode int, message string) error {
    return &ClientError{
        StatusCode: statusCode,
        Message:    message,
    }
}

var ErrSendMessageFailed = NewClientError(-1, "send message failed")
