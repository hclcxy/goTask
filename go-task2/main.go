package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// 1.定义一个整数指针
	var ptr *int = new(int)      // 使用new函数创建一个指向整数的指针
	*ptr = 10                    // 初始化指针指向的值为5
	fmt.Println("指针修改之前:", *ptr) // 输出修改前的值
	modifyValue(ptr)             // 调用函数修改值
	fmt.Println("指针修改之后:", *ptr) // 输出修改后的值

	// 2.输出修改后的值
	var slicePtr *[]int = &[]int{1, 2, 3, 4, 5} // 创建一个整数切片的指针
	fmt.Println("切片修改之前:", *slicePtr)           // 输出修改前的切片
	modifySlice(slicePtr)                       // 调用函数修改切片
	fmt.Println("切片修改之后:", *slicePtr)           // 输出修改后的切片

	//3.channel的使用
	channelExample() // 调用channel示例函数
	// 4.使用sync.Mutex保护共享计数器
	syneExample() // 调用sync.Mutex示例函数
	// 5.使用原子操作实现无锁计数器
	atomicCounterExample() // 调用原子操作示例函数
}

// 1.题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
// 在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。
func modifyValue(ptr *int) {
	// 2.将该整数指针指向的值增加10。
	*ptr += 10

}

// 2。题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。
func modifySlice(slicePtr *[]int) {
	for i := range *slicePtr {
		(*slicePtr)[i] *= 2 // 将切片中的每个元素乘以2
	}
}

// 3.channel的使用
func channelExample() {
	ch := make(chan int) // 创建一个整数类型的通道

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i // 向通道发送数据
		}
		close(ch) // 关闭通道
	}()

	for value := range ch { // 从通道接收数据
		fmt.Println("从通道接收到的值:", value)
	}
}

// 4.题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，
// 每个协程对计数器进行1000次递增操作，最后输出计数器的值
func syneExample() {
	var counter int
	var mu sync.Mutex // 创建一个互斥锁

	var wg sync.WaitGroup // 创建一个等待组

	for i := 0; i < 10; i++ {
		wg.Add(1) // 每启动一个协程，等待组计数加1
		go func() {
			defer wg.Done() // 协程结束时，等待组计数减1
			for j := 0; j < 1000; j++ {
				mu.Lock()   // 锁定互斥锁
				counter++   // 对计数器进行递增操作
				mu.Unlock() // 解锁互斥锁
			}
		}()
	}

	wg.Wait()                           // 等待所有协程完成
	fmt.Println("互斥锁最终计数器的值:", counter) // 输出计数器的值
}

// 5.题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，
// 每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。
func atomicCounterExample() {
	var counter int64 // 使用int64类型的计数器

	var wg sync.WaitGroup // 创建一个等待组

	for i := 0; i < 10; i++ {
		wg.Add(1) // 每启动一个协程，等待组计数加1
		go func() {
			defer wg.Done() // 协程结束时，等待组计数减1
			for j := 0; j < 1000; j++ {
				// 使用原子操作对计数器进行递增
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	wg.Wait()                                // 等待所有协程完成
	fmt.Println("使用原子操作最终无锁计数器的值:", counter) // 输出计数器的值
}
