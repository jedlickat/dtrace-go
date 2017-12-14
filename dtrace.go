package main

import (
	"fmt"
	"github.com/jedlickat/libdtrace-go"
)

func main() {
	handle := dtrace.NewDTraceHandle()
	fmt.Println("DTrace-go")
	fmt.Println("available probes:")
	for _, probe := range dtrace.GetProbes(handle) {
		fmt.Printf("%d: %s: %s: %s: %s\n", probe.ProbeID,
			probe.ProviderName, probe.ModuleName,
			probe.FunctionName, probe.ProbeName);
	}
	dtrace.FreeDTraceHandle(handle)
}
