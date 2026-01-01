package plugins

import (
	"context"

	"github.com/ismailtsdln/socialrecon/internal/models"
)

// Plugin defines the interface all social media modules must implement
type Plugin interface {
	Name() string
	Description() string
	// Check assesses the presence/risk of a specific indicator (username/brand)
	Check(ctx context.Context, target string) ([]models.Finding, error)
}
