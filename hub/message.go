package hub

import (
	"SoFast/models"
	"sync"
)

var (
	MsgCache   = new(MessageCache)
	msgPipline = make(chan models.Message, 10)
)

type MessageCache struct {
	Messages []models.Message
	sync.RWMutex
}

func pushCacheMsg(msg models.Message) {
	defer MsgCache.Unlock()
	MsgCache.Lock()
	MsgCache.Messages = append(MsgCache.Messages, msg)
}

func GetAllCacheMsg() []models.Message {
	defer MsgCache.RUnlock()
	MsgCache.RLock()
	return MsgCache.Messages
}

func BroadcastMessage(msg models.Message) {
	msgPipline <- msg
	pushCacheMsg(msg)
}
