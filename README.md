# rule-browser-selector

## Install

If you have `GOPATH` set, run following command:
```bash
make install
```

## Config file

Create a file `~/.config/rule-browser-selector/config` with a configuration inside:

```yaml
rules:
  - match: 
      - "*notion.so/*"
      - "*parabol.co*"
      - "*gemini*"
      - "*datadog*"
    browser: 
      command: google-chrome
      args:
        - "--profile-directory=Profile\ 1"
  - match: 
      - "*github.com*"
    browser: 
      command: google-chrome
      args:
        - "--profile-directory=Profile\ 2"
default_browser:
  command: google-chrome
  args:
    - "--profile-directory=Default"
```

Matching rules are parsed with [IGLOU-EU/go-wildcard](https://github.com/IGLOU-EU/go-wildcard)
