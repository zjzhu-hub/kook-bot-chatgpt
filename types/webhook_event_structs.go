package types

import "kook-bot-chatgpt/constants"

// 这里主要定义事件需要对外用到的一些结构

// 发送频道信息
type SendMessage struct {
    Type         constants.Type    	`json:"type,omitempty"`
    TargetID     string 			`json:"target_id,omitempty"`
    Content      string 			`json:"content,omitempty"`
    Quote        string 			`json:"quote,omitempty"`
    Nonce        string 			`json:"nonce,omitempty"`
    TempTargetID string 			`json:"temp_target_id,omitempty"`
    MsgID        string             `json:"msg_id,omitempty"`
}


// 卡片信息结构
type Card struct {
    Type    string              `json:"type"`
    Theme   constants.Theme     `json:"theme"`
    Size    constants.Size      `json:"size"`
    Modules []Module            `json:"modules"`
    Color   string              `json:"color"`
}

type Module struct {
    Type string     `json:"type"`
    Text ModuleText `json:"text"`
}

type ModuleText struct {
    Type    string `json:"type"`
    Content string `json:"content"`
}
