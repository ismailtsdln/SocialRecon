package report

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ismailtsdln/socialrecon/internal/models"
)

// Reporter handles outputting scan results
type Reporter struct{}

func NewReporter() *Reporter {
	return &Reporter{}
}

// ExportJSON writes the result to a JSON file or stdout
func (r *Reporter) ExportJSON(result *models.ScanResult, filename string) error {
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}

	if filename == "" {
		fmt.Println(string(data))
		return nil
	}

	return os.WriteFile(filename, data, 0644)
}

// PrintSummary outputs a text-based summary to the console
func (r *Reporter) PrintSummary(result *models.ScanResult) {
	fmt.Printf("\n--- Scan Summary ---\n")
	fmt.Printf("Target:     %s\n", result.Target)
	fmt.Printf("Findings:   %d\n", len(result.Findings))
	fmt.Printf("Risk Score: %.2f/100\n", result.RiskScore)
	fmt.Printf("Duration:   %v\n", result.EndTime.Sub(result.StartTime))
	fmt.Printf("-------------------\n")
}
