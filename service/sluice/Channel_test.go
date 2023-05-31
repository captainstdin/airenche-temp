package sluice

import (
	"githun.com/captainstdin/airenche-temp/service"
	"testing"
)

func TestChannelInstance_Open(t *testing.T) {

	err := service.LoadEnv()
	if err != nil {
		t.Error(err)
		return
	}

	channel := ChannelInstance{
		Cid: "1111222",
		Sid: service.EnvSite.Sid,
	}
	channelErr := channel.Open("浙D88888")
	if err != nil {
		t.Error(channelErr)
		return
	}
}
