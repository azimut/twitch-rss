SRCS=$(wildcard *.go)

twitch-rss: $(SRCS)
	~/go/bin/go build .

.PHONY: clean
clean:
	rm -f twich-rss

.PHONY: install
install: twitch-rss
	mv twitch-rss $(HOME)/.newsboat/feeds/
