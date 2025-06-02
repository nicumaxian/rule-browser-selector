package cfg

import "github.com/pkg/errors"

type Rule struct {
	Match   []string `yaml:"match"`
	Browser Browser  `yaml:"browser"`
}

type Config struct {
	Rules          []Rule  `yaml:"rules"`
	DefaultBrowser Browser `yaml:"default_browser"`
}

type Browser struct {
	Command string   `yaml:"command"`
	Args    []string `yaml:"args"`
}

func (c Config) Validate() error {
	if c.DefaultBrowser.Command == "" {
		return errors.New("default browser is required")
	}

	return nil
}
