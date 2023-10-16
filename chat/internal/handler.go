package internal

import (
	"reflect"
)

func Init() {
	
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}
