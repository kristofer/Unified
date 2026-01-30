package wasm

import (
	"context"
	"fmt"
	"os"
	"unified-compiler/internal/bytecode"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

// Runtime wraps wazero runtime for executing WASM modules
type Runtime struct {
	runtime wazero.Runtime
	module  api.Module
	ctx     context.Context
}

// NewRuntime creates a new WASM runtime
func NewRuntime() *Runtime {
	ctx := context.Background()
	rt := wazero.NewRuntime(ctx)

	return &Runtime{
		runtime: rt,
		ctx:     ctx,
	}
}

// LoadModule loads a WASM module from binary data
func (r *Runtime) LoadModule(wasmBytes []byte) error {
	// Compile and instantiate the module
	compiledModule, err := r.runtime.CompileModule(r.ctx, wasmBytes)
	if err != nil {
		return fmt.Errorf("failed to compile module: %w", err)
	}

	// Instantiate the module
	moduleConfig := wazero.NewModuleConfig()
	module, err := r.runtime.InstantiateModule(r.ctx, compiledModule, moduleConfig)
	if err != nil {
		return fmt.Errorf("failed to instantiate module: %w", err)
	}

	r.module = module
	return nil
}

// Run executes the main function and returns the result
func (r *Runtime) Run() (bytecode.Value, error) {
	if r.module == nil {
		return bytecode.NewNullValue(), fmt.Errorf("no module loaded")
	}

	// Find and call the main function
	mainFunc := r.module.ExportedFunction("main")
	if mainFunc == nil {
		return bytecode.NewNullValue(), fmt.Errorf("no main function found")
	}

	// Execute main
	results, err := mainFunc.Call(r.ctx)
	if err != nil {
		return bytecode.NewNullValue(), fmt.Errorf("error executing main: %w", err)
	}

	// Convert result to bytecode.Value
	if len(results) == 0 {
		return bytecode.NewIntValue(0), nil
	}

	// Return the first result as i64
	return bytecode.NewIntValue(int64(results[0])), nil
}

// Close closes the runtime and releases resources
func (r *Runtime) Close() error {
	if r.runtime != nil {
		return r.runtime.Close(r.ctx)
	}
	return nil
}

// CompileAndRun is a convenience function that compiles and runs a WASM module
func CompileAndRun(wasmBytes []byte) (bytecode.Value, error) {
	runtime := NewRuntime()
	defer func() {
		if err := runtime.Close(); err != nil {
			// Log error but don't override return value
			fmt.Fprintf(os.Stderr, "Warning: error closing WASM runtime: %v\n", err)
		}
	}()

	if err := runtime.LoadModule(wasmBytes); err != nil {
		return bytecode.NewNullValue(), err
	}

	return runtime.Run()
}
