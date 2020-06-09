package helper

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	_ "reflect"
	"runtime"
	"strings"
)

var skip  = 3
var DEEP  = 1

func Log(s string,p string,l int) {
	if l >3 {
		ENTRY("")
		ENTRY(s+" p=%s", p)
		//DEBUG("Test %s %s", "Hello", "World")
	}
}

func Error(p string,l int) {
	skip = 3
	DEEP = 10
	ENTRY("")
	ENTRY(" p=%s", p)
	//DEBUG("Test %s %s", "Hello", "World")
	os.Exit(1)
}

func DEBUG(formating string, args... interface{}) {
	LOG("DEBUG", formating, args...)
}

func ENTRY(formating string, args... interface{}) {
	LOG("ENTRY", formating, args...)
}

func LOG(level string, formating string, args... interface{}) {

	for i:=0;i<DEEP;i++{
		filename, line, funcname := "???", 0, "???"
		pc, filename, line, ok := runtime.Caller(skip+i)
		// fmt.Println(reflect.TypeOf(pc), reflect.ValueOf(pc))
		if ok {
			funcname = runtime.FuncForPC(pc).Name()       // main.(*MyStruct).foo
			funcname = filepath.Ext(funcname)             // .foo
			funcname = strings.TrimPrefix(funcname, ".")  // foo

			filename = filepath.Base(filename)  // /full/path/basename.go => basename.go
		}

		log.Printf("%s:%d:%s: %s: %s\n", filename, line, funcname, level, fmt.Sprintf(formating, args...))
	}
}
