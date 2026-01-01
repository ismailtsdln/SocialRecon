package socialrecon

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/ismailtsdln/socialrecon/internal/engine"
	"github.com/ismailtsdln/socialrecon/internal/models"
	"github.com/ismailtsdln/socialrecon/internal/plugins"
	"github.com/ismailtsdln/socialrecon/internal/plugins/github"
	"github.com/ismailtsdln/socialrecon/internal/plugins/instagram"
	"github.com/ismailtsdln/socialrecon/internal/plugins/twitter"
	"github.com/ismailtsdln/socialrecon/internal/report"
	"github.com/ismailtsdln/socialrecon/internal/scanner"
	"github.com/ismailtsdln/socialrecon/internal/scoring"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "socialrecon",
		Short: "SocialRecon is a high-performance OSINT security scanner",
		Long:  `SocialRecon identifies social media presence, impersonation risks, and brand abuse.`,
	}

	scanCmd = &cobra.Command{
		Use:   "scan [target]",
		Short: "Scan a domain, brand, or username",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runScan,
	}

	// Flags
	jsonOutput bool
	verbose    bool
)

func init() {
	scanCmd.Flags().BoolVar(&jsonOutput, "json", false, "Output results in JSON format")
	scanCmd.Flags().BoolVar(&verbose, "verbose", false, "Enable verbose output")
	rootCmd.AddCommand(scanCmd)
}

func Execute() error {
	return rootCmd.Execute()
}

func runScan(cmd *cobra.Command, args []string) error {
	target := args[0]

	if !jsonOutput {
		color.Cyan("🚀 Starting SocialRecon scan for: %s", target)
	}

	cfg := models.Config{
		MaxConcurrency: 10,
		Timeout:        30 * time.Second,
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	// 1. Initial Discovery (if target is a domain)
	foundUsernames := make(map[string]bool)
	isDomain := strings.Contains(target, ".") || strings.HasPrefix(target, "http")

	if isDomain {
		extractor := scanner.NewExtractor()
		if !jsonOutput {
			color.Yellow("🔍 Extracting social links from domain...")
		}
		links, err := extractor.ExtractSocialLinks(target)
		if err == nil {
			for _, l := range links {
				foundUsernames[l.Username] = true
			}
		}
	} else {
		foundUsernames[target] = true
	}

	// 2. Setup Plugins & Engine
	enabledPlugins := []plugins.Plugin{
		github.NewPlugin(),
		twitter.NewPlugin(),
		instagram.NewPlugin(),
	}
	eng := engine.NewEngine(cfg, enabledPlugins)

	// 3. Run Scanning for each found username
	finalResult := &models.ScanResult{
		Target:    target,
		Findings:  []models.Finding{},
		StartTime: time.Now(),
	}

	for username := range foundUsernames {
		if !jsonOutput && verbose {
			fmt.Printf("   -> Scanning username: %s\n", username)
		}
		res, err := eng.Run(ctx, username)
		if err == nil {
			finalResult.Findings = append(finalResult.Findings, res.Findings...)
		}
	}
	finalResult.EndTime = time.Now()

	// 4. Calculate risk score
	scorer := scoring.NewScoringEngine()
	finalResult.RiskScore = scorer.Calculate(finalResult)

	reporter := report.NewReporter()
	if jsonOutput {
		reporter.ExportJSON(finalResult, "")
	} else {
		color.Green("✅ Scan completed in %v", finalResult.EndTime.Sub(finalResult.StartTime))
		reporter.PrintSummary(finalResult)
	}

	return nil
}
