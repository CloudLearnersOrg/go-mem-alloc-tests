package model

import "time"

// API struct definitions

// APIOptimizedStruct represents an optimized API request struct
type APIOptimizedStruct struct {
	// 8-byte fields first
	RequestID uint64
	UserID    uint64
	Timestamp int64
	SessionID uint64
	// 4-byte fields
	StatusCode int32
	Latency    float32
	// 2-byte fields
	APIVersion uint16
	// 1-byte fields
	Method        byte
	Authenticated bool
	Cached        bool
}

// APIUnoptimizedStruct represents an unoptimized API request struct
type APIUnoptimizedStruct struct {
	// 1-byte fields first (causing padding)
	Method        byte
	Authenticated bool
	// 8-byte field
	UserID uint64
	// 1-byte field
	Cached bool
	// 4-byte field
	StatusCode int32
	// 2-byte field
	APIVersion uint16
	// 8-byte fields
	RequestID uint64
	Timestamp int64
	SessionID uint64
	// 4-byte field
	Latency float32
}

// Config struct definitions

// ConfigOptimizedStruct represents an optimized configuration struct
type ConfigOptimizedStruct struct {
	// String fields (contain pointers, which are 8 bytes)
	Name        string
	Description string
	Environment string
	// 8-byte fields
	UpdatedAt int64
	CreatedAt int64
	// 4-byte fields
	MaxConnections int32
	Timeout        int32
	// 2-byte fields
	Port uint16
	// 1-byte fields
	Debug   bool
	Enabled bool
}

// ConfigUnoptimizedStruct represents an unoptimized configuration struct
type ConfigUnoptimizedStruct struct {
	// 1-byte fields first (causing padding)
	Debug   bool
	Enabled bool
	// 2-byte field
	Port uint16
	// String field (contains pointer)
	Name string
	// 4-byte field
	Timeout int32
	// String field
	Environment string
	// 4-byte field
	MaxConnections int32
	// 8-byte field
	CreatedAt int64
	// String field
	Description string
	// 8-byte field
	UpdatedAt int64
}

// GraphQL struct definitions

// GraphQLOptimizedStruct represents an optimized GraphQL query struct
type GraphQLOptimizedStruct struct {
	// String pointers (8 bytes)
	QueryID   string
	Operation string
	ClientID  string
	// 8-byte fields
	Timestamp int64
	Duration  int64
	// 4-byte fields
	Depth           int32
	ComplexityScore float32
	// 2-byte fields
	FragmentCount uint16
	// 1-byte fields
	IsMutation   bool
	HasVariables bool
	Cached       bool
}

// GraphQLUnoptimizedStruct represents an unoptimized GraphQL query struct
type GraphQLUnoptimizedStruct struct {
	// 1-byte fields first (causing padding)
	IsMutation bool
	Cached     bool
	// 4-byte field
	Depth int32
	// String field
	Operation string
	// 1-byte field
	HasVariables bool
	// 2-byte field
	FragmentCount uint16
	// 8-byte field
	Timestamp int64
	// String field
	QueryID string
	// 4-byte field
	ComplexityScore float32
	// String field
	ClientID string
	// 8-byte field
	Duration int64
}

// DB entity struct definitions

// DBEntityOptimizedStruct represents an optimized database entity struct
type DBEntityOptimizedStruct struct {
	// String fields (pointers, 8 bytes)
	ID    string
	Name  string
	Email string
	// 8-byte fields
	CreatedAt   time.Time // Time is larger than 8 bytes
	UpdatedAt   time.Time
	LastLoginAt time.Time
	// 4-byte fields
	LoginCount int32
	Status     int32
	// 2-byte fields
	AccessLevel uint16
	// 1-byte fields
	IsActive bool
	IsAdmin  bool
	HasMFA   bool
}

// DBEntityUnoptimizedStruct represents an unoptimized database entity struct
type DBEntityUnoptimizedStruct struct {
	// 1-byte fields first (causing padding)
	IsActive bool
	IsAdmin  bool
	// 2-byte field
	AccessLevel uint16
	// String field
	Email string
	// 4-byte field
	Status int32
	// Time fields (large)
	LastLoginAt time.Time
	// 1-byte field
	HasMFA bool
	// String fields
	ID   string
	Name string
	// 4-byte field
	LoginCount int32
	// Time fields
	CreatedAt time.Time
	UpdatedAt time.Time
}
