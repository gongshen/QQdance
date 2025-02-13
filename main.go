package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("QQ炫舞自动化脚本已启动...")
	keys := []string{"up", "down", "left", "right", "space"}

	// 捕获屏幕

	// 模拟按键
	for i := 0; i < 10; i++ { // 模拟 10 次按键
		for _, key := range keys {
			// 按下按键
			robotgo.KeyTap(key)
			fmt.Printf("按下按键: %s\n", key)

			// 等待一段时间（模拟按键间隔）
			time.Sleep(200 * time.Millisecond)
		}
	}
	fmt.Println("模拟按键完成！")
}
