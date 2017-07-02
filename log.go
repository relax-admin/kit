package kit

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func Info(act interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	fmt.Printf("%s:%s:%d:--method:%v--got: %#v-\n\n", Log_Info, filepath.Base(file), line, runtime.FuncForPC(pc).Name(), act)
}

func Debug(act interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	fmt.Printf("%s:%s:%d:--method:%v--got: %#v-\n\n", Log_Info, filepath.Base(file), line, runtime.FuncForPC(pc).Name(), act)
}
