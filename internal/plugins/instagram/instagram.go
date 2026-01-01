package instagram

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ismailtsdln/socialrecon/internal/models"
)

type InstagramPlugin struct {
	client *http.Client
}

func NewPlugin() *InstagramPlugin {
	return &InstagramPlugin{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (p *InstagramPlugin) Name() string {
	return "Instagram"
}

func (p *InstagramPlugin) Description() string {
	return "Checks for Instagram profiles"
}

func (p *InstagramPlugin) Check(ctx context.Context, target string) ([]models.Finding, error) {
	url := fmt.Sprintf("https://www.instagram.com/%s/", target)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

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
			Indicator:   "instagram_profile",
			Value:       target,
			Status:      "exists",
			Severity:    models.SeverityInfo,
			Description: fmt.Sprintf("Instagram profile found: %s", url),
			Timestamp:   time.Now(),
		})
	case http.StatusNotFound:
		findings = append(findings, models.Finding{
			PluginName:  p.Name(),
			Indicator:   "instagram_profile",
			Value:       target,
			Status:      "available",
			Severity:    models.SeverityLow,
			Description: fmt.Sprintf("Instagram username '%s' is available", target),
			Timestamp:   time.Now(),
		})
	}

	return findings, nil
}
