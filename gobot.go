package main

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/RockinTechBomb/gobot/config"
)

var (
	api = nil
)

func main() {
	config.SetConfig()
	anaconda.SetConsumerKey(config.ConsumerKey())
	anaconda.SetConsumerSecret(config.ConsumerSecret())
	api := anaconda.NewTwitterApi(config.AccessToken(), config.AccessTokenSecret())
}
