package internal

import (
	"reflect"
	"server/msg"

	"leaf/gate"
	"leaf/log"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	//UserLogin消息交予handleUserLogin处理
	handleMsg(&msg.UserLogin{}, handleUserLogin)
}

func handleUserLogin(args []interface{}) {
	m := args[0].(*msg.UserLogin)
	a := args[1].(gate.Agent)

	log.Debug("Hello %v", m.Name)
	gate.UserOnline(m.Name, a)

	a.WriteMsg(&msg.UserLoginRsp{
		Success: true,
	})

}
