package scoring

import (
	"testing"

	"github.com/ismailtsdln/socialrecon/internal/models"
)

func TestScoringEngine_Calculate(t *testing.T) {
	scorer := NewScoringEngine()

	tests := []struct {
		name     string
		findings []models.Finding
		expected float64
	}{
		{
			name:     "No findings",
			findings: []models.Finding{},
			expected: 0.0,
		},
		{
			name: "Single existing profile",
			findings: []models.Finding{
				{Indicator: "github_profile", Status: "exists", Severity: models.SeverityInfo},
			},
			expected: 2.0, // weight 10 * 0.2
		},
		{
			name: "Available profile (hijack risk)",
			findings: []models.Finding{
				{Indicator: "twitter_profile", Status: "available", Severity: models.SeverityHigh},
			},
			expected: 30.0, // weight 15 * 2.0
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := &models.ScanResult{
				Findings: tt.findings,
			}
			score := scorer.Calculate(result)
			if score != tt.expected {
				t.Errorf("Calculate() = %v, want %v", score, tt.expected)
			}
		})
	}
}

func TestScoringEngine_GetOverallSeverity(t *testing.T) {
	scorer := NewScoringEngine()

	result := &models.ScanResult{
		Findings: []models.Finding{
			{Severity: models.SeverityInfo},
			{Severity: models.SeverityHigh},
			{Severity: models.SeverityLow},
		},
	}

	if got := scorer.GetOverallSeverity(result); got != models.SeverityHigh {
		t.Errorf("GetOverallSeverity() = %v, want %v", got, models.SeverityHigh)
	}
}
