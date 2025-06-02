build: cfg cmd launcher matcher applog
	go build -o build/rule-browser-selector cmd/main.go
install: build
	cp build/rule-browser-selector $(GOPATH)/bin/
