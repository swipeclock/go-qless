package qless

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

type TaggedReply struct {
	Total int
	Jobs  StringSlice
}

type Client struct {
	conn redis.Conn
	host string
	port string
	db   int

	events *Events
	lua    *redis.Script
}

func Dial(host, port string, db int) (*Client, error) {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port), redis.DialDatabase(db))
	if err != nil {
		return nil, err
	}

	return &Client{
		host: host,
		port: port,
		db:   db,
		lua:  redis.NewScript(0, qlessLua),
		conn: conn,
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) Events() *Events {
	if c.events != nil {
		return c.events
	}
	c.events = NewEvents(c.host, c.port, c.db)
	return c.events
}

func scriptReply(reply interface{}, err error) (interface{}, error) {
	return reply, parseError(err)
}

func (c *Client) Do(args ...interface{}) (interface{}, error) {
	return scriptReply(c.lua.Do(c.conn, args...))
}

func (c *Client) Queue(name string) Queue {
	return &queue{name: name, c: c}
}

func (c *Client) Queues() (queues []Queue, err error) {
	args := []interface{}{"queues", timestamp()}

	byts, err := redis.Bytes(c.Do(args...))
	if err != nil {
		return nil, err
	}

	var infos []QueueInfo
	err = json.Unmarshal(byts, &infos)
	if err != nil {
		return nil, err
	}

	queues = make([]Queue, len(infos))
	for i, info := range infos {
		queues[i] = &queue{name: info.Name, info: info, c: c}
	}
	return queues, err
}

func (c *Client) queueInfo(name string, info *QueueInfo) error {
	args := []interface{}{"queues", timestamp(), name}

	byts, err := redis.Bytes(c.Do(args...))
	if err != nil {
		return err
	}

	return json.Unmarshal(byts, info)
}

// Track the jid
func (c *Client) Track(jid string) (bool, error) {
	return Bool(c.Do("track", timestamp(), "track", jid, ""))
}

// Untrack the jid
func (c *Client) Untrack(jid string) (bool, error) {
	return Bool(c.Do("track", timestamp(), 0, "untrack", jid))
}

// Returns all the tracked jobs
func (c *Client) Tracked() (string, error) {
	return redis.String(c.Do("track", timestamp()))
}

func (c *Client) Get(jid string) (interface{}, error) {
	job, err := c.GetJob(jid)
	if err == redis.ErrNil {
		rjob, err := c.GetRecurringJob(jid)
		return rjob, err
	}
	return job, err
}

func (c *Client) GetJob(jid string) (Job, error) {
	byts, err := redis.Bytes(c.Do("get", timestamp(), jid))
	if err != nil {
		return nil, err
	}

	var d jobData
	err = json.Unmarshal(byts, d)
	if err != nil {
		return nil, err
	}
	return &job{d: &d, c: c}, err
}

func (c *Client) GetRecurringJob(jid string) (*RecurringJob, error) {
	byts, err := redis.Bytes(c.Do("recur", timestamp(), "get", jid))
	if err != nil {
		return nil, err
	}

	job := NewRecurringJob(c)
	err = json.Unmarshal(byts, job)
	if err != nil {
		return nil, err
	}
	return job, err
}

func (c *Client) Completed(start, count int) ([]string, error) {
	reply, err := redis.Values(c.Do("jobs", timestamp(), "complete"))
	if err != nil {
		return nil, err
	}

	ret := []string{}
	for _, val := range reply {
		s, _ := redis.String(val, err)
		ret = append(ret, s)
	}
	return ret, err
}

func (c *Client) Tagged(tag string, start, count int) (*TaggedReply, error) {
	byts, err := redis.Bytes(c.Do("tag", timestamp(), "get", tag, start, count))
	if err != nil {
		return nil, err
	}

	t := &TaggedReply{}
	err = json.Unmarshal(byts, t)
	return t, err
}

func (c *Client) GetConfig(option string) (string, error) {
	interf, err := c.Do("config.get", timestamp(), option)
	if err != nil {
		return "", err
	}

	var contentStr string
	switch interf.(type) {
	case []uint8:
		contentStr, err = redis.String(interf, nil)
	case int64:
		var contentInt64 int64
		contentInt64, err = redis.Int64(interf, nil)
		if err == nil {
			contentStr = strconv.Itoa(int(contentInt64))
		}
	default:
		err = errors.New("The redis return type is not []uint8 or int64")
	}
	if err != nil {
		return "", err
	}

	return contentStr, err
}

func (c *Client) SetConfig(option string, value interface{}) {
	intf, err := c.Do("config.set", timestamp(), option, value)
	if err != nil {
		fmt.Println("setconfig, c.Do fail. interface:", intf, " err:", err)
	}
}

func (c *Client) UnsetConfig(option string) {
	c.Do("config.unset", timestamp(), option)
}
