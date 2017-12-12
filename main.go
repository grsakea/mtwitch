package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/grsakea/go-twitch"
)

type Config struct {
	Streamers []string `json:"streamers"`
}

func main() {
	s := twitch.NewSession(os.Getenv("CLIENT_ID"))
	conf, err := loadConfig("config.json")
	if err != nil {
		os.Exit(1)
	}

	followStream(conf.Streamers, s)

	select {}
}

func loadConfig(path string) (Config, error) {
	var conf Config
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	err = json.Unmarshal(dat, &conf)
	if err != nil {
		return Config{}, err
	}
	return conf, nil
}

func isOnline(name string, s twitch.GetStreamer) (bool, error) {
	data, err := s.GetStream(twitch.GetStreamInput{UserLogin: name})
	if err != nil {
		return false, err
	}
	return len(data.Data) == 1, nil
}
