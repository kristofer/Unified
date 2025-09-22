package main

import (
	"fmt"
	"tinygo.org/x/go-llvm"
)

func main() {
	context := llvm.NewContext()
	defer context.Dispose()
	
	module := context.NewModule("test")
	defer module.Dispose()
	
	builder := context.NewBuilder()
	defer builder.Dispose()
	
	fmt.Printf("LLVM context: %v\n", context)
	fmt.Printf("LLVM module: %s\n", module.String())
}
