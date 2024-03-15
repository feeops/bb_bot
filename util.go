package main

import (
	"fmt"
	"os"
	"time"
)

func waitExit() {
	fmt.Println("请按任意键退出,如果没有操作，5分钟后自动退出")
	ch := make(chan string)
	go func() {
		var input string
		_, _ = fmt.Scanln(&input)
		ch <- input
	}()

	select {
	case <-ch:
		os.Exit(0)
	case <-time.After(5 * time.Minute):
		fmt.Println("超时退出")
		os.Exit(0)
	}

}
