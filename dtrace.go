package main

import (
	"fmt"
	"github.com/jedlickat/libdtrace-go"
)

func main() {
	handle := dtrace.NewDTraceHandle()
	fmt.Printf("handle %p\n", handle)
	dtrace.FreeDTraceHandle(handle)
}
