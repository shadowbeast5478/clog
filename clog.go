package clog

import (
    "fmt"
    "strings"
    "runtime"
    "path/filepath"
)

const (
    black int = 30 + iota
    red
    green
    yellow
    blue
    magenta
    cyan
    white
)


func log(color int, level string, formating string, args... interface{}) {
    filename, line, funcname := "", 0, ""
    pc, filename, line, ok := runtime.Caller(3)
    if ok {
        funcname = runtime.FuncForPC(pc).Name()
        funcname = filepath.Ext(funcname)
        funcname = strings.TrimPrefix(funcname, ".")

        filename = filepath.Base(filename)
    }

    fmt.Printf("\033[0;40;%dm %s \033[0m\033[1;44;32m%s:%d # %s:\033[0m %s\n", color, level, filename, line, funcname, fmt.Sprintf(formating, args...))
}

func Info(formating string, args... interface{}) {
    log(green, "[INFO]  ", formating, args...)
}

func Debug(formating string, args... interface{}) {
    log(blue, "[DEBUG] ", formating, args...)
}

func Warning(formating string, args... interface{}) {
    log(yellow, "[WARNING]", formating, args...)
}

func Error(formating string, args... interface{}) {
    log(red, "[ERROR] ", formating, args...)
}
