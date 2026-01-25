package vm

import (
	"testing"
	"unified-compiler/internal/bytecode"
)

// TestVMEnumAllocation tests enum variant allocation
func TestVMEnumAllocation(t *testing.T) {
	tests := []struct {
		name        string
		buildCode   func(*bytecode.Bytecode)
		wantType    bytecode.ValueType
		wantEnum    string
		wantVariant string
		wantTag     int
		wantDataLen int
	}{
		{
			name: "Simple enum no data",
			buildCode: func(bc *bytecode.Bytecode) {
				// Enum Color::Red (no data)
				enumNameIdx := bc.AddConstant(bytecode.NewStringValue("Color"))
				variantNameIdx := bc.AddConstant(bytecode.NewStringValue("Red"))
				tagIdx := bc.AddConstant(bytecode.NewIntValue(0))
				
				bc.AddInstruction(bytecode.OpPush, int64(enumNameIdx))
				bc.AddInstruction(bytecode.OpPush, int64(variantNameIdx))
				bc.AddInstruction(bytecode.OpPush, int64(tagIdx))
				bc.AddInstruction(bytecode.OpAllocEnum, 0) // 0 data elements
			},
			wantType:    bytecode.ValueTypeEnum,
			wantEnum:    "Color",
			wantVariant: "Red",
			wantTag:     0,
			wantDataLen: 0,
		},
		{
			name: "Enum with single data",
			buildCode: func(bc *bytecode.Bytecode) {
				// Enum Option::Some(42)
				dataIdx := bc.AddConstant(bytecode.NewIntValue(42))
				enumNameIdx := bc.AddConstant(bytecode.NewStringValue("Option"))
				variantNameIdx := bc.AddConstant(bytecode.NewStringValue("Some"))
				tagIdx := bc.AddConstant(bytecode.NewIntValue(1))
				
				bc.AddInstruction(bytecode.OpPush, int64(dataIdx))       // data
				bc.AddInstruction(bytecode.OpPush, int64(enumNameIdx))
				bc.AddInstruction(bytecode.OpPush, int64(variantNameIdx))
				bc.AddInstruction(bytecode.OpPush, int64(tagIdx))
				bc.AddInstruction(bytecode.OpAllocEnum, 1) // 1 data element
			},
			wantType:    bytecode.ValueTypeEnum,
			wantEnum:    "Option",
			wantVariant: "Some",
			wantTag:     1,
			wantDataLen: 1,
		},
		{
			name: "Enum with multiple data",
			buildCode: func(bc *bytecode.Bytecode) {
				// Enum Result::Ok(100, "success")
				data1Idx := bc.AddConstant(bytecode.NewIntValue(100))
				data2Idx := bc.AddConstant(bytecode.NewStringValue("success"))
				enumNameIdx := bc.AddConstant(bytecode.NewStringValue("Result"))
				variantNameIdx := bc.AddConstant(bytecode.NewStringValue("Ok"))
				tagIdx := bc.AddConstant(bytecode.NewIntValue(0))
				
				bc.AddInstruction(bytecode.OpPush, int64(data1Idx))
				bc.AddInstruction(bytecode.OpPush, int64(data2Idx))
				bc.AddInstruction(bytecode.OpPush, int64(enumNameIdx))
				bc.AddInstruction(bytecode.OpPush, int64(variantNameIdx))
				bc.AddInstruction(bytecode.OpPush, int64(tagIdx))
				bc.AddInstruction(bytecode.OpAllocEnum, 2) // 2 data elements
			},
			wantType:    bytecode.ValueTypeEnum,
			wantEnum:    "Result",
			wantVariant: "Ok",
			wantTag:     0,
			wantDataLen: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := bytecode.NewBytecode()
			bc.AddFunction("main", 0)
			tt.buildCode(bc)
			bc.AddInstruction(bytecode.OpReturnValue, 0)
			bc.AddInstruction(bytecode.OpHalt, 0)

			vm := NewVM(bc)
			result, err := vm.Run()
			if err != nil {
				t.Fatalf("VM execution failed: %v", err)
			}

			// result is returned from vm.Run()
			if result.Type != tt.wantType {
				t.Errorf("Got type %v, want %v", result.Type, tt.wantType)
			}
			if result.Enum == nil {
				t.Fatal("Expected enum value, got nil")
			}
			if result.Enum.EnumName != tt.wantEnum {
				t.Errorf("Got enum name %v, want %v", result.Enum.EnumName, tt.wantEnum)
			}
			if result.Enum.VariantName != tt.wantVariant {
				t.Errorf("Got variant name %v, want %v", result.Enum.VariantName, tt.wantVariant)
			}
			if result.Enum.VariantTag != tt.wantTag {
				t.Errorf("Got tag %v, want %v", result.Enum.VariantTag, tt.wantTag)
			}
			if len(result.Enum.Data) != tt.wantDataLen {
				t.Errorf("Got data length %v, want %v", len(result.Enum.Data), tt.wantDataLen)
			}
		})
	}
}

