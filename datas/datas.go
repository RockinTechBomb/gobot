package datas

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/RockinTechBomb/gobot/config"
)

type Datas struct {
	Tweets []string `json:"tweets"`
}

var (
	tweets = []string{}
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SetData() {
	var all_tweet Datas
	{
		data, err_file := ioutil.ReadFile(config.TWEETS_PATH)
		check(err_file)
		err_json := json.Unmarshal(data, &all_tweet)
		check(err_json)
	}
	tweets = all_tweet.Tweets
}

func Tweets() []string {
	return tweets
}

func Choice(s []string) string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(s))
	return s[i]
}
