package main

import (
	"log"
	"path"
	"regexp"
	"time"

	twitch "github.com/grsakea/go-twitch"
	"github.com/grsakea/hls"
)

func isOnline(name string, s twitch.GetStreamer) (bool, error) {
	data, err := s.GetStream(twitch.GetStreamInput{UserLogin: name})
	if err != nil {
		return false, err
	}
	return len(data.Data) == 1, nil
}

func recordStream(channel string, s twitch.Interface, d hls.Downloader) {
	log.Println("Starting recording of", channel)
	st, _ := s.GetStream(twitch.GetStreamInput{UserLogin: channel})
	stURL, _ := s.ExtractStreamUrl(channel)
	filename := streamFilename(st.Data[0], time.Now())
	d.Download(stURL[0].URL, path.Join(channel, filename))
}

func streamFilename(s twitch.Stream, t time.Time) string {
	outTime := t.Format("06-01-02-15:04")
	re := regexp.MustCompile(`[ \|\!|@]+`)
	outChan := re.ReplaceAllString(s.Title, "_")

	return outTime + "-" + outChan + ".mp4"
}
