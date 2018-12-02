package chapter5

import (
	"fmt"
	"time"
)

func Do() {
	fmt.Println("[Do] start!", time.Now())

	// [goroutine]
	// Go 프로그램 안에서 동시에 독립적으로 실행되는 흐름의 단위! - 스레드와 비슷
	// 스레드와 달리 고루틴은 수 킬로바이트 정도의 아주 적은 리소스에서 동작하므로
	// 한 프로세스에 수천, 수만 개의 고루틴을 동작시킬 수 있다.
	// 고루틴은 정보를 공유하는 방식이 아닌 서로 메시지를 주고 받는 방식이라 Memory Lock 처리가 필요없고, 구현도 어렵지 않다.
	go long()
	go short()

	time.Sleep(5 * time.Second) // 5초 대기
	fmt.Println("[Do] stop!", time.Now())

	// [goroutine 사용시 주의 사항!]
	// 실행 중인 goroutine 이 있어도 Main 함수가 종료되면 프로그램이 종료된다.
	// 그래서 아직 실행 중인 고루틴이 있디마녀 메인 함수가 종료되지 않게 해야 한다.
	// 프로그램이 비정상 종료가 되는 것을 피하기 위해서는 Main 함수는 충분히 긴 시간 동안 대기...ㅠㅠ
	// 근데 각 goroutine 이 얼마나 오랫동안 수행할지 모르니.... 이건 좀..
	// Go에서 제안하는 방법
	// goroutine 이 종료 상황을  확인할 수 있게 채널을 만들고, 만든 채널을 통해서 종료 메시지를 대기 시키는 것!
}

func long() {
	fmt.Println("func long start!", time.Now())
	time.Sleep(3 * time.Second) // 3초 대기
	fmt.Println("func long stop!", time.Now())
}

func short() {
	fmt.Println("func short start!", time.Now())
	time.Sleep(1 * time.Second) // 1초 대기
	fmt.Println("func short stop!", time.Now())
}

func DoWithChannel() {
	fmt.Println("[DoWithChannel] start!", time.Now())

	// [channel : goroutine 끼리 정보를 교환하고 실행의 흐름을 동기화하기 위해 사용.]
	// channel 선언은 일반 변수와 동일하게 선언하고, make() 함수로 생성!
	// channel을 정의할 떼는 chan 키워드로 채널을 통해 주고받을 데이터의 타입을 지정!!!
	done := make(chan bool)
	// channel을 정의할 때 지정한 테이터 타입만 채널을 통해 주고 받을 수 있다!
	// ch <- "msg" // ch 채널에 msg 전송
	// m := <- ch // ch 채널로부터 메시지 수신

	// channel에 값을 전송하거나 수신할 때 채널이 준비되지 않으면 준비될 때까지 대기한다.
	// 즉, 채널이 비어 있지 않으면 전송되지 않고, 채널에 값이 들어오기 전에는 수신되지 않는다.

	go longWithChannel(done)
	go shortWithChannel(done)

	<-done
	<-done

	time.Sleep(5 * time.Second) // 5초 대기
	fmt.Println("[DoWithChannel] stop!", time.Now())

	// [channel 사용시 주의 사항]
	// 함수와 마찬가지로 채널도 값에 의한 호출 방식으로 값을 전달한다. 즉, 실제 값이 복사되어 전달되므로
	// bool, int, float64 등의 값을 전달하는 것이 안전하다! 문자열과 배열도 변하지 않는 값이므로 채널의 값으로 사용해도 안전
	// 하지만 포인터 변수나 참조 값(슬라이스, 맵)을 채널로 전달할 때는 주소 값이 전달되므로 값을 보내는 고루틴과
	// 값을 받는 고루틴에서 값을 동시에 수정하면 예상치 못한 결과가 발생할 수 있다.
	// 그래서 포인터나 참조 값을 채널로 전달할 때는 여러 고루틴에서 값을 동시 수정하지 않게 주의!
	// 가장 간단한 방법은 여러 고루틴에서 참조 값에 동시에 접근할 수 없게 뮤텍스로 제한하는 것.

	// [channel 방향]
	// 기본적으로 양방향 통신이 가능한 상태로 만들어지지만 실제로는 채널을 구조체의 필드로 사용하거나 함수의 매개변수로 전달
	// 하는 것이 일반적인데, 이때는 채널이 대부분 단방향으로만 사용된다! 즉, 수신 전용 채널 또는 송신 전용 채널이다!
	// chan<- string // 송신 전용 채널
	// <-chan string // 수신 전용 채널

	// [channel close]
	// close(ch)
	// 채널을 닫은 후에 메시지를 전송하면 에러 발생!
	// v, ok := <-ch // ok가 false 라면 채널에 더는 수신할 값이 없고 채널이 닫힌 상태이다.
	// 채널을 닫는 것은 필수가 아니다. 수신자가 채널에 더 이상 들어올 값이 없다는 것을 알아야 할 때만 채널을 닫아주면 된다.
}

func longWithChannel(done chan bool) {
	fmt.Println("func longWithChannel start!", time.Now())
	time.Sleep(3 * time.Second) // 3초 대기
	fmt.Println("func longWithChannel stop!", time.Now())

	done <- true
}

func shortWithChannel(done chan bool) {
	fmt.Println("func shortWithChannel start!", time.Now())
	time.Sleep(1 * time.Second) // 1초 대기
	fmt.Println("func shortWithChannel stop!", time.Now())

	done <- true
}

func DoChannelWithDeadLock() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	c <- 3 // fatal error: all goroutines are asleep - deadlock!

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func DoChannel() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	go func() { c <- 3 }()

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func fibonacci(c, quit chan int) {
	// [select]
	// select문은 하나의 고루틴이 여러 채널과 통신할 때 사용한다.
	// case로 여러 채널을 대기시키고 있다가 실행 가능 상태가 된 채널이 있으면 해당 케이스를 수행.

	// select문에서 default case를 지정하면 case에 지정된 모든 채널이 사용 가능 상태가 아닐 때 default case를 수행

	/*
		c1 := make(chan int)
		c2 := make(chan int)
		// ...
		select {
		case <-c1:
			// c1 채널에 값이 전달됐을 때 수행
		case <-c2:
			// c2 채널에 값이 전달됐을 때 수행
		default:
			// case에서 대기하고 있는 채널에 값이 전달되지 않았을 때 수행
		}
	*/

	x, y := 0, 1
	for i := 0; ; i++ {
		select {
		case c <- x:
			fmt.Println("SEND??? ", x)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
			// default:
			// 	fmt.Println("......... ", i)
		}
	}
}

func DoSelect() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("RECEIVE??? ", <-c) // 수신?
		}
		quit <- 0 // 송신?
	}()
	fibonacci(c, quit)
}
