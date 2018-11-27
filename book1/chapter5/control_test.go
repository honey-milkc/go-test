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
