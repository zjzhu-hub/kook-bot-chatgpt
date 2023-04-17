package utils

import (
	"bufio"
	"bytes"
	"kook-bot-chatgpt/config"
	"kook-bot-chatgpt/context"
	"kook-bot-chatgpt/types"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

// 请求
type ChatGptClient struct{}

// 响应

// 定义一个结构体来映射响应体
type ChatCompletion struct {
    ID      string      `json:"id"`
    Object  string      `json:"object"`
    Created int64       `json:"created"`
    Model   string      `json:"model"`
    Choices []Completion `json:"choices"`
}

type Completion struct {
    Delta         Message `json:"delta"`
    Index         int     `json:"index"`
    FinishReason  string  `json:"finish_reason"`
}

type Message struct {
    Content string `json:"content"`
}

func (c ChatGptClient) CreateChat(ch chan<- string, content, userID string) {
    // 关闭通道避免结束不了无法处理下一个指令
    defer close(ch)

    chat := types.Chat{
        Model:  "gpt-3.5-turbo",
        Stream: true,
        Messages: []types.ChatMessage{
            { Role: "user", Content: content },
        },
    }
    if context.GetContext(userID) == nil {
        // 没有上下文创建一个创建一个新的会话
        context.StoreContext(userID, chat.Messages[0])
    } else {
        // 获取之前的上下文
        chatContext := context.GetContext(userID)
        chat.Messages = chatContext.Chat.Messages
        // 最新的消息接上去
        newChatMessage := types.ChatMessage{ Role: "user", Content: content }
        chat.Messages = append(chat.Messages, newChatMessage)
        // 保留最新的上下文
        context.StoreContext(userID, newChatMessage)
    }

    payload, err := json.Marshal(chat)
	if err != nil {
        log.Println("Error struct to json failed", err)
		return
	}

    req, err := http.NewRequest(http.MethodPost, config.GlobalConfig.OpenAI.URL + "/v1/chat/completions", bytes.NewBuffer(payload))
	if err != nil {
        log.Println("Error NewRequest failed URL: " + config.GlobalConfig.OpenAI.URL + "/v1/chat/completions", err)
		return
	}

	req.Header.Set("Authorization", config.GlobalConfig.OpenAI.Authorization)
	req.Header.Set("Content-Type", "application/json")


    client := &http.Client{}
	res, err := client.Do(req)
    if err != nil {
        log.Println("Error send request failed URL: " + config.GlobalConfig.OpenAI.URL + "/v1/chat/completions", err )
	}

    defer res.Body.Close()

    // 读取响应体
    reader := bufio.NewReader(res.Body)
    msg := strings.Builder{}

    for {
        line, err := reader.ReadBytes('\n')
        if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("Error reading response body:", err)
			return
		}
        
        log.Println(string(line))

        // 结束符
        if (bytes.HasPrefix(line, []byte("data: [DONE]"))) {
            break;
        }

        // 检查该行是否包含JSON数据 必须以data: 开头
        if bytes.HasPrefix(line, []byte("data: ")) {
            // 去除前缀data: 
            dataWithoutPrefix  := bytes.TrimPrefix(line, []byte("data: "))
            
            chatCompletion := ChatCompletion{}
			err = json.Unmarshal(dataWithoutPrefix, &chatCompletion)
			if err != nil {
				log.Println("Error unmarshalling JSON:", err)
				continue
			}
            msg.WriteString(chatCompletion.Choices[0].Delta.Content)
            // 10个字符一组
            if len(msg.String()) > 10 { 
                ch <- msg.String()
                msg.Reset()
            }
        }
    }

    // 处理缓冲区的剩余数据
	if len(msg.String()) > 0 {
		ch <- msg.String()
        msg.Reset()
	}
}

func NewChatGptClient() *ChatGptClient {
    return &ChatGptClient{}
}
