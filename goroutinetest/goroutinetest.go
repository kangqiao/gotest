package goroutinetest

import (
	"fmt"
	"runtime"
	"sync"
)

// wg用来等待程序完成
var wg sync.WaitGroup

func Test() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 计数器加2, 表示要等待两个goroutine

	// 创建两个goroutine
	log("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	// 等待goroutine结束
	log("Waiting To Finish")
	wg.Wait()

	log("Terminating Program")
}

func log(msg string) {
	fmt.Println(msg)
}

// printPrime 显示5000以内的素数值
func printPrime(prefix string) {
	log("printPrime")
	// 在函数退出时调用Done来通知main函数工作已经完成
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Println(">>", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
