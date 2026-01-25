package bytecode

import (
	"testing"
)

// TestEnumValue tests creating and manipulating enum values
func TestEnumValue(t *testing.T) {
	tests := []struct {
		name        string
		enumName    string
		variantName string
		variantTag  int
		data        []Value
		wantType    ValueType
	}{
		{
			name:        "Simple enum variant",
			enumName:    "Color",
			variantName: "Red",
			variantTag:  0,
			data:        []Value{},
			wantType:    ValueTypeEnum,
		},
		{
			name:        "Enum with single data",
			enumName:    "Option",
			variantName: "Some",
			variantTag:  1,
			data:        []Value{NewIntValue(42)},
			wantType:    ValueTypeEnum,
		},
		{
			name:        "Enum with multiple data",
			enumName:    "Result",
			variantName: "Ok",
			variantTag:  0,
			data:        []Value{NewIntValue(100), NewStringValue("success")},
			wantType:    ValueTypeEnum,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val := NewEnumValue(tt.enumName, tt.variantName, tt.variantTag, tt.data)

			if val.Type != tt.wantType {
				t.Errorf("NewEnumValue() type = %v, want %v", val.Type, tt.wantType)
			}
			if val.Enum == nil {
				t.Fatal("NewEnumValue() produced nil Enum field")
			}
			if val.Enum.EnumName != tt.enumName {
				t.Errorf("NewEnumValue() enum name = %v, want %v", val.Enum.EnumName, tt.enumName)
			}
			if val.Enum.VariantName != tt.variantName {
				t.Errorf("NewEnumValue() variant name = %v, want %v", val.Enum.VariantName, tt.variantName)
			}
			if val.Enum.VariantTag != tt.variantTag {
				t.Errorf("NewEnumValue() variant tag = %v, want %v", val.Enum.VariantTag, tt.variantTag)
			}
			if len(val.Enum.Data) != len(tt.data) {
				t.Errorf("NewEnumValue() data length = %v, want %v", len(val.Enum.Data), len(tt.data))
			}
		})
	}
}

// TestEnumValueString tests the string representation of enum values
func TestEnumValueString(t *testing.T) {
	tests := []struct {
		name        string
		value       Value
		wantContain string
	}{
		{
			name:        "Simple enum variant string",
			value:       NewEnumValue("Color", "Red", 0, []Value{}),
			wantContain: "Color::Red",
		},
		{
			name:        "Enum with data string",
			value:       NewEnumValue("Option", "Some", 1, []Value{NewIntValue(42)}),
			wantContain: "Option::Some",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.value.String()
			if got != "enum:"+tt.wantContain {
				t.Errorf("Value.String() = %v, want to contain %v", got, "enum:"+tt.wantContain)
			}
		})
	}
}

// TestOpCodeEnumInstructions tests enum opcode string representations
func TestOpCodeEnumInstructions(t *testing.T) {
	tests := []struct {
		name   string
		opcode OpCode
		want   string
	}{
		{"AllocEnum", OpAllocEnum, "ALLOC_ENUM"},
		{"MatchVariant", OpMatchVariant, "MATCH_VARIANT"},
		{"ExtractVariant", OpExtractVariant, "EXTRACT_VARIANT"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.opcode.String(); got != tt.want {
				t.Errorf("OpCode.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestEnumTypeInfo tests enum type info structure
func TestEnumValueDataAccess(t *testing.T) {
	data := []Value{
		NewIntValue(42),
		NewStringValue("test"),
		NewBoolValue(true),
	}
	
	enumVal := NewEnumValue("TestEnum", "TestVariant", 0, data)
	
	if len(enumVal.Enum.Data) != 3 {
		t.Errorf("Expected 3 data elements, got %d", len(enumVal.Enum.Data))
	}
	
	if enumVal.Enum.Data[0].Int != 42 {
		t.Errorf("Expected first data element to be 42, got %d", enumVal.Enum.Data[0].Int)
	}
	
	if enumVal.Enum.Data[1].Str != "test" {
		t.Errorf("Expected second data element to be 'test', got %s", enumVal.Enum.Data[1].Str)
	}
	
	if enumVal.Enum.Data[2].Bool != true {
		t.Errorf("Expected third data element to be true, got %t", enumVal.Enum.Data[2].Bool)
	}
}
