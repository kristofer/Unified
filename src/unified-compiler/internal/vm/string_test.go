package vm

import (
	"testing"
	"unified-compiler/internal/bytecode"
)

// Test 1: String concatenation with + operator
func TestStringConcatenation(t *testing.T) {
	// Test: "Hello" + " " + "World"
	bc := bytecode.NewBytecode()
	
	// Add constants
	hello := bc.AddConstant(bytecode.NewStringValue("Hello"))
	space := bc.AddConstant(bytecode.NewStringValue(" "))
	world := bc.AddConstant(bytecode.NewStringValue("World"))
	
	// Main function
	bc.AddFunction("main", bc.CurrentPosition())
	bc.AddInstruction(bytecode.OpPush, int64(hello))
	bc.AddInstruction(bytecode.OpPush, int64(space))
	bc.AddInstruction(bytecode.OpConcat, 0)
	bc.AddInstruction(bytecode.OpPush, int64(world))
	bc.AddInstruction(bytecode.OpConcat, 0)
	bc.AddInstruction(bytecode.OpReturnValue, 0)
	bc.AddInstruction(bytecode.OpHalt, 0)
	
	vm := NewVM(bc)
	result, err := vm.Run()
	
	if err != nil {
		t.Fatalf("VM execution failed: %v", err)
	}
	
	if result.Type != bytecode.ValueTypeString {
		t.Errorf("Expected string type, got %v", result.Type)
	}
	
	if result.Str != "Hello World" {
		t.Errorf("Expected 'Hello World', got '%s'", result.Str)
	}
}

// Test 2: String length method
func TestStringLength(t *testing.T) {
	// Test: "Hello".len()
	bc := bytecode.NewBytecode()
	
	str := bc.AddConstant(bytecode.NewStringValue("Hello"))
	
	bc.AddFunction("main", bc.CurrentPosition())
	bc.AddInstruction(bytecode.OpPush, int64(str))
	bc.AddInstruction(bytecode.OpStrLen, 0)
	bc.AddInstruction(bytecode.OpReturnValue, 0)
	bc.AddInstruction(bytecode.OpHalt, 0)
	
	vm := NewVM(bc)
	result, err := vm.Run()
	
	if err != nil {
		t.Fatalf("VM execution failed: %v", err)
	}
	
	if result.Type != bytecode.ValueTypeInt {
		t.Errorf("Expected int type, got %v", result.Type)
	}
	
	if result.Int != 5 {
		t.Errorf("Expected length 5, got %d", result.Int)
	}
}

// Test 3: String is_empty method
func TestStringIsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected bool
	}{
		{"Empty string", "", true},
		{"Non-empty string", "Hello", false},
		{"Whitespace", "  ", false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := bytecode.NewBytecode()
			str := bc.AddConstant(bytecode.NewStringValue(tt.str))
			
			bc.AddFunction("main", bc.CurrentPosition())
			bc.AddInstruction(bytecode.OpPush, int64(str))
			bc.AddInstruction(bytecode.OpStrIsEmpty, 0)
			bc.AddInstruction(bytecode.OpReturnValue, 0)
			bc.AddInstruction(bytecode.OpHalt, 0)
			
			vm := NewVM(bc)
			result, err := vm.Run()
			
			if err != nil {
				t.Fatalf("VM execution failed: %v", err)
			}
			
			if result.Type != bytecode.ValueTypeBool {
				t.Errorf("Expected bool type, got %v", result.Type)
			}
			
			if result.Bool != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result.Bool)
			}
		})
	}
}

