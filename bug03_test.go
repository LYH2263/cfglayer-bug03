package cfglayer_test

import (
	"errors"
	"testing"

	cl "github.com/LYH2263/go-cfglayer"
)

func TestBug03_PopEmptyStackErr(t *testing.T) {
	m := newMerger(t)
	_, err := m.PopLayer()
	if err == nil {
		t.Fatal("want err")
	}
	if !errors.Is(err, cl.ErrEmptyStack) {
		t.Fatalf("want ErrEmptyStack got %v", err)
	}
}
