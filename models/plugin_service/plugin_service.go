package plugin_service

import (
	"time"
)

type PluginService struct {
	ID            string
	Name          string
	BigCategoryID string
	CreatedAt     time.Time
}
