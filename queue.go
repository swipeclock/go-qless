package qless

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var _ = fmt.Sprint("")

type Queue interface {
	Name() string
	Info() QueueInfo
	Jobs(state string, start, count int) ([]string, error)
	CancelAll()
	Pause()
	Resume()
	Put(class string, data interface{}, opt ...putOptionFn) (string, error)
	PopOne() (j Job, err error)
	Pop(count int) ([]Job, error)
	Recur(jid, class string, data interface{}, interval, offset, priority int, tags []string, retries int) (string, error)
	Len() (int64, error)
}

type queue struct {
	name string
	info QueueInfo
	c    *Client
}

func (q *queue) Name() string {
	return q.name
}

func (q *queue) Info() QueueInfo {
	q.c.queueInfo(q.name, &q.info)
	return q.info
}

func (q *queue) Jobs(state string, start, count int) ([]string, error) {
	reply, err := redis.Values(q.c.Do("jobs", timestamp(), state, q.name))
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

// Cancel all jobs in this queue
func (q *queue) CancelAll() {
	for _, state := range jobStates {
		var jids []string
		for {
			jids, _ = q.Jobs(state, 0, 100)
			for _, jid := range jids {
				j, err := q.c.GetRecurringJob(jid)
				if j != nil && err == nil {
					j.Cancel()
				}
			}

			if len(jids) < 100 {
				break
			}
		}
	}
}

func (q *queue) Pause() {
	q.c.Do("pause", timestamp(), q.name)
}

func (q *queue) Resume() {
	q.c.Do("unpause", timestamp(), q.name)
}

type putData struct {
	jid   string
	delay int
	args  []interface{}
}

func newPutData() putData {
	return putData{}
}

func (p *putData) setOptions(opt []putOptionFn) {
	for _, fn := range opt {
		fn(p)
	}

	if p.jid == "" {
		p.jid = generateJID()
	}
}

type putOptionFn func(d *putData)

func putOptionNoOp(*putData) {}

func PutJID(v string) putOptionFn {
	return func(p *putData) {
		p.jid = v
	}
}

func PutDelay(v int) putOptionFn {
	return func(p *putData) {
		p.delay = v
	}
}

func PutPriority(v int) putOptionFn {
	return func(p *putData) {
		p.args = append(p.args, "priority", v)
	}
}

func PutRetries(v int) putOptionFn {
	return func(p *putData) {
		p.args = append(p.args, "retries", v)
	}
}

func PutTags(v []string) putOptionFn {
	if v == nil {
		return putOptionNoOp
	}
	return func(p *putData) {
		p.args = append(p.args, "tags", marshal(v))
	}
}

func PutDepends(v []string) putOptionFn {
	if v == nil {
		return putOptionNoOp
	}
	return func(p *putData) {
		p.args = append(p.args, "depends", marshal(v))
	}
}

func PutResources(v []string) putOptionFn {
	if v == nil {
		return putOptionNoOp
	}
	return func(p *putData) {
		p.args = append(p.args, "resources", marshal(v))
	}
}

// Put enqueues a job to the named queue
func (q *queue) Put(class string, data interface{}, opt ...putOptionFn) (string, error) {
	pd := newPutData()
	pd.setOptions(opt)
	args := []interface{}{"put", timestamp(), "", q.name, pd.jid, class, marshal(data), pd.delay}
	args = append(args, pd.args...)

	return redis.String(q.c.Do(args...))
}

func (q *queue) PopOne() (j Job, err error) {
	var jobs []Job
	if jobs, err = q.Pop(1); err == nil && len(jobs) == 1 {
		j = jobs[0]
	}
	return
}

// Pops a job off the queue.
func (q *queue) Pop(count int) ([]Job, error) {
	if count == 0 {
		count = 1
	}

	reply, err := redis.Bytes(q.c.Do("pop", timestamp(), q.name, workerName(), count))
	if err != nil {
		return nil, err
	}

	if len(reply) == 2 {
		return nil, nil
	}

	var jobsData []jobData
	err = json.Unmarshal(reply, &jobsData)
	if err != nil {
		return nil, err
	}

	jobs := make([]Job, len(jobsData))
	for i, v := range jobsData {
		jobs[i] = &job{jd: &v, c: q.c}
	}

	return jobs, nil
}

// Put a recurring job in this queue
func (q *queue) Recur(jid, class string, data interface{}, interval, offset, priority int, tags []string, retries int) (string, error) {
	if jid == "" {
		jid = generateJID()
	}
	if interval == -1 {
		interval = 0
	}
	if offset == -1 {
		offset = 0
	}
	if priority == -1 {
		priority = 0
	}
	if retries == -1 {
		retries = 5
	}

	return redis.String(q.c.Do(
		"recur", timestamp(), "on", q.name, jid, class,
		data, "interval",
		interval, offset, "priority", priority,
		"tags", marshal(tags), "retries", retries))
}

func (q *queue) Len() (int64, error) {
	reply, err := redis.Int64(q.c.Do("length", timestamp(), q.name))
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	return reply, nil
}
