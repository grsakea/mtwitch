package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/grsakea/go-twitch"
)

type Config struct {
	Streamers []string `json:"streamers"`
}

func main() {
	s := twitch.NewSession()
	conf, err := loadConfig("config.json")
	if err != nil {
		os.Exit(1)
	}

	for i := 0; i < len(conf.Streamers); i++ {
		onl, err := isOnline(conf.Streamers[i], s)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println(conf.Streamers[i], onl)
	}
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
