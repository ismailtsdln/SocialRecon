package engine

import (
	"context"
	"sync"
	"time"

	"github.com/ismailtsdln/socialrecon/internal/models"
	"github.com/ismailtsdln/socialrecon/internal/plugins"
)

// Engine orchestrates the scanning process
type Engine struct {
	config  models.Config
	plugins []plugins.Plugin
}

// NewEngine creates a new scanning engine
func NewEngine(cfg models.Config, enabledPlugins []plugins.Plugin) *Engine {
	return &Engine{
		config:  cfg,
		plugins: enabledPlugins,
	}
}

// Run executes the scan across all configured plugins
func (e *Engine) Run(ctx context.Context, target string) (*models.ScanResult, error) {
	result := &models.ScanResult{
		Target:    target,
		Findings:  []models.Finding{},
		StartTime: time.Now(),
	}

	findingChan := make(chan models.Finding)
	errorChan := make(chan error)
	var wg sync.WaitGroup

	// Worker pool pattern for plugins
	semaphore := make(chan struct{}, e.config.MaxConcurrency)

	for _, p := range e.plugins {
		wg.Add(1)
		go func(pl plugins.Plugin) {
			defer wg.Done()

			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			findings, err := pl.Check(ctx, target)
			if err != nil {
				errorChan <- err
				return
			}
			for _, f := range findings {
				findingChan <- f
			}
		}(p)
	}

	// Closer goroutine
	go func() {
		wg.Wait()
		close(findingChan)
		close(errorChan)
	}()

	// Collect findings
CollectLoop:
	for {
		select {
		case f, ok := <-findingChan:
			if !ok {
				break CollectLoop
			}
			result.Findings = append(result.Findings, f)
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-errorChan:
			// Log error but continue scanning with other plugins
			continue
		}
	}

	result.EndTime = time.Now()
	return result, nil
}
