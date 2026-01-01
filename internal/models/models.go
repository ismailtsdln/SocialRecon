package models

import "time"

// Severity defines the risk level of a finding
type Severity string

const (
	SeverityInfo     Severity = "INFO"
	SeverityLow      Severity = "LOW"
	SeverityMedium   Severity = "MEDIUM"
	SeverityHigh     Severity = "HIGH"
	SeverityCritical Severity = "CRITICAL"
)

// Finding represents a single discovery by a plugin
type Finding struct {
	PluginName  string    `json:"plugin_name"`
	Indicator   string    `json:"indicator"`   // e.g., "twitter.com/user"
	Value       string    `json:"value"`       // e.g., "johndoe"
	Status      string    `json:"status"`      // e.g., "exists", "available", "suspended"
	Severity    Severity  `json:"severity"`
	Description string    `json:"description"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Timestamp   time.Time `json:"timestamp"`
}

// ScanResult is the final output of a scan
type ScanResult struct {
	Target    string    `json:"target"`
	Findings  []Finding `json:"findings"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	RiskScore float64   `json:"risk_score"`
}

// Config holds engine configuration
type Config struct {
	MaxConcurrency int
	Timeout        time.Duration
	RateLimit      int // requests per second
}
