package matcher_test

import (
	"fmt"
	"github.com/nicumaxian/rule-browser-selector/cfg"
	"github.com/nicumaxian/rule-browser-selector/matcher"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchURL(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		rules      []cfg.Rule
		expected   string
		expectsErr bool
	}{
		{
			name: "Match with exact pattern",
			url:  "https://example.com",
			rules: []cfg.Rule{
				{Match: "https://example.com", Browser: "firefox"},
			},
			expected:   "firefox",
			expectsErr: false,
		},
		{
			name: "Match with wildcard pattern",
			url:  "https://example.com/page",
			rules: []cfg.Rule{
				{Match: "https://example.com/*", Browser: "chrome"},
			},
			expected:   "chrome",
			expectsErr: false,
		},
		{
			name: "No matching rule",
			url:  "https://notmatched.com",
			rules: []cfg.Rule{
				{Match: "https://example.com/*", Browser: "chrome"},
			},
			expected:   "",
			expectsErr: true,
		},
		{
			name: "Multiple rules, match first",
			url:  "https://example.com",
			rules: []cfg.Rule{
				{Match: "https://example.com", Browser: "firefox"},
				{Match: "https://example.com/*", Browser: "chrome"},
			},
			expected:   "firefox",
			expectsErr: false,
		},
		{
			name: "Error in pattern syntax",
			url:  "https://example.com",
			rules: []cfg.Rule{
				{Match: "[invalid-pattern", Browser: "firefox"},
			},
			expected:   "",
			expectsErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			browser, err := matcher.MatchURL(tt.url, tt.rules)

			if tt.expectsErr {
				assert.Error(t, err, "expected an error but got none")
				assert.Empty(t, browser, "expected no browser to be returned")
			} else {
				assert.NoError(t, err, fmt.Sprintf("did not expect an error but got: %v", err))
				assert.Equal(t, tt.expected, browser, "expected browser mismatch")
			}
		})
	}
}
