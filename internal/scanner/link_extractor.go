package scanner

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// Info represents a discovered social media indicator
type Info struct {
	Platform string
	Username string
	URL      string
}

// Extractor handles fetching and parsing social links
type Extractor struct {
	client *http.Client
}

func NewExtractor() *Extractor {
	return &Extractor{
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

// ExtractSocialLinks scans a URL for social media profiles
func (e *Extractor) ExtractSocialLinks(targetURL string) ([]Info, error) {
	if !strings.HasPrefix(targetURL, "http") {
		targetURL = "http://" + targetURL
	}

	resp, err := e.client.Get(targetURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch URL: %s (status %d)", targetURL, resp.StatusCode)
	}

	return e.parseHTML(resp.Body), nil
}

func (e *Extractor) parseHTML(r io.Reader) []Info {
	var infos []Info
	z := html.NewTokenizer(r)

	// Regex for social platforms
	patterns := map[string]*regexp.Regexp{
		"Twitter":   regexp.MustCompile(`(?:https?://)?(?:www\.)?twitter\.com/([a-zA-Z0-9_]{1,15})`),
		"GitHub":    regexp.MustCompile(`(?:https?://)?(?:www\.)?github\.com/([a-zA-Z0-9-]{1,39})`),
		"Instagram": regexp.MustCompile(`(?:https?://)?(?:www\.)?instagram\.com/([a-zA-Z0-9_\.]{1,30})`),
		"TikTok":    regexp.MustCompile(`(?:https?://)?(?:www\.)?tiktok\.com/@([a-zA-Z0-9_\.]{2,24})`),
		"LinkedIn":  regexp.MustCompile(`(?:https?://)?(?:www\.)?linkedin\.com/in/([a-zA-Z0-9-]{3,100})`),
	}

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return infos
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == "a" {
				for _, a := range t.Attr {
					if a.Key == "href" {
						val := a.Val
						for platform, re := range patterns {
							matches := re.FindStringSubmatch(val)
							if len(matches) > 1 {
								infos = append(infos, Info{
									Platform: platform,
									Username: matches[1],
									URL:      val,
								})
							}
						}
					}
				}
			}
		}
	}
}
