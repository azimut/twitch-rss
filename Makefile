SRCS=$(wildcard *.go)

twitch-rss: $(SRCS)
	go build -v

.PHONY: clean
clean: ; go clean

.PHONY: install
install: twitch-rss
	mv twitch-rss $(HOME)/.newsboat/feeds/
