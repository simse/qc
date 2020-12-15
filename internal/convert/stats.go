package convert

import "github.com/simse/qc/internal/format"

/*
// Error represents a conversion error
type Error struct {
	Reason string
}

// Warning represents a conversion warning
type Warning struct {
}

// Success represents a conversion that completed
type Success struct {
	Key string
}

// Ignored represents a conversion that was skipped
type Ignored struct {
	Reason string
}

// Stats contains statistics about a convert jobs
type Stats struct {
	Errors    []Error
	Warnings  []Warning
	Successes []Success
	Ignored   []Ignored
}
*/

// ConversionResult stores information about a conversion like status, speed and so on
type ConversionResult struct {
	InputPath    string
	OutputPath   string
	InputFormat  format.Info
	OutputFormat format.Info
	Duration     int64
	Error        string
	Warning      string
	Errored      bool
	Warned       bool
	Success      bool
	Skipped      bool
	SkipReason   string
	InputSize    int64
	OutputSize   int64
}
