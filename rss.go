package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gorilla/feeds"
	"github.com/nicklaw5/helix"
)

func toFeed(streams []helix.Stream) (string, error) {
	now := time.Now()
	feed := &feeds.Feed{
		Title:       "Twitch's `" + categoryName + "` streams",
		Link:        &feeds.Link{Href: "https://www.twitch.tv/directory/game/" + categoryName},
		Description: categoryName + " streams",
		Author:      &feeds.Author{Name: "twitch.tv"},
		Created:     now,
	}
	loadStreams(feed, streams)
	atom, err := feed.ToAtom()
	if err != nil {
		return "", err
	}
	return atom, nil
}

func loadStreams(feed *feeds.Feed, streams []helix.Stream) {
	var items []*feeds.Item
	for _, stream := range streams {
		items = append(items, toItem(stream))
	}
	feed.Items = items
	return
}

func toItem(stream helix.Stream) *feeds.Item {
	item := &feeds.Item{}
	item.Title = fmt.Sprintf("/%s/ %s",
		strings.ToUpper(stream.Language), stream.Title)
	item.Created = stream.StartedAt
	item.Author = &feeds.Author{Name: stream.UserName}
	item.Link = &feeds.Link{Href: "https://www.twitch.tv/" + stream.UserLogin}
	item.Description += fmt.Sprintf(
		"<a href='https://www.twitch.tv/%s/videos?filter=archives%%26sort%%3Dtime'>%s's videos archive</a><br/>",
		stream.UserLogin,
		stream.UserName,
	)
	item.Description += fmt.Sprintf(
		"<img alt='thumbnail' src='%s'/><br/>",
		strings.Replace(stream.ThumbnailURL, "{width}x{height}", "320x240", 1),
	)
	item.Description += fmt.Sprintf(
		"<a href='https://www.twitch.tv/popout/%s/chat?popout='>Live Chat</a>",
		stream.UserLogin,
	)
	return item
}
