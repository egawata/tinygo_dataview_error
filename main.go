package main

import (
	"fmt"
	"syscall/js"
)

type hoge struct {
	el  js.Value
	cnt int
}

func (h *hoge) Start() {
	h.cnt = 0
	h.el = js.Global().Get("document").Call("getElementById", "output")
	h.el.Call("addEventListener", "incr", js.FuncOf(h.incr))
}

func (h *hoge) incr(this js.Value, p []js.Value) any {
	b := make([]byte, 100_000_000)
	h.cnt += len(b) / 1_000_000
	h.el.Set("innerText", fmt.Sprintf("Allocated %d mbytes", h.cnt))
	return nil
}

func main() {
	h := &hoge{}
	h.Start()

	<-make(chan struct{})
}
