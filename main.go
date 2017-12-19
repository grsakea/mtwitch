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
	s, conf, err := initApp(os.Getenv("CLIENT_ID"), "config.json")
	if err != nil {
		os.Exit(1)
	}

	followStream(conf.Streamers, s)

	select {}
}

func initApp(cID, confPath string) (twitch.Session, Config, error) {
	s := twitch.NewSession(cID)
	conf, err := loadConfig(confPath)
	return s, conf, err
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
