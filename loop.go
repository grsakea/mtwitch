package main

import (
	"log"
	"time"

	twitch "github.com/grsakea/go-twitch"
	"github.com/grsakea/hls"
)

var sleepFunc = time.Sleep

func followStream(conf Config, s twitch.Interface, d hls.Downloader) {
	for _, channel := range conf.Streamers {
		go startRecord(channel, conf, s, d)
		sleepFunc(3 * time.Second)
	}
}

func startRecord(channel string, conf Config, s twitch.Interface, d hls.Downloader) {
	log.Println("start", channel)
	for {
		loopStreamRecord(channel, conf, s, d)
		sleepFunc(30 * time.Second)
	}
}

func loopStreamRecord(channel string, conf Config, s twitch.Interface, d hls.Downloader) {
	state, err := isOnline(channel, s)
	if err != nil {
		log.Println("error :", err)
	} else if state {
		err := recordStream(channel, conf, s, d)
		if err != nil {
			log.Println("Error during recording :", err)
			return
		}
		log.Println("Stopping recording of", channel)
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
