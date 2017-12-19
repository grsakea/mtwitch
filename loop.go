package main

import (
	"log"
	"regexp"
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
			recordStream(channel, s)
			log.Println("Stopping recording of", channel)
		}

		time.Sleep(30 * time.Second)
	}
}

func recordStream(channel string, s twitch.Interface) {
	log.Println("Starting recording of", channel)
	st, _ := s.GetStream(twitch.GetStreamInput{UserLogin: channel})
	stURL, _ := s.ExtractStreamUrl(channel)
	filename := streamFilename(st.Data[0], time.Now())
	hls.Download(stURL[0].URL, filename)
}

func streamFilename(s twitch.Stream, t time.Time) string {
	outTime := t.Format("06-01-02_15:04")
	re := regexp.MustCompile(`[ \|\!|@]+`)
	outChan := re.ReplaceAllString(s.Title, "_")

	return outTime + "-" + outChan + ".mp4"
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
