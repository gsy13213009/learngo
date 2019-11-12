package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 3, 4, 5, 4, 6, 6}
	var grid [4][5]bool
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(grid)
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	// 使用range，意义明确，美观，相当于java的for each
	for i := range arr3 {
		fmt.Println(arr3[i])
	}
	// 可通过下划线省略变量
	for _, v := range arr3 {
		fmt.Println(v)
	}
}

// 获取数组中最大值
// []int 是切片，[5]int 是数组
func getMax(arr []int) (maxi, maxv int) {
	maxi = -1
	maxv = -1
	for i, v := range arr {
		if v > maxv {
			maxi, maxv = i, v
		}
	}
	return
}
