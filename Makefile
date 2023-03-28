SRCS=$(wildcard *.go)

twitch-rss: $(SRCS)
	go build -v -ldflags="-s -w"
	ls -lh $@

.PHONY: clean
clean: ; go clean

.PHONY: install
install: twitch-rss
	mv twitch-rss $(HOME)/.newsboat/feeds/
