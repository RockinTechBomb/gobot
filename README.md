gobot
===============
This is twitterbot made with go

## Using

Please create your twitter acount & getting  API keys

Reference URL Links

[English](http://www.amelt.net/en/iwm/programming-iwm/web-app-development/3241/)

[Japanese](http://www.amelt.net/imc/programming/web-app/3215/)

Fetch using `go get`

```
$ go get github.com/RockinTechBomb/gobot
$ go get github.com/ChimeraCoder/anaconda
```

The following environment variables will need to be set with their obvious values

```
src/github.com/RockinTechBomb/gobot/config/config.json
  "consumer_key"
  "consumer_secret"
  "access_token"
  "access_token_secret"
```

Please refer to the connection test
```
$ go run check_connect.go
```

Set the tweets
```
src/github.com/RockinTechBomb/gobot/datas/tweets.json
```

Any select it and tweet

```
$ go run gobot.go
```
