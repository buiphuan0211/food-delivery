package common

import (
	"math/rand"
	"time"
)

var letters = []rune("abcxyzABCXYZ")

func randSequence(n int) string {
	b := make([]rune, n)

	// Nếu sử dựng trục tiếp  rand.Intn()
	// => tất cá kết quá sẽ giống nhau, phải chạy rand.Seet mới rand ra các gtri khac nhau được
	// => tự custom
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := range b {
		b[i] = letters[r1.Intn(99999)%len(letters)]
	}

	return string(b)
}

func GenSalt(length int) string {
	if length < 0 {
		length = 50
	}

	return randSequence(length)
}

// Tại sao rand cần Seet => 1 chuỗi rand, chỉ cần có được Seet -> sẽ lấy được toàn bộ chuỗi của vòng random (số tiếp theo là gì)
