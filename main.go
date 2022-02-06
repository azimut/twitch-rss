package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gorilla/feeds"
	"github.com/nicklaw5/helix"
)

var (
	categoryName = "Software and Game Development"
	clientID     string
	clientSecret string
)

func init() {
	clientID = os.Getenv("TWITCH_ID")
	clientSecret = os.Getenv("TWITCH_SECRET")
	if clientID == "" || clientSecret == "" {
		log.Fatal("needs TWITCH_ID and TWITCH_SECRET environment variables set")
		os.Exit(1)
	}
}

func login() (*helix.Client, error) {
	client, err := helix.NewClient(&helix.Options{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	})
	if err != nil {
		return nil, err
	}
	resp, err := client.RequestAppAccessToken([]string{"user:read:email"})
	if err != nil {
		return nil, err
	}
	client.SetAppAccessToken(resp.Data.AccessToken)
	return client, nil
}

func toItem(stream helix.Stream) *feeds.Item {
	item := &feeds.Item{}
	item.Title = fmt.Sprintf("(%s|%s|%d) %s", strings.ToUpper(stream.Language), stream.UserLogin, stream.ViewerCount, stream.Title)
	item.Created = stream.StartedAt
	item.Author = &feeds.Author{Name: stream.UserName}
	item.Link = &feeds.Link{Href: "https://www.twitch.tv/" + stream.UserLogin}
	return item
}

func loadStreams(feed *feeds.Feed, streams []helix.Stream) {
	var items []*feeds.Item
	for _, stream := range streams {
		items = append(items, toItem(stream))
	}
	feed.Items = items
	return
}

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

func main() {
	client, err := login()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	games, err := client.GetGames(&helix.GamesParams{
		Names: []string{categoryName},
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	gameID := games.Data.Games[0].ID
	streams, err := client.GetStreams(&helix.StreamsParams{
		First:    10,
		Type:     "live",
		Language: []string{"en", "es"},
		GameIDs:  []string{gameID},
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	atom, err := toFeed(streams.Data.Streams)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", atom)
}
