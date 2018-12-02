package chapter5

import (
	"runtime"
	"sync"
)

type iCounter interface {
	increment()
	cnt() int64
}

type counter struct {
	i int64
}

func newCounter() *counter {
	return &counter{0}
}

func (c *counter) increment() {
	c.i++
}

func (c *counter) cnt() int64 {
	return c.i
}

func DoSyncMutex(c iCounter) int64 {
	// 모든 CPU를 사용하도록
	runtime.GOMAXPROCS(runtime.NumCPU())
	// runtime.NumCPU() : CPU 코어 수
	// runtime.GOMAXPROCS() : 현재 프로그램에서 사용할 CPU의 최대 수 셋팅

	done := make(chan struct{}) // 완료 신호 수신용 채널

	// c.increment()를 실행하는 고루틴 1000개 실행
	for i := 0; i < 1000; i++ {
		go func() {
			c.increment()      // 카운터 값을 1 증가시킴
			done <- struct{}{} // done 채널에 완료 신호 전송
		}()
	}

	for i := 0; i < 1000; i++ {
		<-done
	}

	return c.cnt()

}

type counterM struct {
	i  int64
	mu sync.Mutex // 공유 데이터 i를 보호하기 위한 뮤텍스
}

func newCounterM() *counterM {
	return &counterM{i: 0}
}

func (c *counterM) increment() {
	c.mu.Lock()   // i 값을 변경하는 부분(임계영역)을 뮤텍스로 잠금
	c.i++         // 공유 데이터 변경
	c.mu.Unlock() // i 값을 변경 완료한 후 뮤텍스 잠금해제
}

func (c *counterM) cnt() int64 {
	return c.i
}

// [sync.Once]
// 특정 함수를 한 번만 수행해야 할 때!
type counterO struct {
	i    int64
	mu   sync.Mutex // 공유 데이터 i를 보호하기 위한 뮤텍스
	once sync.Once  // 한 번만 수행될 함수를 지정하기 위한 Once 구조체
}

const initalValue = -500

func newCounterO() *counterO {
	return &counterO{i: 0}
}
func (c *counterO) increment() {
	// i 값 초기화 작업은 한 번만 수행되도록 once의 Do() method로 실행
	c.once.Do(func() {
		c.i = initalValue
	})

	c.mu.Lock()
	c.i++
	c.mu.Unlock()
}

func (c *counterO) cnt() int64 {
	return c.i
}

func DoWaitGroup() int64 {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counterM{i: 0}
	wg := sync.WaitGroup{} // WaitGroup 생성

	for i := 0; i < 1000; i++ {
		wg.Add(1) // 고루틴 수 증가
		go func() {
			defer wg.Done() // 고루틴 종료 시
			c.increment()
		}()
	}

	wg.Wait()

	return c.cnt()
}
