package wasm

import (
	"bytes"
)

// Encode encodes a WASM module to binary format
func (m *Module) Encode() ([]byte, error) {
	var buf bytes.Buffer

	// Magic number
	buf.Write([]byte{0x00, 0x61, 0x73, 0x6D}) // \0asm

	// Version
	buf.Write([]byte{0x01, 0x00, 0x00, 0x00}) // version 1

	// Type section
	if err := m.encodeTypeSection(&buf); err != nil {
		return nil, err
	}

	// Function section
	if err := m.encodeFunctionSection(&buf); err != nil {
		return nil, err
	}

	// Global section
	if err := m.encodeGlobalSection(&buf); err != nil {
		return nil, err
	}

	// Memory section
	if err := m.encodeMemorySection(&buf); err != nil {
		return nil, err
	}

	// Export section
	if err := m.encodeExportSection(&buf); err != nil {
		return nil, err
	}

	// Code section
	if err := m.encodeCodeSection(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// encodeTypeSection encodes the type section
func (m *Module) encodeTypeSection(buf *bytes.Buffer) error {
	if len(m.FunctionTypes) == 0 {
		return nil
	}

	var section bytes.Buffer

	// Number of types
	writeULEB128(&section, uint64(len(m.FunctionTypes)))

	// Encode each type
	for _, ft := range m.FunctionTypes {
		section.WriteByte(0x60) // func type

		// Parameters
		writeULEB128(&section, uint64(len(ft.Params)))
		for _, p := range ft.Params {
			section.WriteByte(byte(p))
		}

		// Returns
		writeULEB128(&section, uint64(len(ft.Returns)))
		for _, r := range ft.Returns {
			section.WriteByte(byte(r))
		}
	}

	// Write section
	buf.WriteByte(0x01) // type section ID
	writeULEB128(buf, uint64(section.Len()))
	buf.Write(section.Bytes())

	return nil
}

// encodeFunctionSection encodes the function section
func (m *Module) encodeFunctionSection(buf *bytes.Buffer) error {
	if len(m.Functions) == 0 {
		return nil
	}

	var section bytes.Buffer

	// Number of functions
	writeULEB128(&section, uint64(len(m.Functions)))

	// Encode each function's type index
	for _, fn := range m.Functions {
		writeULEB128(&section, uint64(fn.TypeIndex))
	}

	// Write section
	buf.WriteByte(0x03) // function section ID
	writeULEB128(buf, uint64(section.Len()))
	buf.Write(section.Bytes())

	return nil
}

// encodeGlobalSection encodes the global section
func (m *Module) encodeGlobalSection(buf *bytes.Buffer) error {
	if len(m.Globals) == 0 {
		return nil
	}

	var section bytes.Buffer

	// Number of globals
	writeULEB128(&section, uint64(len(m.Globals)))

	// Encode each global
	for _, global := range m.Globals {
		// Global type
		section.WriteByte(byte(global.Type))
		
		// Mutability
		if global.Mutable {
			section.WriteByte(0x01)
		} else {
			section.WriteByte(0x00)
		}
		
		// Initializer expression
		section.Write(global.Init)
	}

	// Write section
	buf.WriteByte(0x06) // global section ID
	writeULEB128(buf, uint64(section.Len()))
	buf.Write(section.Bytes())

	return nil
}

// encodeMemorySection encodes the memory section
func (m *Module) encodeMemorySection(buf *bytes.Buffer) error {
	if m.Memory == nil {
		return nil
	}

	var section bytes.Buffer

	// Number of memories (always 1)
	writeULEB128(&section, 1)

	// Memory limits
	if m.Memory.Max > 0 {
		section.WriteByte(0x01) // has max
		writeULEB128(&section, uint64(m.Memory.Min))
		writeULEB128(&section, uint64(m.Memory.Max))
	} else {
		section.WriteByte(0x00) // no max
		writeULEB128(&section, uint64(m.Memory.Min))
	}

	// Write section
	buf.WriteByte(0x05) // memory section ID
	writeULEB128(buf, uint64(section.Len()))
	buf.Write(section.Bytes())

	return nil
}

// encodeExportSection encodes the export section
func (m *Module) encodeExportSection(buf *bytes.Buffer) error {
	if len(m.Exports) == 0 {
		return nil
	}

	var section bytes.Buffer

	// Number of exports
	writeULEB128(&section, uint64(len(m.Exports)))

	// Encode each export
	for _, exp := range m.Exports {
		// Name
		writeULEB128(&section, uint64(len(exp.Name)))
		section.WriteString(exp.Name)

		// Kind
		section.WriteByte(exp.Kind)

		// Index
		writeULEB128(&section, uint64(exp.Index))
	}

	// Write section
	buf.WriteByte(0x07) // export section ID
	writeULEB128(buf, uint64(section.Len()))
	buf.Write(section.Bytes())

	return nil
}

// encodeCodeSection encodes the code section
func (m *Module) encodeCodeSection(buf *bytes.Buffer) error {
	if len(m.Functions) == 0 {
		return nil
	}

	var section bytes.Buffer

	// Number of function bodies
	writeULEB128(&section, uint64(len(m.Functions)))

	// Encode each function body
	for _, fn := range m.Functions {
		var fnBody bytes.Buffer

		// Locals
		writeULEB128(&fnBody, uint64(len(fn.Locals)))
		for _, local := range fn.Locals {
			writeULEB128(&fnBody, uint64(local.Count))
			fnBody.WriteByte(byte(local.Type))
		}

		// Body code
		fnBody.Write(fn.Body)

		// Write function body with size
		writeULEB128(&section, uint64(fnBody.Len()))
		section.Write(fnBody.Bytes())
	}

	// Write section
	buf.WriteByte(0x0A) // code section ID
	writeULEB128(buf, uint64(section.Len()))
	buf.Write(section.Bytes())

	return nil
}

// writeULEB128 writes an unsigned LEB128 encoded integer
func writeULEB128(buf *bytes.Buffer, value uint64) {
	for {
		b := byte(value & 0x7F)
		value >>= 7
		if value == 0 {
			buf.WriteByte(b)
			break
		}
		buf.WriteByte(b | 0x80)
	}
}

// writeLEB128 writes a signed LEB128 encoded integer
func writeLEB128(buf *bytes.Buffer, value int64) {
	for {
		b := byte(value & 0x7F)
		value >>= 7
		if (value == 0 && (b&0x40) == 0) || (value == -1 && (b&0x40) != 0) {
			buf.WriteByte(b)
			break
		}
		buf.WriteByte(b | 0x80)
	}
}

// writeBytes writes a byte slice to the buffer with length prefix
func writeBytes(buf *bytes.Buffer, data []byte) {
	writeULEB128(buf, uint64(len(data)))
	buf.Write(data)
}
