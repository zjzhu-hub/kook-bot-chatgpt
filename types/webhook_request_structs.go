package types

import "kook-bot-chatgpt/constants"

// 这里主要定义kook平台发送的一些结构

// 接受信息结构
type ReceivedMessage struct {
	S int  `json:"s"` // 信令类型
	D Data `json:"d"` // 数据
}

type Data struct {
	ChannelType constants.ChannelType `json:"channel_type"`    // 消息通道类型
	Type        constants.Type        `json:"type"`            // 消息类型
	TargetID    string                `json:"target_id"`       // 来源ID 一般是频道id
	AuthorID    string                `json:"author_id"`       // 发送者 id, 1 代表系统
	Content     string                `json:"content"`         // 消息内容, 文件，图片，视频时，content 为 url
	MessageID   string                `json:"msg_id"`          // 消息的 id
	MessageTime int64                 `json:"msg_timestamp"`   // 消息发送时间的毫秒时间戳
	Nonce       string                `json:"nonce"`           // 随机串，与用户消息发送 api 中传的 nonce 保持一致
	Challenge   string                `json:"challenge"`       // 客户端需要原样返回
	VerifyToken string                `json:"verify_token"`    // 机器人的 verify token（不等同于 token）
	Extra       interface{}           `json:"extra,omitempty"` // 不同的消息类型，结构不一致
}
