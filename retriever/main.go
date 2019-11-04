package main

import (
	"../retriever/real"
	"fmt"
)
const url = "http://www.imooc.com"
// 定义接口
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

// 接口的组合
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) {
	s.Post(url, map[string]string{
		"contents": "another facked imooc form",
	})
	s.Get("")
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com", map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}
func main() {
	//r := mock.Retriever{Contents: "hhhhhh"}
	raa := real.Retriever{}
	
	fmt.Println(download(raa))
}
