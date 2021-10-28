package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	// создаём контекст, который отменится при получении сигнала Interrupt (нажатие ctrl+c)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// вызовем медленную операцию с поддержкой отмены контекста
	slowOperation(ctx)
}

// Медленная операция, которую можно отменить через контекст
func slowOperation(ctx context.Context) {
	fmt.Println("начинаем медленную операцию")

	// select похож на switch/case, но выполняет он не ветку с подходящим булевым условием,
	// а ту ветку, данные из которой получены первыми.
	select {
	case <-ctx.Done():
		// контекст был отменён, а медленная операция не успела завершиться.
		// здесь у нас есть шанс выполнить компенсирующую операцию:
		// - вернуть незавершённую задачу в очередь
		// - пометить задачу в БД как "отменённую".
		// - или хотя бы просто написать в лог.
		fmt.Println("медленная операция прервана")
		break
	case <-time.After(time.Second * 5):
		fmt.Println("медленная операция завершена")
		break
	}
}
