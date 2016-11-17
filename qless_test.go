package qless

import (
	"os"
	"strconv"
	"testing"
)

var (
	redisHost string
	redisPort string
	redisDB   int
)

func init() {
	redisHost = os.Getenv("REDIS_HOST")
	if redisHost == "" {
		panic("invalid REDIS_HOST")
	}
	redisPort = os.Getenv("REDIS_PORT")

	if redisPort == "" {
		panic("invalid REDIS_PORT")
	}

	if db, err := strconv.Atoi(os.Getenv("REDIS_DB")); err != nil {
		panic("invalid REDIS_DB")
	} else {
		redisDB = db
	}
}

func newClient() *Client {
	c, err := Dial(redisHost, redisPort, redisDB)
	if err != nil {
		panic(err.Error())
	}

	return c
}

func newClientFlush() *Client {
	c, err := Dial(redisHost, redisPort, redisDB)
	if err != nil {
		panic(err.Error())
	}

	c.conn.Do("FLUSHDB")

	return c
}

func flushDB() {
	c := newClient()
	defer c.Close()
	c.conn.Do("FLUSHDB")
}

func TestMain(m *testing.M) {
	flushDB()
	os.Exit(m.Run())
}
