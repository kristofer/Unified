package wasm

import (
	"bytes"
)

// MemoryAllocator manages WASM linear memory allocation
type MemoryAllocator struct {
	currentOffset uint32
	dataSegments  []DataSegment
}

// NewMemoryAllocator creates a new memory allocator
func NewMemoryAllocator() *MemoryAllocator {
	return &MemoryAllocator{
		currentOffset: 1024, // Start allocations after first 1KB (reserved for heap pointer, etc.)
		dataSegments:  make([]DataSegment, 0),
	}
}

// Allocate reserves a block of memory and returns its offset
func (ma *MemoryAllocator) Allocate(size uint32) uint32 {
	offset := ma.currentOffset
	ma.currentOffset += size
	// Align to 8-byte boundary
	if ma.currentOffset%8 != 0 {
		ma.currentOffset += 8 - (ma.currentOffset % 8)
	}
	return offset
}

// AllocateWithData reserves memory and initializes it with data
func (ma *MemoryAllocator) AllocateWithData(data []byte) uint32 {
	offset := ma.currentOffset
	ma.currentOffset += uint32(len(data))
	
	// Add data segment
	ma.dataSegments = append(ma.dataSegments, DataSegment{
		Offset: offset,
		Data:   data,
	})
	
	// Align to 8-byte boundary
	if ma.currentOffset%8 != 0 {
		ma.currentOffset += 8 - (ma.currentOffset % 8)
	}
	
	return offset
}

// GetDataSegments returns all data segments for module initialization
func (ma *MemoryAllocator) GetDataSegments() []DataSegment {
	return ma.dataSegments
}

// GetCurrentOffset returns the current memory offset
func (ma *MemoryAllocator) GetCurrentOffset() uint32 {
	return ma.currentOffset
}

// emitHeapAlloc generates WASM code to allocate memory at runtime
// This uses a simple bump allocator stored in a global variable
func (g *Generator) emitHeapAlloc(body *bytes.Buffer, size uint32) {
	// Load current heap pointer (from global 0)
	body.WriteByte(0x23) // global.get
	body.WriteByte(0x00) // global index 0
	
	// Save it for return value
	body.WriteByte(0x22) // local.tee
	g.emitULEB128(body, uint64(g.localVarCount)) // Use a temp local
	
	// Add size to heap pointer
	body.WriteByte(0x41) // i32.const
	g.emitULEB128(body, uint64(size))
	body.WriteByte(0x6A) // i32.add
	
	// Store new heap pointer
	body.WriteByte(0x24) // global.set
	body.WriteByte(0x00) // global index 0
	
	// Original pointer is now on stack (from local.tee)
	body.WriteByte(0x20) // local.get
	g.emitULEB128(body, uint64(g.localVarCount))
}

// InitializeHeapPointer generates code to initialize the heap pointer global
func (g *Generator) InitializeHeapPointer() {
	// Add a global variable for the heap pointer
	// This will be initialized to the start of the heap (after static data)
	g.module.Globals = append(g.module.Globals, Global{
		Type:    I32,
		Mutable: true,
		Init:    []byte{0x41, 0x00, 0x04, 0x00, 0x0B}, // i32.const 1024, end
	})
}
