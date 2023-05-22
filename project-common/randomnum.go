package common

import (
	"math/rand"
	"time"
)

// RandomNum 生成一个[min,max]随机数
func RandomNum(min int, max int) int {
	rand.New(rand.NewSource(time.Now().UnixNano())) //以当前纳秒数作为随机数种子
	return rand.Intn(max+1) + min
}
