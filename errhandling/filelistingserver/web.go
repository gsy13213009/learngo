package main

import (
	"../filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

// 定义appHandler，处理逻辑并抛出错误
type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		// 使用hander去请求内容，如果有错误handler会返回错误
		err := handler(writer, request)
		if err != nil {
			log.Printf("error occurred handling request: %s", err.Error())
			// 处理自定义的用户错误
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}

	}
}

// 自定义的用户错误
type userError interface {
	error
	Message() string
}

func main() {
	// 使用errWrapper函数包装请求逻辑，处理请求错误
	http.HandleFunc("/", errWrapper(filelisting.HandlerFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
