package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/name5566/leaf/log"
	"github.com/shopspring/decimal"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func ConvertString2Structure(s string, v interface{}) error {
	if s == "" {
		return nil
	}
	if err := json.Unmarshal([]byte(s), v); err != nil {
		log.Error("ConvertString2Structure error:%+v", err, v)
		return err
	}
	return nil
}

func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	if bytes, ok := v.([]byte); ok {
		return string(bytes)
	}

	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.String:
		return v.(string)
		//case reflect.Int, reflect.Float32, reflect.Float64, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8, reflect.Uint, reflect.Uint32, reflect.Uint64, reflect.Uint8:
		//	bytes, _ := json.Marshal(v)
		//	return string(bytes)

	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8, reflect.Uint, reflect.Uint32, reflect.Uint64, reflect.Uint8:
		//i, _ := v.(int)
		//return strconv.Itoa(i)
		return fmt.Sprintf("%d", v)
	case reflect.Float32, reflect.Float64:
		bytes, _ := json.Marshal(v)
		return string(bytes)
		//return str

	case reflect.Struct:
		return ToJson(v)
	default:
		return ToJson(v)
	}
}

func ToJson(v interface{}) string {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.Encode(&v)
	return strings.TrimRight(buf.String(), "\n\r")
}

func ToFloat(a int) float32 {
	return float32(a) / float32(10000)
}

func SplicingId(order, lv int) int {
	strconv.Itoa(order)
	var level string
	if lv < 10 {
		level = "00" + strconv.Itoa(lv)
	}
	if lv >= 10 && lv < 100 {
		level = "0" + strconv.Itoa(lv)
	}
	if lv >= 100 {
		level = strconv.Itoa(lv)
	}
	id := "1" + strconv.Itoa(order) + level
	Id, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err.Error())
	}
	if order == 9 && lv == 120 {
		return 20000
	}
	return Id
}

func Reserve4bits(num float64) float32 {
	f, _ := decimal.NewFromFloat(num).Round(4).Float64()
	return float32(f)
}

func SortMapKey(data map[int64][]int32) int64 {
	arr := make([]int, 0)
	for k, _ := range data {
		arr = append(arr, int(k))
	}
	sort.Ints(arr)
	return int64(arr[len(arr)-1])
}

func SplicingMythicsLvId(id, lv int32) int {
	strId := strconv.Itoa(int(id))
	lvId := "1" + strId[6:] + "000"
	atoi, _ := strconv.Atoi(lvId)
	return atoi + int(lv)
}

func DeleteSlice(idx int, arr []int) []int {
	copyArr := make([]int, 0)
	for k, v := range arr {
		if k != idx {
			copyArr = append(copyArr, v)
		}
	}
	return copyArr
}

// SliceSort 从大到小
func SliceSort(arr []int) []int {
	newArr := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		newArr = append(newArr, arr[i])
	}
	return newArr
}
