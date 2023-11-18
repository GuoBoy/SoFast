package hub

import (
	"SoFast/models"
	"sync"
)

var (
	wsUsers sync.Map
)

func AddWsUser(user models.User) {
	wsUsers.Store(user.RemoteAddr, user)
}

func RemoveUser(remoteAddr string) {
	wsUsers.Delete(remoteAddr)
}

func AllUser() (users []models.User) {
	wsUsers.Range(func(key, value any) bool {
		users = append(users, value.(models.User))
		return true
	})
	return
}
