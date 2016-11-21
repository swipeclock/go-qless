package qless

import (
	"reflect"
	"strings"
)

type RecurringJob struct {
	Tags     StringSlice
	Jid      string
	Retries  int
	Data     interface{}
	Queue    string
	Interval int
	Count    int
	Klass    string
	Priority int

	c *Client
}

func (r *RecurringJob) Update(opts map[string]interface{}) {
	args := []interface{}{"recur", timestamp(), "update", r.Jid}

	vOf := reflect.ValueOf(r).Elem()
	for key, value := range opts {
		key = strings.ToLower(key)
		v := vOf.FieldByName(ucfirst(key))
		if v.IsValid() {
			setv := reflect.ValueOf(value)
			if key == "data" {
				setv = reflect.ValueOf(marshal(value))
			}
			v.Set(setv)
			args = append(args, key, value)
		}
	}

	r.c.Do(args...)
}

func (r *RecurringJob) Cancel() {
	r.c.Do("recur", timestamp(), "off", r.Jid)
}

func (r *RecurringJob) Tag(tags ...interface{}) {
	args := []interface{}{"recur", timestamp(), "tag", r.Jid}
	args = append(args, tags...)
	r.c.Do(args...)
}

func (r *RecurringJob) Untag(tags ...interface{}) {
	args := []interface{}{"recur", timestamp(), "untag", r.Jid}
	args = append(args, tags...)
	r.c.Do(args...)
}
