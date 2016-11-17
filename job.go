package qless

import (
	"github.com/garyburd/redigo/redis"
	"reflect"
	"strings"
)

var (
	JOBSTATES   = []string{"stalled", "running", "scheduled", "depends", "recurring"}
	finishBytes = []byte(`{"finish":"yes"}`)
)

type History struct {
	When   int64
	Q      string
	What   string
	Worker string
}

type Job interface {
	JID() string
	Class() string
	State() string
	Queue() string
	Worker() string
	Tracked() bool
	Priority() int
	Expires() int64
	Retries() int
	Remaining() int
	Data() interface{}
	Tags() []string
	History() []History
	Failure() interface{}
	Dependents() []string
	Dependencies() interface{}

	// operations
	Fail(typ, message string) (bool, error)
	CompleteWithNoData() (string, error)
	HeartbeatWithNoData() (bool, error)
}

//easyjson:json
type job struct {
	data *jobData
	c    *Client
}

func (j *job) JID() string {
	return j.data.Jid
}
func (j *job) Class() string {
	return j.data.Klass
}
func (j *job) State() string {
	return j.data.State
}
func (j *job) Queue() string {
	return j.data.Queue
}
func (j *job) Worker() string {
	return j.data.Worker
}
func (j *job) Tracked() bool {
	return j.data.Tracked
}
func (j *job) Priority() int {
	return j.data.Priority
}
func (j *job) Expires() int64 {
	return j.data.Expires
}
func (j *job) Retries() int {
	return j.data.Retries
}
func (j *job) Remaining() int {
	return j.data.Remaining
}
func (j *job) Data() interface{} {
	return j.data.Data
}
func (j *job) Tags() []string {
	return j.data.Tags
}
func (j *job) History() []History {
	return j.data.History
}
func (j *job) Failure() interface{} {
	return j.data.Failure
}
func (j *job) Dependents() []string {
	return j.data.Dependents
}
func (j *job) Dependencies() interface{} {
	return j.data.Dependencies
}

// Move this from it's current queue into another
func (j *job) Move(queueName string) (string, error) {
	return redis.String(j.c.Do("put", timestamp(), queueName, j.data.Jid, j.data.Klass, marshal(j.data.Data), 0))
}

// Fail this job
// return success, error
func (j *job) Fail(typ, message string) (bool, error) {
	return Bool(j.c.Do("fail", timestamp(), j.data.Jid, j.data.Worker, typ, message, marshal(j.data.Data)))
}

// Heartbeats this job
// return success, error
func (j *job) Heartbeat() (bool, error) {
	return Bool(j.c.Do("heartbeat", timestamp(), j.data.Jid, j.data.Worker, marshal(j.data.Data)))
}

// Completes this job
// returns state, error
func (j *job) Complete() (string, error) {
	return redis.String(j.c.Do("complete", timestamp(), j.data.Jid, j.data.Worker, j.data.Queue, marshal(j.data.Data)))
}

//for big job, save memory in redis
func (j *job) CompleteWithNoData() (string, error) {
	return redis.String(j.c.Do("complete", timestamp(), j.data.Jid, j.data.Worker, j.data.Queue, finishBytes))
}

func (j *job) HeartbeatWithNoData() (bool, error) {
	return Bool(j.c.Do("heartbeat", timestamp(), j.data.Jid, j.data.Worker))
}

// Cancels this job
func (j *job) Cancel() {
	j.c.Do("cancel", timestamp(), j.data.Jid)
}

// Track this job
func (j *job) Track() (bool, error) {
	return Bool(j.c.Do("track", timestamp(), "track", j.data.Jid))
}

// Untrack this job
func (j *job) Untrack() (bool, error) {
	return Bool(j.c.Do("track", timestamp(), "untrack", j.data.Jid))
}

func (j *job) Tag(tags ...interface{}) (string, error) {
	args := []interface{}{"tag", timestamp(), "add", j.data.Jid}
	args = append(args, tags...)
	return redis.String(j.c.Do(args...))
}

func (j *job) Untag(tags ...interface{}) (string, error) {
	args := []interface{}{"tag", timestamp(), "remove", j.data.Jid}
	args = append(args, tags...)
	return redis.String(j.c.Do(args...))
}

func (j *job) Retry(delay int) (int, error) {
	return redis.Int(j.c.Do("retry", timestamp(), j.data.Jid, j.data.Queue, j.data.Worker, delay))
}

func (j *job) Depend(jids ...interface{}) (string, error) {
	args := []interface{}{"depends", timestamp(), j.data.Jid, "on"}
	args = append(args, jids...)
	return redis.String(j.c.Do(args...))
}

func (j *job) Undepend(jids ...interface{}) (string, error) {
	args := []interface{}{"depends", timestamp(), j.data.Jid, "off"}
	args = append(args, jids...)
	return redis.String(j.c.Do(args...))
}

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

	cli *Client
}

func NewRecurringJob(cli *Client) *RecurringJob {
	return &RecurringJob{cli: cli}
}

// example: job.Update(map[string]interface{}{"priority": 5})
// options:
//   priority int
//   retries int
//   interval int
//   data interface{}
//   klass string
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

	r.cli.Do(args...)
}

func (r *RecurringJob) Cancel() {
	r.cli.Do("recur", timestamp(), "off", r.Jid)
}

func (r *RecurringJob) Tag(tags ...interface{}) {
	args := []interface{}{"recur", timestamp(), "tag", r.Jid}
	args = append(args, tags...)
	r.cli.Do(args...)
}

func (r *RecurringJob) Untag(tags ...interface{}) {
	args := []interface{}{"recur", timestamp(), "untag", r.Jid}
	args = append(args, tags...)
	r.cli.Do(args...)
}
