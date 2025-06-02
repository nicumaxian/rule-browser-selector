package cfg

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func Load(cfg *Config) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/rule-browser-selector/")
	viper.AddConfigPath("$HOME/.config/rule-browser-selector")

	err := viper.ReadInConfig()
	if err != nil {
		return errors.Wrap(err, "failed to read config")
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal config")
	}

	err = viper.UnmarshalKey("default_browser", &cfg.DefaultBrowser)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal config")
	}

	return nil
}