// Test 4: String substring method
func TestStringSubstring(t *testing.T) {
	// Test: "Hello World".substring(0, 5) = "Hello"
	bc := bytecode.NewBytecode()
	
	str := bc.AddConstant(bytecode.NewStringValue("Hello World"))
	start := bc.AddConstant(bytecode.NewIntValue(0))
	end := bc.AddConstant(bytecode.NewIntValue(5))
	
	bc.AddFunction("main", bc.CurrentPosition())
	bc.AddInstruction(bytecode.OpPush, int64(str))
	bc.AddInstruction(bytecode.OpPush, int64(start))
	bc.AddInstruction(bytecode.OpPush, int64(end))
	bc.AddInstruction(bytecode.OpStrSubstring, 0)
	bc.AddInstruction(bytecode.OpReturnValue, 0)
	bc.AddInstruction(bytecode.OpHalt, 0)
	
	vm := NewVM(bc)
	result, err := vm.Run()
	
	if err != nil {
		t.Fatalf("VM execution failed: %v", err)
	}
	
	if result.Type != bytecode.ValueTypeString {
		t.Errorf("Expected string type, got %v", result.Type)
	}
	
	if result.Str != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", result.Str)
	}
}

// Test 5: String contains method
func TestStringContains(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		substring string
		expected  bool
	}{
		{"Contains substring", "Hello World", "World", true},
		{"Does not contain", "Hello World", "xyz", false},
		{"Contains at start", "Hello World", "Hello", true},
		{"Empty substring", "Hello", "", true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := bytecode.NewBytecode()
			str := bc.AddConstant(bytecode.NewStringValue(tt.str))
			sub := bc.AddConstant(bytecode.NewStringValue(tt.substring))
			
			bc.AddFunction("main", bc.CurrentPosition())
			bc.AddInstruction(bytecode.OpPush, int64(str))
			bc.AddInstruction(bytecode.OpPush, int64(sub))
			bc.AddInstruction(bytecode.OpStrContains, 0)
			bc.AddInstruction(bytecode.OpReturnValue, 0)
			bc.AddInstruction(bytecode.OpHalt, 0)
			
			vm := NewVM(bc)
			result, err := vm.Run()
			
			if err != nil {
				t.Fatalf("VM execution failed: %v", err)
			}
			
			if result.Type != bytecode.ValueTypeBool {
				t.Errorf("Expected bool type, got %v", result.Type)
			}
			
			if result.Bool != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result.Bool)
			}
		})
	}
}

// Test 6: String starts_with method
func TestStringStartsWith(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		prefix   string
		expected bool
	}{
		{"Starts with prefix", "Hello World", "Hello", true},
		{"Does not start with", "Hello World", "World", false},
		{"Empty prefix", "Hello", "", true},
		{"Same string", "Hello", "Hello", true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := bytecode.NewBytecode()
			str := bc.AddConstant(bytecode.NewStringValue(tt.str))
			prefix := bc.AddConstant(bytecode.NewStringValue(tt.prefix))
			
			bc.AddFunction("main", bc.CurrentPosition())
			bc.AddInstruction(bytecode.OpPush, int64(str))
			bc.AddInstruction(bytecode.OpPush, int64(prefix))
			bc.AddInstruction(bytecode.OpStrStartsWith, 0)
			bc.AddInstruction(bytecode.OpReturnValue, 0)
			bc.AddInstruction(bytecode.OpHalt, 0)
			
			vm := NewVM(bc)
			result, err := vm.Run()
			
			if err != nil {
				t.Fatalf("VM execution failed: %v", err)
			}
			
			if result.Type != bytecode.ValueTypeBool {
				t.Errorf("Expected bool type, got %v", result.Type)
			}
			
			if result.Bool != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result.Bool)
			}
		})
	}
}

// Test 7: String ends_with method
func TestStringEndsWith(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		suffix   string
		expected bool
	}{
		{"Ends with suffix", "Hello World", "World", true},
		{"Does not end with", "Hello World", "Hello", false},
		{"Empty suffix", "Hello", "", true},
		{"Same string", "Hello", "Hello", true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := bytecode.NewBytecode()
			str := bc.AddConstant(bytecode.NewStringValue(tt.str))
			suffix := bc.AddConstant(bytecode.NewStringValue(tt.suffix))
			
			bc.AddFunction("main", bc.CurrentPosition())
			bc.AddInstruction(bytecode.OpPush, int64(str))
			bc.AddInstruction(bytecode.OpPush, int64(suffix))
			bc.AddInstruction(bytecode.OpStrEndsWith, 0)
			bc.AddInstruction(bytecode.OpReturnValue, 0)
			bc.AddInstruction(bytecode.OpHalt, 0)
			
			vm := NewVM(bc)
			result, err := vm.Run()
			
			if err != nil {
				t.Fatalf("VM execution failed: %v", err)
			}
			
			if result.Type != bytecode.ValueTypeBool {
				t.Errorf("Expected bool type, got %v", result.Type)
			}
			
			if result.Bool != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result.Bool)
			}
		})
	}
}

