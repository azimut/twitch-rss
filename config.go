package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func readConfig() error {
	currentPaths := strings.Split(os.Args[0], "/")
	currentPath := strings.Join(currentPaths[:len(currentPaths)-1], "/")
	secretPath := currentPath + "/twitch-rss.secret"
	bytes, err := ioutil.ReadFile(secretPath)
	if err != nil {
		return err
	}
	lines := strings.Split(string(bytes), "\n")
	clientID = strings.TrimSpace(lines[0])
	clientSecret = strings.TrimSpace(lines[1])
	return nil
}
