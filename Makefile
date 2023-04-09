SRCS=$(wildcard *.go)

twitch-rss: $(SRCS)
	go build -v -ldflags="-s -w"
	ls -lh $@

.PHONY: clean
clean: ; go clean

.PHONY: update
update:
	go list -m -u all
	go get -u
	go mod tidy

.PHONY: install
install: twitch-rss
	upx twitch-rss
	mv twitch-rss $(HOME)/.newsboat/feeds/
