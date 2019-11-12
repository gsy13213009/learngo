package main

import "testing"

func TestLengthOfNonRepeating(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},
		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbb", 1},
		// chiness support
		{"哈哈哈哈哈哈", 1},
	}

	for _, tt := range tests {
		if actual := lengthOfNonRepeating(tt.s); actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}
}

// 性能测试
func BenchLengthOfNonRepeating(b *testing.B) {
	s := "了劳动节啦浪费打两份了大家廖慧峰量大富丽达奥拉夫量大付了定金撒了反馈了"
	ans := 8

	for i := 0; i < b.N; i++ {
		actual :=  lengthOfNonRepeating(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}