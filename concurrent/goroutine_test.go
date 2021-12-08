// 协程泄露是指协程创建后，长时间得不到释放，并且还在不断地创建新的协程，最终导致内存耗尽，程序崩溃。常见的导致协程泄露的场景有以下几种
// 缺少接收器，导致发送阻塞
// 这个例子中，每执行一次 query，则启动10个协程向信道 ch 发送数字 0，但只接收了一次，导致 9 个协程被阻塞，不能退出
package concurrent

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"testing"
	"time"
)

func query() int {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func() { ch <- 0 }()
	}
	return <-ch
}

func TestQuery(t *testing.T) {
	for i := 0; i < 4; i++ {
		query()
		fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
	}
}

/*
缺少发送器，导致接收阻塞
那同样的，如果启动 1000 个协程接收信道的信息，但信道并不会发送那么多次的信息，也会导致接收协程被阻塞，不能退出。

死锁(dead lock)
两个或两个以上的协程在执行过程中，由于竞争资源或者由于彼此通信而造成阻塞，这种情况下，也会导致协程被阻塞，不能退出。

无限循环(infinite loops)
这个例子中，为了避免网络等问题，采用了无限重试的方式，发送 HTTP 请求，直到获取到数据。那如果 HTTP 服务宕机，永远不可达，导致协程不能退出，发生泄漏。
*/

func request(url string, wg *sync.WaitGroup) {
	i := 0
	for {
		if _, err := http.Get(url); err == nil {
			// write to db
			break
		}
		i++
		if i >= 3 {
			break
		}
		time.Sleep(time.Second)
	}
	wg.Done()
}

func TestRequest(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go request(fmt.Sprintf("https://whatever.com/%d", i), &wg)
	}
	wg.Wait()
}
