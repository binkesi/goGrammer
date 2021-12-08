// 对于无缓冲的 channel，发送方将阻塞该信道，直到接收方从该信道接收到数据为止，而接收方也将阻塞该信道，直到发送方将数据发送到该信道中为止。
// 对于有缓存的 channel，发送方在没有空插槽（缓冲区使用完）的情况下阻塞，而接收方在信道为空的情况下阻塞。
package concurrent

import (
	"fmt"
	"testing"
	"time"
)

// channel without buffer
func TestChannel(t *testing.T) {
	st := time.Now()
	ch := make(chan bool)
	go func() {
		time.Sleep(time.Second * 2)
		<-ch
	}()
	ch <- true
	fmt.Printf("cost %s\n", time.Since(st))
	time.Sleep(time.Second * 5)
}

// channel with buffer
func TestChannelBuf(t *testing.T) {
	st := time.Now()
	ch := make(chan bool, 2)
	go func() {
		time.Sleep(time.Second * 2)
		<-ch
	}()
	ch <- true
	ch <- true
	fmt.Printf("cost %s \n", time.Since(st))
	ch <- true
	fmt.Printf("cost %s\n", time.Since(st))
	time.Sleep(time.Second * 5)
}
