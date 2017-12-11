package main

import (
	"github.com/jedlickat/libdtrace-go"
)

func main() {
	handle := dtrace.NewDTraceHandle()
	dtrace.GetProbes(handle)
	dtrace.FreeDTraceHandle(handle)
}
