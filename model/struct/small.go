package model

import (
	"time"
)

// OptimizedStruct arranges fields from largest to smallest data types
type OptimizedStruct struct {
	// String fields (pointers - 8 bytes or 16 bytes on 64-bit systems)
	StringField  string
	StringFieldB string

	// Slices (pointers + length + capacity - 24 bytes on 64-bit systems)
	SliceField []int
	MapField   map[string]int

	// Complex data types
	DurationField time.Duration // 8 bytes (int64)
	TimeField     time.Time     // 24 bytes typically

	// 8-byte aligned fields
	Float64Field float64 // 8 bytes
	Int64Field   int64   // 8 bytes
	Uint64Field  uint64  // 8 bytes
	Int64FieldB  int64   // 8 bytes

	// Interface (16 bytes typically)
	InterfaceField interface{}

	// 4-byte aligned fields
	Float32Field float32 // 4 bytes
	Int32Field   int32   // 4 bytes
	Uint32Field  uint32  // 4 bytes
	Int32FieldB  int32   // 4 bytes

	// 2-byte aligned fields
	Int16Field  int16  // 2 bytes
	Uint16Field uint16 // 2 bytes
	Int16FieldB int16  // 2 bytes

	// 1-byte fields at the end
	Int8Field  int8  // 1 byte
	Uint8Field uint8 // 1 byte
	BoolField  bool  // 1 byte
	BoolFieldB bool  // 1 byte
	ByteField  byte  // 1 byte
	RuneField  rune  // technically 4 bytes (alias for int32)
}

// UnoptimizedStruct arranges fields in a suboptimal order
type UnoptimizedStruct struct {
	// 1-byte fields first (causing padding)
	BoolField  bool  // 1 byte (will be padded)
	Int8Field  int8  // 1 byte
	Uint8Field uint8 // 1 byte
	ByteField  byte  // 1 byte

	// 4-byte field
	RuneField rune // 4 bytes

	// 2-byte fields not aligned properly
	Int16Field int16 // 2 bytes (might get padded)

	// 8-byte field
	Int64Field int64 // 8 bytes

	// 1-byte field in the middle (causing padding)
	BoolFieldB bool // 1 byte (will be padded)

	// 4-byte fields
	Int32Field   int32   // 4 bytes
	Float32Field float32 // 4 bytes

	// String (pointer)
	StringField string // pointer (will cause padding)

	// 2-byte fields
	Uint16Field uint16 // 2 bytes
	Int16FieldB int16  // 2 bytes

	// Complex types in suboptimal positions
	TimeField     time.Time     // 24 bytes
	DurationField time.Duration // 8 bytes

	// More mixing to cause padding
	Int32FieldB  int32   // 4 bytes
	Uint32Field  uint32  // 4 bytes
	Float64Field float64 // 8 bytes
	Uint64Field  uint64  // 8 bytes

	// Interface type
	InterfaceField interface{} // 16 bytes typically

	// Slices and maps at the end
	SliceField   []int          // pointer + len + cap
	MapField     map[string]int // pointer + ...
	StringFieldB string         // pointer
	Int64FieldB  int64          // 8 bytes
}