// TestVMMatchVariant tests enum variant matching
func TestVMMatchVariant(t *testing.T) {
	tests := []struct {
		name      string
		buildCode func(*bytecode.Bytecode)
		wantMatch bool
	}{
		{
			name: "Matching variant",
			buildCode: func(bc *bytecode.Bytecode) {
				// Create Color::Red
				enumNameIdx := bc.AddConstant(bytecode.NewStringValue("Color"))
				variantNameIdx := bc.AddConstant(bytecode.NewStringValue("Red"))
				tagIdx := bc.AddConstant(bytecode.NewIntValue(0))
				
				bc.AddInstruction(bytecode.OpPush, int64(enumNameIdx))
				bc.AddInstruction(bytecode.OpPush, int64(variantNameIdx))
				bc.AddInstruction(bytecode.OpPush, int64(tagIdx))
				bc.AddInstruction(bytecode.OpAllocEnum, 0)
				
				// Match against "Red"
				matchNameIdx := bc.AddConstant(bytecode.NewStringValue("Red"))
				bc.AddInstruction(bytecode.OpPush, int64(matchNameIdx))
				bc.AddInstruction(bytecode.OpMatchVariant, 0)
			},
			wantMatch: true,
		},
		{
			name: "Non-matching variant",
			buildCode: func(bc *bytecode.Bytecode) {
				// Create Color::Red
				enumNameIdx := bc.AddConstant(bytecode.NewStringValue("Color"))
				variantNameIdx := bc.AddConstant(bytecode.NewStringValue("Red"))
				tagIdx := bc.AddConstant(bytecode.NewIntValue(0))
				
				bc.AddInstruction(bytecode.OpPush, int64(enumNameIdx))
				bc.AddInstruction(bytecode.OpPush, int64(variantNameIdx))
				bc.AddInstruction(bytecode.OpPush, int64(tagIdx))
				bc.AddInstruction(bytecode.OpAllocEnum, 0)
				
				// Match against "Blue"
				matchNameIdx := bc.AddConstant(bytecode.NewStringValue("Blue"))
				bc.AddInstruction(bytecode.OpPush, int64(matchNameIdx))
				bc.AddInstruction(bytecode.OpMatchVariant, 0)
			},
			wantMatch: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := bytecode.NewBytecode()
			bc.AddFunction("main", 0)
			tt.buildCode(bc)
			bc.AddInstruction(bytecode.OpReturnValue, 0)
			bc.AddInstruction(bytecode.OpHalt, 0)

			vm := NewVM(bc)
			result, err := vm.Run()
			if err != nil {
				t.Fatalf("VM execution failed: %v", err)
			}

			// result is returned from vm.Run()
			if result.Type != bytecode.ValueTypeBool {
				t.Errorf("Got type %v, want Bool", result.Type)
			}
			if result.Bool != tt.wantMatch {
				t.Errorf("Got match result %v, want %v", result.Bool, tt.wantMatch)
			}
		})
	}
}

