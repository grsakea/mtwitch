package main

import (
	"errors"
	"testing"

	twitch "github.com/grsakea/go-twitch"
)

func TestLoadConfig(t *testing.T) {
	conf, err := loadConfig("fixtures/config.json")
	if err != nil {
		t.Error("Error during conf parsing, ", err)
	}
	if len(conf.Streamers) != 1 {
		t.Error("Invalid number of streamer")
	}
	if conf.Streamers[0] != "twitch" {
		t.Error("Conf parsing failed")
	}
}

func TestLoadAbsentConfig(t *testing.T) {
	_, err := loadConfig("fixtures/invalid_config.json")
	if err == nil {
		t.Error("Should have been an error")
	}
}

type fakeTwitch struct {
}

func (s fakeTwitch) GetStream(input twitch.GetStreamInput) (twitch.StreamList, error) {
	if input.UserLogin == "twitch" {
		sl := twitch.StreamList{Data: []twitch.Stream{{}}}
		return sl, nil
	} else {
		return twitch.StreamList{}, errors.New("test")
	}
}

func TestInitApp(t *testing.T) {
	fakeCID := "TEST"
	fakeConf := "fixtures/config.json"

	_, _, err := initApp(fakeCID, fakeConf)
	if err != nil {
		t.Fail()
	}
}
