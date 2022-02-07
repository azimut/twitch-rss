package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nicklaw5/helix"
)

var (
	categoryName string
	clientID     string
	clientSecret string
	nResults     uint
)

func init() {
	flag.UintVar(&nResults, "n", 10, "number of results to return")
	if err := readConfig(); err != nil {
		log.Fatal("could not read credentials from twitch-rss.secret")
		os.Exit(1)
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("missing category argument")
		os.Exit(1)
	}
	categoryName = flag.Args()[0]

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
		First:    int(nResults),
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
