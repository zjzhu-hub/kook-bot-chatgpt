package context

import (
	"kook-bot-chatgpt/config"
	"kook-bot-chatgpt/types"
	"sync"
)

// 存储用户上下文信息
type Context struct {
	UserID string
	Chat   types.Chat
}

var globalContextStore = struct {
	store map[string]*Context
	sync.RWMutex
	maxContexts int // 用于控制最大保留上下文信息的数量
}{
	store:       make(map[string]*Context),
	maxContexts: config.GlobalConfig.Storage.MaxContexts, // 默认最大保留10条上下文信息
}

func StoreContext(userID string, context types.ChatMessage) {
	globalContextStore.Lock()
	defer globalContextStore.Unlock()

	if _, ok := globalContextStore.store[userID]; !ok {
		globalContextStore.store[userID] = &Context{UserID: userID}
	}

	if globalContextStore.maxContexts > 0 {
		// 如果设置了最大保留上下文信息的数量，则只保留最近的maxContexts条上下文信息
		if len(globalContextStore.store[userID].Chat.Messages) >= globalContextStore.maxContexts {
			globalContextStore.store[userID].Chat.Messages = globalContextStore.store[userID].Chat.Messages[1:]
		}
	}

	globalContextStore.store[userID].Chat.Messages = append(globalContextStore.store[userID].Chat.Messages, context)
}

func GetContext(userID string) *Context {
	globalContextStore.RLock()
	defer globalContextStore.RUnlock()

	return globalContextStore.store[userID]
}

func ResetContext(userID string) {
	globalContextStore.RLock()
	defer globalContextStore.RUnlock()

	contenxt := globalContextStore.store[userID]
	if contenxt != nil {
		delete(globalContextStore.store, userID)
	}
}
