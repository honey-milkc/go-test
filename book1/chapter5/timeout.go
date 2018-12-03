package chapter5

import (
	"fmt"
	"time"
)

func DoTimeout() {
	quit := make(chan struct{})
	done := process(quit)
	timeout := time.After(2 * time.Second)

	select {
	case d := <-done:
		fmt.Println(d)
	case <-timeout:
		fmt.Println("timeout!!!!")
		quit <- struct{}{}
	}
}

// timeout에 대한 처리는 다음 3가지로 처리할 수 있다.
// 1. 아무일도 하지 않음
//    - process() 작업 이후 코드에 영향을 주지 않는다면 따로 처리 안해도 된다.
//      단, select 구문이 종료된 후에도 process() 함수는 계속 동작해 처리가 완료된 후 done 채널로 값 전송.
//      fatal error: all goroutines are asleep - deadlock!
// 2. done 채널을 닫음
//    - process() 함수에 타임아웃 처리가 되었다는 것을 알리기 위해서 done 채널을 닫을 수 있지만
//      done 채널을 임의로 닫아버리면 process() 함수에서 처리를 완료한 후 done 채널에 처리한 결과를 전송할 대 런타임 에러 발생
//      (panic: send on closed channel)
//      process()에서  done 채널은 외부에서 닫힐 수 있다는 것을 고려해서 작성해야 한다.
// 3. process() 함수에 타임아웃 메시지 전송 <- 이방법을 처리
//    - process() 에서 처리하는 작업이 리소스를 많이 사용하는 무거운 작업이라면 타임아웃 메시지를 process()로 전달하여
//      타임아웃 후에 작업을 바로 종료하게 하는 것이 좋다.
func process(quit <-chan struct{}) chan string {
	done := make(chan string)

	go func() {
		go func() {
			defer fmt.Println("long long")
			time.Sleep(10 * time.Second) // heavy job
			done <- "Complete!"
		}()

		<-quit
		return
	}()

	return done
}
