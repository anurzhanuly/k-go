package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	// создаём контекст, который отменяется через 2 секунды
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

	// вызовем из нашей программы bash-скрипт "sleep 5"
	output, err := exec.CommandContext(ctx, "sleep", "5").CombinedOutput()

	// ожидаем, что вернётся ошибка. какая?
	fmt.Println("command output: ", string(output))
	fmt.Println("error: ", err)

	// на эту строчку пока можно не обращать внимания
	cancel()
}
