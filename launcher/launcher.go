package launcher

import (
	"github.com/nicumaxian/rule-browser-selector/cfg"
	"os/exec"
)

func OpenURL(url string, browser cfg.Browser) error {
	args := make([]string, 0, 1+len(browser.Args))
	args = append(args, browser.Args...)
	args = append(args, url)

	cmd := exec.Command(
		browser.Command,
		args...,
	)
	return cmd.Start()
}
