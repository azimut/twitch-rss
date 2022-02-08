package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nicklaw5/helix"
)

var (
	categoryName string
	clientID     string
	clientSecret string
	nResults     uint
	rawlang      string
)

func init() {
	flag.StringVar(
		&rawlang,
		"lang",
		"en,es",
		"comma separated list of 2 digit languages to filter by",
	)
	flag.UintVar(&nResults, "n", 10, "number of results to return")
	if err := readConfig(); err != nil {
		log.Fatal("could not read credentials from twitch-rss.secret")
		os.Exit(1)
	}
}

func main() {
	if rss, err := run(); err != nil {
		flag.Usage()
		log.Fatal(err)
	} else {
		fmt.Println(rss)
	}
}

func run() (string, error) {
	flag.Parse()
	if flag.NArg() != 1 {
		return "", errors.New("missing category argument")
	}
	categoryName = flag.Args()[0]
	client, err := login()
	if err != nil {
		return "", err
	}
	games, err := client.GetGames(&helix.GamesParams{
		Names: []string{categoryName},
	})
	if err != nil {
		return "", err
	}
	gameID := games.Data.Games[0].ID
	lang := strings.Split(rawlang, ",")
	streams, err := client.GetStreams(&helix.StreamsParams{
		First:    int(nResults),
		Type:     "live",
		Language: lang,
		GameIDs:  []string{gameID},
	})
	if err != nil {
		return "", err
	}
	atom, err := toFeed(streams.Data.Streams)
	if err != nil {
		return "", err
	}
	return atom, nil
}