// Test 8: String to_upper method
func TestStringToUpper(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"All lowercase", "hello", "HELLO"},
		{"Mixed case", "HeLLo", "HELLO"},
		{"Already uppercase", "HELLO", "HELLO"},
		{"With spaces", "hello world", "HELLO WORLD"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := bytecode.NewBytecode()
			str := bc.AddConstant(bytecode.NewStringValue(tt.input))
			
			bc.AddFunction("main", bc.CurrentPosition())
			bc.AddInstruction(bytecode.OpPush, int64(str))
			bc.AddInstruction(bytecode.OpStrToUpper, 0)
			bc.AddInstruction(bytecode.OpReturnValue, 0)
			bc.AddInstruction(bytecode.OpHalt, 0)
			
			vm := NewVM(bc)
			result, err := vm.Run()
			
			if err != nil {
				t.Fatalf("VM execution failed: %v", err)
			}
			
			if result.Type != bytecode.ValueTypeString {
				t.Errorf("Expected string type, got %v", result.Type)
			}
			
			if result.Str != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result.Str)
			}
		})
	}
}

// Test 9: String to_lower method
func TestStringToLower(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"All uppercase", "HELLO", "hello"},
		{"Mixed case", "HeLLo", "hello"},
		{"Already lowercase", "hello", "hello"},
		{"With spaces", "HELLO WORLD", "hello world"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := bytecode.NewBytecode()
			str := bc.AddConstant(bytecode.NewStringValue(tt.input))
			
			bc.AddFunction("main", bc.CurrentPosition())
			bc.AddInstruction(bytecode.OpPush, int64(str))
			bc.AddInstruction(bytecode.OpStrToLower, 0)
			bc.AddInstruction(bytecode.OpReturnValue, 0)
			bc.AddInstruction(bytecode.OpHalt, 0)
			
			vm := NewVM(bc)
			result, err := vm.Run()
			
			if err != nil {
				t.Fatalf("VM execution failed: %v", err)
			}
			
			if result.Type != bytecode.ValueTypeString {
				t.Errorf("Expected string type, got %v", result.Type)
			}
			
			if result.Str != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result.Str)
			}
		})
	}
}

// Test 10: String trim method
func TestStringTrim(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Leading spaces", "  hello", "hello"},
		{"Trailing spaces", "hello  ", "hello"},
		{"Both sides", "  hello  ", "hello"},
		{"No spaces", "hello", "hello"},
		{"Only spaces", "   ", ""},
		{"Tabs and newlines", "\t\nhello\n\t", "hello"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := bytecode.NewBytecode()
			str := bc.AddConstant(bytecode.NewStringValue(tt.input))
			
			bc.AddFunction("main", bc.CurrentPosition())
			bc.AddInstruction(bytecode.OpPush, int64(str))
			bc.AddInstruction(bytecode.OpStrTrim, 0)
			bc.AddInstruction(bytecode.OpReturnValue, 0)
			bc.AddInstruction(bytecode.OpHalt, 0)
			
			vm := NewVM(bc)
			result, err := vm.Run()
			
			if err != nil {
				t.Fatalf("VM execution failed: %v", err)
			}
			
			if result.Type != bytecode.ValueTypeString {
				t.Errorf("Expected string type, got %v", result.Type)
			}
			
			if result.Str != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result.Str)
			}
		})
	}
}
