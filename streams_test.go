package main

import (
	"errors"
	"testing"
	"time"

	twitch "github.com/grsakea/go-twitch"
)

type fakeTwitchGetter struct {
}

func (s fakeTwitchGetter) GetStream(input twitch.GetStreamInput) (twitch.StreamList, error) {
	if input.UserLogin == "twitch" {
		sl := twitch.StreamList{Data: []twitch.Stream{{}}}
		return sl, nil
	} else {
		return twitch.StreamList{}, errors.New("test")
	}
}

func TestIsOnline(t *testing.T) {
	s := fakeTwitchGetter{}
	out, err := isOnline("twitch", s)
	if err != nil {
		t.Fail()
	}
	if !out {
		t.Fail()
	}
}

func TestIsOnlineFail(t *testing.T) {
	s := fakeTwitchGetter{}
	_, err := isOnline("not_twitch", s)
	if err == nil {
		t.Fail()
	}
}

func TestStreamFilename(t *testing.T) {
	s := twitch.Stream{Title: "@fake_stream!|"}
	tim, _ := time.Parse(time.RFC3339, "2017-01-01T15:04:05Z")

	out := streamFilename(s, tim)
	expected := "17-01-01_15:04-_fake_stream_.mp4"
	if out != expected {
		t.Fatal(out, expected)
	}
}

type fakeDownloader struct {
}

func (d fakeDownloader) Download(playlistURL string, target string) {
}

func TestRecordStream(t *testing.T) {
	tw := fakeTwitchFollowStream{t}
	d := fakeDownloader{}

	recordStream("test_stream", tw, d)
}
