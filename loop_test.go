package main

import (
	"testing"

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

func TestFollowStream(t *testing.T) {
	followStream([]string{"test_stream"}, fakeTwitchFollowStream{})
}
