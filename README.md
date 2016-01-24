gobot
===============
This is twitterbot made with go

## Using

Fetch using `go get`

```
$ go get github.com/RockinTechBomb/gobot
```

Or clone this repository and install dependencies:

```
$ go get github.com/ChimeraCoder/anaconda
```

The following environment variables will need to be set with their obvious values:

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
