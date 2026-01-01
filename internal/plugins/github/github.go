package github

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ismailtsdln/socialrecon/internal/models"
)

type GitHubPlugin struct {
	client *http.Client
}

func NewPlugin() *GitHubPlugin {
	return &GitHubPlugin{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (p *GitHubPlugin) Name() string {
	return "GitHub"
}

func (p *GitHubPlugin) Description() string {
	return "Checks for GitHub profiles and repository availability"
}

func (p *GitHubPlugin) Check(ctx context.Context, target string) ([]models.Finding, error) {
	url := fmt.Sprintf("https://github.com/%s", target)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var findings []models.Finding

	switch resp.StatusCode {
	case http.StatusOK:
		findings = append(findings, models.Finding{
			PluginName:  p.Name(),
			Indicator:   "github_profile",
			Value:       target,
			Status:      "exists",
			Severity:    models.SeverityInfo,
			Description: fmt.Sprintf("GitHub profile found: %s", url),
			Timestamp:   time.Now(),
		})
	case http.StatusNotFound:
		findings = append(findings, models.Finding{
			PluginName:  p.Name(),
			Indicator:   "github_profile",
			Value:       target,
			Status:      "available",
			Severity:    models.SeverityLow,
			Description: fmt.Sprintf("GitHub username '%s' is available for registration", target),
			Timestamp:   time.Now(),
		})
	}

	return findings, nil
}
