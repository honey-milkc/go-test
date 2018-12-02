package chapter5

import (
	"testing"
)

func TestDo(t *testing.T) {
	t.Log("test Do")

	Do()
}

func TestDoWithChannel(t *testing.T) {
	t.Log("test Do With Channel")
	DoWithChannel()
}

func TestDoChannelWithDeadLock(t *testing.T) {
	t.Log("test Do Channel With DeadLock")
	DoChannelWithDeadLock()
}

func TestDoChannel(t *testing.T) {
	t.Log("test Do Channel")
	DoChannel()
}

func TestDoSelect(t *testing.T) {
	t.Log("test Do Select")
	DoSelect()
}

func TestDoSyncMutex(t *testing.T) {
	t.Log("test Do Sync Mutex\n")
	for i := 0; i < 10; i++ {
		// 1000이 나올 것 같았지만 1000보다 작은 값이 출력.
		// 이는 여러 고루틴이 counter 내부 필드 i의 값을 동시에 수정하려고 해서 경쟁 상태가 만들어지고 이로 인해 정확한 결과 NONO!
		notUseMutexC := newCounter()
		t.Logf("[%d] Not Use Mutex : %d\n", i, DoSyncMutex(notUseMutexC))

		// mutex를 이용한 경우는 늘 1000이 나온다.
		useMutexC := newCounterM()
		t.Logf("[%d] Use Mutex : %d\n", i, DoSyncMutex(useMutexC))

		// once 수행
		useMutexOnceC := newCounterO()
		t.Logf("[%d] Use Mutex Once : %d\n\n", i, DoSyncMutex(useMutexOnceC))
	}
}

func TestDoWaitGroup(t *testing.T) {
	t.Log("test Do Sync WaitGroup\n")
	t.Log(DoWaitGroup())
}
