package main

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/RockinTechBomb/gobot/config"
	"github.com/RockinTechBomb/gobot/datas"
)

func main() {
	config.SetConfig()
	anaconda.SetConsumerKey(config.ConsumerKey())
	anaconda.SetConsumerSecret(config.ConsumerSecret())
	api := anaconda.NewTwitterApi(config.AccessToken(), config.AccessTokenSecret())
	datas.SetData()
	tweet := datas.Choice(datas.Tweets())
	// tweet
	api.PostTweet(tweet, nil)
}
