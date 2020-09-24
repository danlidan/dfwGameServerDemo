package gate

import (
	"fmt"
	"sync"
)

//Agent->用户名映射
var UserToAgent = struct {
	sync.RWMutex
	m map[Agent]string
}{m: make(map[Agent]string)}

//处理用户上线
func UserOnline(id string, a Agent) {
	UserToAgent.Lock()
	UserToAgent.m[a] = id
	fmt.Println("online", UserToAgent.m)
	UserToAgent.Unlock()
}

//处理用户下线
func UserOffline(a Agent) {
	UserToAgent.Lock()
	delete(UserToAgent.m, a)
	fmt.Println("offline", UserToAgent.m)
	UserToAgent.Unlock()
}
