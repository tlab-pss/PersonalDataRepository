package plugin_service

import (
	"time"
)

type PluginService struct {
	ID            string
	Name          string
	BigCategoryId string
	CreatedAt     time.Time
}