// TestVMExtractVariant tests extracting data from enum variants
func TestVMExtractVariant(t *testing.T) {
	tests := []struct {
		name      string
		buildCode func(*bytecode.Bytecode)
		wantValue bytecode.Value
	}{
		{
			name: "Extract first element",
			buildCode: func(bc *bytecode.Bytecode) {
				// Create Option::Some(42)
				dataIdx := bc.AddConstant(bytecode.NewIntValue(42))
				enumNameIdx := bc.AddConstant(bytecode.NewStringValue("Option"))
				variantNameIdx := bc.AddConstant(bytecode.NewStringValue("Some"))
				tagIdx := bc.AddConstant(bytecode.NewIntValue(1))
				
				bc.AddInstruction(bytecode.OpPush, int64(dataIdx))
				bc.AddInstruction(bytecode.OpPush, int64(enumNameIdx))
				bc.AddInstruction(bytecode.OpPush, int64(variantNameIdx))
				bc.AddInstruction(bytecode.OpPush, int64(tagIdx))
				bc.AddInstruction(bytecode.OpAllocEnum, 1)
				
				// Extract data at index 0
				indexIdx := bc.AddConstant(bytecode.NewIntValue(0))
				bc.AddInstruction(bytecode.OpPush, int64(indexIdx))
				bc.AddInstruction(bytecode.OpExtractVariant, 0)
			},
			wantValue: bytecode.NewIntValue(42),
		},
		{
			name: "Extract second element",
			buildCode: func(bc *bytecode.Bytecode) {
				// Create Result::Ok(100, "success")
				data1Idx := bc.AddConstant(bytecode.NewIntValue(100))
				data2Idx := bc.AddConstant(bytecode.NewStringValue("success"))
				enumNameIdx := bc.AddConstant(bytecode.NewStringValue("Result"))
				variantNameIdx := bc.AddConstant(bytecode.NewStringValue("Ok"))
				tagIdx := bc.AddConstant(bytecode.NewIntValue(0))
				
				bc.AddInstruction(bytecode.OpPush, int64(data1Idx))
				bc.AddInstruction(bytecode.OpPush, int64(data2Idx))
				bc.AddInstruction(bytecode.OpPush, int64(enumNameIdx))
				bc.AddInstruction(bytecode.OpPush, int64(variantNameIdx))
				bc.AddInstruction(bytecode.OpPush, int64(tagIdx))
				bc.AddInstruction(bytecode.OpAllocEnum, 2)
				
				// Extract data at index 1
				indexIdx := bc.AddConstant(bytecode.NewIntValue(1))
				bc.AddInstruction(bytecode.OpPush, int64(indexIdx))
				bc.AddInstruction(bytecode.OpExtractVariant, 0)
			},
			wantValue: bytecode.NewStringValue("success"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := bytecode.NewBytecode()
			bc.AddFunction("main", 0)
			tt.buildCode(bc)
			bc.AddInstruction(bytecode.OpReturnValue, 0)
			bc.AddInstruction(bytecode.OpHalt, 0)

			vm := NewVM(bc)
			result, err := vm.Run()
			if err != nil {
				t.Fatalf("VM execution failed: %v", err)
			}

			// result is returned from vm.Run()
			if result.Type != tt.wantValue.Type {
				t.Errorf("Got type %v, want %v", result.Type, tt.wantValue.Type)
			}
			
			switch tt.wantValue.Type {
			case bytecode.ValueTypeInt:
				if result.Int != tt.wantValue.Int {
					t.Errorf("Got value %v, want %v", result.Int, tt.wantValue.Int)
				}
			case bytecode.ValueTypeString:
				if result.Str != tt.wantValue.Str {
					t.Errorf("Got value %v, want %v", result.Str, tt.wantValue.Str)
				}
			}
		})
	}
}
