package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
context 优雅的退出 goroutine
可以退出单个或多个 goroutine
*/
var sw sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sw.Add(1)
	go woker(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 调用返回的 cancel函数退出goroutine
	sw.Wait()
	fmt.Println("退出")
}

func woker(ctx context.Context) {
	defer sw.Done()
	for {
		fmt.Println("work........")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): //监听到调用了 cancel() 函数即可退出
			return
		default:

		}
	}
}
