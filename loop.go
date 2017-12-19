package main

import (
	"log"
	"time"

	twitch "github.com/grsakea/go-twitch"
	"github.com/grsakea/hls"
)

var sleepFunc = time.Sleep

func followStream(channels []string, s twitch.Interface) {
	for _, channel := range channels {
		go startRecord(channel, s)
		sleepFunc(3 * time.Second)
	}
}

func startRecord(channel string, s twitch.Interface) {
	log.Println("start", channel)
	for {
		state, err := isOnline(channel, s)
		if err != nil {
			log.Println("error :", err)
		} else if state {
			recordStream(channel, s, hls.HLSDownloader{})
			log.Println("Stopping recording of", channel)
		}

		sleepFunc(30 * time.Second)
	}
}

func channelStatus(channel string, state bool) string {
	var str string
	if state {
		str = "online"
	} else {
		str = "offline"
	}
	return channel + " is now : " + str
}
