package matcher

import (
	"fmt"
	"github.com/IGLOU-EU/go-wildcard/v2"
	"github.com/nicumaxian/rule-browser-selector/applog"
	"github.com/nicumaxian/rule-browser-selector/cfg"
)

func MatchURL(url string, rules []cfg.Rule) (cfg.Browser, error) {
	for _, rule := range rules {
		for _, m := range rule.Match {
			matched := wildcard.Match(m, url)
			applog.Logger.
				With("matched", matched).
				With("match", m).
				Debug("checked for rule")
			if matched {
				return rule.Browser, nil
			}
		}
	}
	return cfg.Browser{}, fmt.Errorf("no match found")
}
