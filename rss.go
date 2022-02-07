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
		Title:       categoryName + " streams",
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
	item.Title = stream.Title
	item.Created = stream.StartedAt
	item.Author = &feeds.Author{Name: stream.UserName}
	item.Link = &feeds.Link{Href: "https://www.twitch.tv/" + stream.UserLogin}
	item.Description = fmt.Sprintf("<a href='https://www.twitch.tv/popout/%s/chat?popout='>%s's Chat</a><br/>",
		stream.UserLogin, stream.UserName)
	item.Description += fmt.Sprintf("(%s|%d)", strings.ToUpper(stream.Language), stream.ViewerCount)
	return item
}
