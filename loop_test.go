package main

import (
	"testing"
	"time"

	twitch "github.com/grsakea/go-twitch"
)

type fakeTwitchFollowStream struct {
	t *testing.T
}

func (s fakeTwitchFollowStream) GetStream(input twitch.GetStreamInput) (twitch.StreamList, error) {
	if input.UserLogin != "test_stream" {
		s.t.Fail()
	}
	return twitch.StreamList{}, nil
}

func (s fakeTwitchFollowStream) ExtractStreamUrl(name string) ([]twitch.HLSStream, error) {
	return nil, nil
}

func fakeSleep(time.Duration) {
}

func TestFollowStream(t *testing.T) {
	sleepFunc = fakeSleep
	followStream([]string{"test_stream"}, fakeTwitchFollowStream{})
}

func TestChannelStatus(t *testing.T) {
	data := []struct {
		channel string
		state   bool
		out     string
	}{
		{"twitch", true, "twitch is now : online"},
		{"aswes", false, "aswes is now : offline"},
	}

	for _, i := range data {
		out := channelStatus(i.channel, i.state)
		if out != i.out {
			t.Fatal("expected :\n", i.out, " got :\n", out)
		}
	}
}

func TestStreamFilename(t *testing.T) {
	s := twitch.Stream{Title: "@fake_stream!|"}
	tim, _ := time.Parse(time.RFC3339, "2017-01-01T15:04:05Z")
	out := streamFilename(s, tim)
	if out != "17-01-01_15:04-_fake_stream_.mp4" {
		t.Fail()
	}
}
