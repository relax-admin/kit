package log

import (
	"fmt"
	"path/filepath"
	"runtime"
)

const (
	Log_Info  = "log_info"
	Log_Error = "log_err"
)

func Info(act interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	fmt.Printf("%s:%s:%d:--method:%v--got: %#v-\n\n", Log_Info, filepath.Base(file), line, runtime.FuncForPC(pc).Name(), act)
}

func Error(act interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	fmt.Printf("%s:%s:%d:--method:%v--got: %#v-\n\n", Log_Info, filepath.Base(file), line, runtime.FuncForPC(pc).Name(), act)
}
