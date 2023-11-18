package hub

import (
	"SoFast/models"
	"log"
	"time"
)

func RunBroadcast() {
	dur := time.Second * 3
	ticker := time.NewTicker(dur)
	defer ticker.Stop()
	for {
		select {
		case m := <-msgPipline:
			sendAll(m)
			ticker.Reset(dur)
			break
		case <-ticker.C:
			continue
		}
	}
}

func sendAll(msg models.Message) {
	wsUsers.Range(func(key, value any) bool {
		go func() {
			user := value.(models.User)
			if err := user.Conn.WriteJSON(msg); err != nil {
				log.Printf("发送消息给 %s 失败， %s", user.Username, err)
				RemoveUser(key.(string))
			}
		}()
		return true
	})
}
