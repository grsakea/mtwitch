package main

import "testing"

func TestIsOnline(t *testing.T) {
	s := fakeTwitch{}
	out, err := isOnline("twitch", s)
	if err != nil {
		t.Fail()
	}
	if !out {
		t.Fail()
	}
}

func TestIsOnlineFail(t *testing.T) {
	s := fakeTwitch{}
	_, err := isOnline("not_twitch", s)
	if err == nil {
		t.Fail()
	}
}
