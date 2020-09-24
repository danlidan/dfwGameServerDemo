package internal

import (
	"fmt"
	"reflect"
	"server/msg"

	"leaf/gate"
	"leaf/log"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	//hello消息交予handleHello处理
	handleMsg(&msg.Hello{}, handleHello)
}

func handleHello(args []interface{}) {
	fmt.Println("hello!!!")
	m := args[0].(*msg.Hello)
	a := args[1].(gate.Agent)

	log.Debug("Hello %v", m.Name)

	a.WriteMsg(&msg.HelloRsp{
		Name: "client",
	})

}
