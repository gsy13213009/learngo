package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 创建
	m := map[string]string{
		"name":   "ccmouse",
		"cource": "golang",
		"site":   "imooc",
	}
	m2 := make(map[string]int)
	var m3 map[string]int // 默认值是nil，nil可以参与运算，可以与map[]混用
	fmt.Println(m)        // map[name:ccmouse cource:golang site:imooc]
	fmt.Println(m2)       // map[]
	fmt.Println(m3)       // map[]
	// map的遍历
	for k, v := range m {
		fmt.Println(k, v)
	}
	// 获取值
	courceName, ok := m["cource"]
	fmt.Println(courceName, ok)    // golang true
	courcehhhh, ok := m["courceh"] // 如果不存在，第二个参数返回false，第一为空串
	fmt.Println(courcehhhh, ok)    //  false
	// 删除元素
	delete(m, "name")
	lengthOfNonRepeating("adbbsafb")
	lengthOfNonRepeating("")
	lengthOfNonRepeating("b")
	lengthOfNonRepeating("bbbbbbbbbbb")
	lengthOfNonRepeating("哈哈哈")
	lengthOfNonRepeatingRoune()
}

// 查找最长不重复子串
// lastOccurred[x] 不存在，或者 < start -> 无需操作
// lastOccurred[x] >= start -> 更新start
// 更新lastOccurred[x] ,更新maxLength
func lengthOfNonRepeating(s string) {
	lastOccurred := make(map[int32]int)
	start, maxLength := 0, 0
	for i, ch := range []rune(s) { // 转成rune，就是国际化的
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		length := i - start + 1
		if length > maxLength {
			maxLength = length
		}
		lastOccurred[ch] = i
	}
	fmt.Println(s, " max = ", maxLength)
}
func lengthOfNonRepeatingRoune() {
	// 每个中文3字节, UTF-8编码，可变长，英文1字节，中文3字节
	s := "Yes我爱慕课网!" // 59 65 73 (E6 88 91) (E7 88 B1) (E6 85 95) (E8 AF BE) (E7 BD 91) 21
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b) // 59 65 73 E6 88 91 E7 88 B1 E6 85 95 E8 AF BE E7 BD 91 21
	}
	fmt.Println()
	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X) ", i, ch) // (0 59) (1 65) (2 73) (3 6211) (6 7231) (9 6155) (12 8BFE) (15 7F51) (18 21)
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s)) // Rune count: 9
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch) // Y e s 我 爱 慕 课 网 !
	}
	fmt.Println()
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
}
