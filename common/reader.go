package common

import (
	"reflect"
	"strings"

	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/recordfile"
)

var config interface{}

const GapCount = 3

func readRf(st interface{}) *recordfile.RecordFile {
	rf, err := recordfile.New(st, GapCount)

	if err != nil {
		log.Fatal("%v", err)
	}
	fn := reflect.TypeOf(st).Name() + ".txt"
	lower := strings.ToLower(fn)
	err = rf.Read("gamedata/" + lower)
	if err != nil {
		log.Fatal("%v: %v", strings.ToLower(fn), err)
	}

	return rf
}
