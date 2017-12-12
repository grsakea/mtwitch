package main

import (
	"fmt"
	"time"

	twitch "github.com/grsakea/go-twitch"
)

func followStream(channels []string, s twitch.Interface) {
	for _, channel := range channels {
		go startRecord(channel, s)
		time.Sleep(3 * time.Second)
	}
}

func startRecord(channel string, s twitch.Interface) {
	fmt.Println("start" + channel)
	state := false
	for {
		new_st, err := isOnline(channel, s)
		if err != nil {
			fmt.Println("error :", err)
		} else if state != new_st {
			state = new_st
			var str string
			if state {
				str = "online"
			} else {
				str = "offline"
			}
			fmt.Println(channel + " is now : " + str)
		}

		time.Sleep(30 * time.Second)
	}
}
