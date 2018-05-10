package main

import (
	"log"
	"os"
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

func recordStream(channel string, conf Config, s twitch.Interface, d hls.Downloader) error {
	log.Println("Starting recording of", channel)
	os.MkdirAll(path.Join(conf.Location, channel), 0775)
	st, _ := s.GetStream(twitch.GetStreamInput{UserLogin: channel})
	stURL, err := s.ExtractStreamUrl(channel)
	if err != nil {
		return err
	}
	filename := streamFilename(st.Data[0], time.Now())
	log.Println(path.Join(conf.Location, channel, filename))
	d.Download(stURL[0].URL, path.Join(conf.Location, channel, filename))
	return nil
}

func streamFilename(s twitch.Stream, t time.Time) string {
	outTime := t.Format("06-01-02-15:04")
	re := regexp.MustCompile(`[\.'\/ \|\!|@]+`)
	outChan := re.ReplaceAllString(s.Title, "_")

	return outTime + "-" + outChan + ".mp4"
}
