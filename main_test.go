package main

import (
	"testing"
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

func TestLoadBadConfig(t *testing.T) {
	_, err := loadConfig("main.go")
	if err == nil {
		t.Error("Should have been an error")
	}
}

func TestLoadAbsentConfig(t *testing.T) {
	_, err := loadConfig("fixtures/invalid_config.json")
	if err == nil {
		t.Error("Should have been an error")
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

func TestInitAppFail(t *testing.T) {
	fakeCID := "TEST"
	fakeConf := "badfile"

	_, _, err := initApp(fakeCID, fakeConf)
	if err == nil {
		t.Fail()
	}
}
