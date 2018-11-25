package book1_chapter4

import (
	"testing"
)

func TestDoItem(t *testing.T) {
	t.Log("Item")
	DoItem()
}

func TestDoInterface(t *testing.T) {
	t.Log("Interface")
	DoInterface()
}
