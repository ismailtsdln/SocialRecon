package twitter

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ismailtsdln/socialrecon/internal/models"
)

type TwitterPlugin struct {
	client *http.Client
}

func NewPlugin() *TwitterPlugin {
	return &TwitterPlugin{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (p *TwitterPlugin) Name() string {
	return "Twitter"
}

func (p *TwitterPlugin) Description() string {
	return "Checks for Twitter/X profiles"
}

func (p *TwitterPlugin) Check(ctx context.Context, target string) ([]models.Finding, error) {
	url := fmt.Sprintf("https://twitter.com/%s", target)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	// Twitter often blocks scrapers, using a real-looking UA might help for passive check
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
			Indicator:   "twitter_profile",
			Value:       target,
			Status:      "exists",
			Severity:    models.SeverityInfo,
			Description: fmt.Sprintf("Twitter profile found: %s", url),
			Timestamp:   time.Now(),
		})
	case http.StatusNotFound:
		findings = append(findings, models.Finding{
			PluginName:  p.Name(),
			Indicator:   "twitter_profile",
			Value:       target,
			Status:      "available",
			Severity:    models.SeverityLow,
			Description: fmt.Sprintf("Twitter username '%s' is available or suspended", target),
			Timestamp:   time.Now(),
		})
	}

	return findings, nil
}
