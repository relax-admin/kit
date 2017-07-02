package kit

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func ToString(raw interface{}) string {
	switch v := raw.(type) {
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(int64(v), 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(uint64(v), 10)
	case bool:
		return strconv.FormatBool(bool(v))

	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(float64(v), 'f', -1, 64)
	case string:
		return string(v)
	default:
		return ""
	}
	return ""
}

//Black(30),Red(31),Green(32),Yellow(33),Blue(34),Magenta(35),Cyan(36),White(37)
func InfoTest(act interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	//fmt.Println("Name of function: " + runtime.FuncForPC(pc).Name())
	fmt.Printf("\033[35m%s:%d:\n\n\tmethod:%v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, runtime.FuncForPC(pc).Name(), act)
}

//Black(90),Red(91),Green(92),Yellow(93),Blue(94),Magenta(95),Cyan(96),White(97)
func Warning(act interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	//fmt.Println("Name of function: " + runtime.FuncForPC(pc).Name())
	fmt.Printf("\033[33m%s:%d:\n\n\tmethod:%v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, runtime.FuncForPC(pc).Name(), act)
}

func LogInfo(act interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	fmt.Printf("%s:%s:%d:--method:%v--got: %#v-\n", LogInfo, filepath.Base(file), line, runtime.FuncForPC(pc).Name(), act)
}

func LogError(act interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	fmt.Printf("%s:%s:%d:--method:%v--got: %#v-\n", LogError, filepath.Base(file), line, runtime.FuncForPC(pc).Name(), act)
}

func Contains(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func ContainsInt(list []int64, value int64) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func ContainsInts(list string, value int64) (has bool, err error) {
	slist := strings.Split(list, ",")
	for _, v := range slist {
		var v1 int64
		v1, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			return
		}
		if v1 == value {
			has = true
			return
		}
	}
	return
}
