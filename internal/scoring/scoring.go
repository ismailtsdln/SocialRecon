package scoring

import (
	"github.com/ismailtsdln/socialrecon/internal/models"
)

// ScoringEngine calculates risk scores based on findings
type ScoringEngine struct {
	weights map[string]float64
}

func NewScoringEngine() *ScoringEngine {
	return &ScoringEngine{
		weights: map[string]float64{
			"github_profile":    10.0,
			"twitter_profile":   15.0,
			"instagram_profile": 12.0,
		},
	}
}

// Calculate assigns a cumulative risk score to the result
func (e *ScoringEngine) Calculate(result *models.ScanResult) float64 {
	var totalScore float64

	for _, finding := range result.Findings {
		weight, ok := e.weights[finding.Indicator]
		if !ok {
			weight = 5.0 // default weight for unknown indicators
		}

		// Adjust weight based on status
		switch finding.Status {
		case "available":
			// Hijack risk is higher than existence
			totalScore += weight * 2.0
			finding.Severity = models.SeverityHigh
		case "suspended":
			totalScore += weight * 0.5
			finding.Severity = models.SeverityMedium
		case "exists":
			totalScore += weight * 0.2
			finding.Severity = models.SeverityInfo
		}
	}

	// Normalize score to 0-100 (cap at 100)
	if totalScore > 100 {
		totalScore = 100
	}

	return totalScore
}

// GetOverallSeverity returns the highest severity level found
func (e *ScoringEngine) GetOverallSeverity(result *models.ScanResult) models.Severity {
	maxSeverity := models.SeverityInfo

	severityMap := map[models.Severity]int{
		models.SeverityInfo:     0,
		models.SeverityLow:      1,
		models.SeverityMedium:   2,
		models.SeverityHigh:     3,
		models.SeverityCritical: 4,
	}

	for _, f := range result.Findings {
		if severityMap[f.Severity] > severityMap[maxSeverity] {
			maxSeverity = f.Severity
		}
	}

	return maxSeverity
}
