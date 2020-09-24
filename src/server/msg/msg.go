package msg

import (
	"leaf/network/protobuf"
)

var Processor = protobuf.NewProcessor()

//protoc --go_out=. xxx.proto
//proto消息格式： len 2Byte | id 2Byte | protobuf message
//len 为 id + message的长度，不包括len本身
//len & id 均为大端编码， 根据id区分消息类型

func init() {
	Processor.Register(&UserLogin{})    // id : 0 每行递增1
	Processor.Register(&UserLoginRsp{}) // 1
}
