package main

import (
	"fmt"
	"time"
)

func main() {
	slowOperation()
}

// Медленная операция, которая прервётся, если запустить программу и нажать ctrl+c
func slowOperation() {
	fmt.Println("начинаем медленную операцию")
	time.Sleep(time.Second * 5)
	fmt.Println("медленная операция завершена")
}
