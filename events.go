package qless

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type Events struct {
	conn redis.Conn
	ch   chan interface{}
	psc  redis.PubSubConn
	host string
	port string
	db   int

	listening bool
}

var channels = []string{"ql:log", "canceled", "completed", "failed", "popped", "stalled", "put", "track", "untrack"}

func NewEvents(host, port string, db int) *Events {
	return &Events{host: host, port: port, db: db}
}

func (e *Events) Listen() (chan interface{}, error) {
	var err error
	e.conn, err = redis.Dial("tcp", fmt.Sprintf("%s:%s", e.host, e.port), redis.DialDatabase(e.db))
	if err != nil {
		return nil, err
	}

	e.psc = redis.PubSubConn{Conn: e.conn}
	ch := make(chan interface{}, 1)

	if err := e.psc.Subscribe(channels...); err != nil {
		close(ch)
		e.close()
		return nil, err
	}

	go func() {
		for {
			val := e.psc.Receive()
			if v, ok := val.(error); ok {
				ch <- v
				close(ch)
				break
			}
			ch <- val
		}
	}()

	e.listening = true

	return ch, nil
}

func (e *Events) Unsubscribe() {
	if e.listening {
		e.close()
		e.listening = false
	}
}

func (e *Events) close() {
	e.psc.Unsubscribe()
	e.psc.Close()
	e.psc = redis.PubSubConn{}
}
