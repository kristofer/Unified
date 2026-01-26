package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestIntegrationSimpleReturn(t *testing.T) {
	testCompileAndRun(t, "../../test/integration/simple_return.uni", 42)
}

func TestIntegrationFunctionCall(t *testing.T) {
	testCompileAndRun(t, "../../test/integration/function_call.uni", 30)
}

func TestIntegrationLocalVariables(t *testing.T) {
	testCompileAndRun(t, "../../test/integration/local_variables.uni", 30)
}

func TestIntegrationWhileFactorial(t *testing.T) {
	testCompileAndRun(t, "../../test/integration/while_factorial.uni", 120)
}

func TestIntegrationForSum(t *testing.T) {
	testCompileAndRun(t, "../../test/integration/for_sum.uni", 55)
}

func TestIntegrationNestedLoops(t *testing.T) {
	testCompileAndRun(t, "../../test/integration/nested_loops.uni", 18)
}

func TestIntegrationFizzBuzz(t *testing.T) {
	testCompileAndRun(t, "../../test/integration/fizzbuzz_complete.uni", 15)
}

// TODO: Enable once if statement parsing is fixed
// func TestIntegrationIfElse(t *testing.T) {
// 	testCompileAndRun(t, "../../test/integration/if_else.uni", 1)
// }

// Helper function to compile and run a .uni file
func testCompileAndRun(t *testing.T, filePath string, expectedExitCode int) {
	t.Helper()

	// Build the compiler first
	buildCmd := exec.Command("go", "build", "-o", "../../bin/unified-test", ".")
	buildOut, err := buildCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to build compiler: %v\nOutput: %s", err, buildOut)
	}

	// Get absolute path to test file
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		t.Fatalf("Failed to get absolute path: %v", err)
	}

	// Run the compiler
	cmd := exec.Command("../../bin/unified-test", "--input", absPath)
	output, err := cmd.CombinedOutput()
	
	// Check exit code
	exitCode := 0
	if exitErr, ok := err.(*exec.ExitError); ok {
		exitCode = exitErr.ExitCode()
	} else if err != nil {
		t.Fatalf("Failed to run compiler: %v\nOutput: %s", err, output)
	}

	if exitCode != expectedExitCode {
		t.Errorf("Expected exit code %d, got %d\nOutput: %s", expectedExitCode, exitCode, output)
	}

	// Clean up test binary
	os.Remove("../../bin/unified-test")
}

func TestMain(m *testing.M) {
	// Clean up any leftover test binaries
	os.Remove("../../bin/unified-test")
	
	// Run tests
	code := m.Run()
	
	// Clean up again
	os.Remove("../../bin/unified-test")
	
	os.Exit(code)
}

func TestIntegrationStructPoint(t *testing.T) {
testCompileAndRun(t, "../../test/point.uni", 10)
}

func TestIntegrationStructRectangle(t *testing.T) {
testCompileAndRun(t, "../../test/rectangle.uni", 50)
}

func TestIntegrationStructNested(t *testing.T) {
testCompileAndRun(t, "../../test/nested_structs.uni", 57) // 12345 % 256 = 57
}
