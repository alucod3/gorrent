package config

import (
	"os"
	"path/filepath"
	"time"
)

// Config stores application settings
type Config struct {
	// General Settings
	AppName    string
	AppVersion string

	// Download Settings
	DownloadPath          string
	Seed                  bool
	ProgressCheckInterval time.Duration

	// Validation Standards
	MagnetPattern    string
	TorrentExtension string
}

// LoadDefaultConfig loads default settings and ensures the download path exists
func LoadDefaultConfig() (*Config, error) {
	cfg := &Config{
		AppName:               "Gorrent",
		AppVersion:            "0.1",
		DownloadPath:          getDefaultDownloadPath(),
		Seed:                  true,
		ProgressCheckInterval: 1 * time.Second,
		MagnetPattern:         `(?i)^magnet:\?xt=urn:btih:[a-zA-Z0-9]{32,40}`,
		TorrentExtension:      ".torrent",
	}

	if err := cfg.EnsureDownloadPath(); err != nil {
		return nil, err
	}

	return cfg, nil
}

// getDefaultDownloadPath returns the default path for downloads
func getDefaultDownloadPath() string {
	return filepath.Join(os.Getenv("HOME"), "Downloads")
}

// EnsureDownloadPath ensures that the download directory exists
func (c *Config) EnsureDownloadPath() error {
	if _, err := os.Stat(c.DownloadPath); os.IsNotExist(err) {
		return os.MkdirAll(c.DownloadPath, 0755)
	}
	return nil
}
