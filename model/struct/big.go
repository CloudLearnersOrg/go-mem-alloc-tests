package model

import (
	"net"
	"time"
)

// LargeOptimizedStruct represents a struct with fields arranged to minimize memory padding
type LargeOptimizedStruct struct {
	// Complex fields with pointers and internal structures
	UserProfile     map[string]string    // Map is a pointer
	ActivityHistory []time.Time          // Slice is a pointer + len + cap
	Settings        map[string]bool      // Another map
	Tags            []string             // Slice of strings
	IPAddress       net.IP               // IP address (slice under the hood)
	Listener        net.Listener         // Interface (pointer + type data)
	Conn            net.Conn             // Another interface
	Metrics         map[string]float64   // Map of metrics
	
	// 8-byte aligned fields
	CreatedAt         time.Time          // 24 bytes typically
	UpdatedAt         time.Time
	ExpiresAt         time.Time
	DueDate           time.Time
	Duration          time.Duration      // 8 bytes (int64)
	Timeout           time.Duration
	TransactionID     uint64
	UserID            uint64
	AccountID         uint64
	OrderID           uint64
	ParentID          uint64
	RequestTimestamp  int64
	ResponseTimestamp int64
	Balance           float64
	Credit            float64
	Score             float64

	// 4-byte fields
	StatusCode       int32
	ResponseCode     int32
	RequestCount     int32
	RetryAttempts    int32
	ErrorCount       int32
	BatchSize        int32
	ServiceTime      float32
	CPUTime          float32
	MemoryUsage      float32
	DiskUsage        float32
	NetworkUsage     float32
	Percentage       float32

	// 2-byte fields
	ErrorCode        uint16
	ProtocolVersion  uint16
	ServerRegion     uint16
	ClientRegion     uint16
	Port             uint16
	BackupPort       uint16

	// 1-byte fields
	IsSuccess        bool
	IsRetry          bool
	IsCached         bool
	IsVerified       bool
	IsActive         bool
	IsAdmin          bool
	IsTest           bool
	IsPriority       bool
	IsFlagged        bool
	Priority         uint8
	CompressionLevel uint8
	Importance       uint8
	Status           byte
	Type             byte
}

// LargeUnoptimizedStruct represents a struct with fields arranged in a suboptimal way
type LargeUnoptimizedStruct struct {
	// 1-byte fields first (causing potential padding)
	IsSuccess bool
	Status    byte
	Priority  uint8

	// 8-byte field surrounded by 1-byte fields (padding)
	UserID uint64

	// 1-byte fields
	IsRetry  bool
	IsActive bool

	// 4-byte fields
	StatusCode  int32
	ServiceTime float32

	// 2-byte field between 4-byte and pointer fields (padding)
	ErrorCode uint16

	// Complex types mixed with primitives
	CreatedAt     time.Time
	IsAdmin       bool
	TransactionID uint64
	IsCached      bool

	// 4-byte field
	ResponseCode int32

	// 2-byte fields
	ProtocolVersion uint16
	ServerRegion    uint16

	// Map field with 1-byte field after
	UserProfile map[string]string
	IsTest      bool

	// More 8-byte fields
	UpdatedAt         time.Time
	OrderID           uint64
	RequestTimestamp  int64
	ResponseTimestamp int64
	Balance           float64

	// Slice surrounded by 1-byte fields
	IsPriority bool
	Tags       []string
	IsFlagged  bool

	// Remaining fields in random order
	Type             byte
	RetryAttempts    int32
	Port             uint16
	CPUTime          float32
	ExpiresAt        time.Time
	ActivityHistory  []time.Time
	CompressionLevel uint8
	ErrorCount       int32
	Importance       uint8
	ClientRegion     uint16
	Credit           float64
	Score            float64
	IPAddress        net.IP
	DueDate          time.Time
	Timeout          time.Duration
	BatchSize        int32
	Listener         net.Listener
	NetworkUsage     float32
	DiskUsage        float32
	Conn             net.Conn
	AccountID        uint64
	ParentID         uint64
	Duration         time.Duration
	MemoryUsage      float32
	BackupPort       uint16
	RequestCount     int32
	IsVerified       bool
	Settings         map[string]bool
	Metrics          map[string]float64
	Percentage       float32
}
