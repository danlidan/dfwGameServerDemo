package gate

import (
	"server/login"
	"server/msg"
)

//该模块用于处理客户端的接入,根据消息不同分发到各个模块处理
func init() {
	msg.Processor.SetRouter(&msg.Hello{}, login.ChanRPC)
}
