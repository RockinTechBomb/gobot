package main

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
	"github.com/RockinTechBomb/gobot/config"
)

func main() {
	config.SetConfig()
	anaconda.SetConsumerKey(config.ConsumerKey())
	anaconda.SetConsumerSecret(config.ConsumerSecret())
	api := anaconda.NewTwitterApi(config.AccessToken(), config.AccessTokenSecret())
	twitterStream := api.PublicStreamSample(nil)
	x := <-twitterStream.C
	if x != nil {
		fmt.Println("OK gobot")
	} else {
		fmt.Println("NG gobot")
	}
}
