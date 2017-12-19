package main

import twitch "github.com/grsakea/go-twitch"

func isOnline(name string, s twitch.GetStreamer) (bool, error) {
	data, err := s.GetStream(twitch.GetStreamInput{UserLogin: name})
	if err != nil {
		return false, err
	}
	return len(data.Data) == 1, nil
}
