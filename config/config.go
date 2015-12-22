package config

import (
	"encoding/json"
	"io/ioutil"
)

type ApiConf struct {
	ConsumerKey       string `json:"consumer_key"`
	ConsumerSecret    string `json:"consumer_secret"`
	AccessToken       string `json:"access_token"`
	AccessTokenSecret string `json:"access_token_secret"`
}

var (
	consumer_key        = ""
	consumer_secret     = ""
	access_token        = ""
	access_token_secret = ""
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SetConfig() {
	var apiConf ApiConf
	{
		data, err_file := ioutil.ReadFile(CONF_PATH)
		check(err_file)
		err_json := json.Unmarshal(data, &apiConf)
		check(err_json)
	}
	consumer_key = apiConf.ConsumerKey
	consumer_secret = apiConf.ConsumerSecret
	access_token = apiConf.AccessToken
	access_token_secret = apiConf.AccessTokenSecret
}

func ConsumerKey() string {
	return consumer_key
}

func ConsumerSecret() string {
	return consumer_secret
}

func AccessToken() string {
	return access_token
}

func AccessTokenSecret() string {
	return access_token_secret
}
