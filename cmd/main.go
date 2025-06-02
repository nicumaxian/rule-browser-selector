package main

import (
	"github.com/nicumaxian/rule-browser-selector/applog"
	"github.com/nicumaxian/rule-browser-selector/cfg"
	"github.com/nicumaxian/rule-browser-selector/launcher"
	"github.com/nicumaxian/rule-browser-selector/matcher"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	defer func() {
		// put an artificial sleep before closing app, otherwise systemctl will not handle this process properly.
		time.Sleep(2 * time.Second)
	}()

	// Default configuration embedded in the application
	useCfg := cfg.Config{}

	rootCmd := &cobra.Command{
		Use:   "browser-selector",
		Short: "Selects the browser based on URL patterns",
	}

	openCmd := &cobra.Command{
		Use:   "open [URL]",
		Short: "Open a URL using the appropriate browser",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			err := cfg.Load(&useCfg)
			if err != nil {
				return err
			}

			return useCfg.Validate()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			url := args[0]

			applog.Logger.
				With("args", args).
				Debug("opening url")

			browser, err := matcher.MatchURL(url, useCfg.Rules)
			if err != nil {
				applog.Logger.Debug("no matches, using default")
				browser = useCfg.DefaultBrowser
			}

			applog.Logger.
				With("browser", browser).
				Debug("opening with chosen browser")
			return launcher.OpenURL(url, browser)
		},
	}

	rootCmd.AddCommand(openCmd)

	if err := rootCmd.Execute(); err != nil {
		applog.Logger.Error(err.Error())
		os.Exit(1)
	}
}
