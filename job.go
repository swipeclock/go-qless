package qless

import (
	"github.com/garyburd/redigo/redis"
)

var (
	jobStates   = []string{"stalled", "running", "scheduled", "depends", "recurring"}
	finishBytes = []byte(`{"finish":"yes"}`)
)

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
	Data() []byte
	Tags() []string
	History() []History
	Failure() interface{}
	Dependents() []string
	Dependencies() []string

	// operations
	Heartbeat() (bool, error)
	Fail(typ, message string) (bool, error)
	Complete() (string, error)
	CompleteWithNoData() (string, error)
	HeartbeatWithNoData() (bool, error)
	Cancel()
	Retry(delay int) (int, error)
}

//easyjson:json
type job struct {
	jd   *jobData
	data interface{}
	c    *Client
}

func (j *job) JID() string {
	return j.jd.JID
}

func (j *job) Class() string {
	return j.jd.Class
}

func (j *job) State() string {
	return j.jd.State
}

func (j *job) Queue() string {
	return j.jd.Queue
}

func (j *job) Worker() string {
	return j.jd.Worker
}

func (j *job) Tracked() bool {
	return j.jd.Tracked
}

func (j *job) Priority() int {
	return j.jd.Priority
}

func (j *job) Expires() int64 {
	return j.jd.Expires
}

func (j *job) Retries() int {
	return j.jd.Retries
}

func (j *job) Remaining() int {
	return j.jd.Remaining
}

func (j *job) Data() []byte {
	return j.jd.Data
}

func (j *job) Tags() []string {
	return j.jd.Tags
}

func (j *job) History() []History {
	return j.jd.History
}

func (j *job) Failure() interface{} {
	return j.jd.Failure
}

func (j *job) Dependents() []string {
	return j.jd.Dependents
}

func (j *job) Dependencies() []string {
	return j.jd.Dependencies
}

// Move this from it's current queue into another
func (j *job) Move(queueName string) (string, error) {
	return redis.String(j.c.Do("put", timestamp(), queueName, j.jd.JID, j.jd.Class, j.jd.Data, 0))
}

// Fail this job
// return success, error
func (j *job) Fail(typ, message string) (bool, error) {
	return Bool(j.c.Do("fail", timestamp(), j.jd.JID, j.jd.Worker, typ, message, j.jd.Data))
}

// Heartbeats this job
// return success, error
func (j *job) Heartbeat() (bool, error) {
	return Bool(j.c.Do("heartbeat", timestamp(), j.jd.JID, j.jd.Worker, j.jd.Data))
}

// Completes this job
// returns state, error
func (j *job) Complete() (string, error) {
	return redis.String(j.c.Do("complete", timestamp(), j.jd.JID, j.jd.Worker, j.jd.Queue, j.jd.Data))
}

//for big job, save memory in redis
func (j *job) CompleteWithNoData() (string, error) {
	return redis.String(j.c.Do("complete", timestamp(), j.jd.JID, j.jd.Worker, j.jd.Queue, finishBytes))
}

func (j *job) HeartbeatWithNoData() (bool, error) {
	return Bool(j.c.Do("heartbeat", timestamp(), j.jd.JID, j.jd.Worker))
}

// Cancels this job
func (j *job) Cancel() {
	j.c.Do("cancel", timestamp(), j.jd.JID)
}

// Track this job
func (j *job) Track() (bool, error) {
	return Bool(j.c.Do("track", timestamp(), "track", j.jd.JID))
}

// Untrack this job
func (j *job) Untrack() (bool, error) {
	return Bool(j.c.Do("track", timestamp(), "untrack", j.jd.JID))
}

func (j *job) AddTags(tags ...interface{}) (string, error) {
	args := []interface{}{"tag", timestamp(), "add", j.jd.JID}
	args = append(args, tags...)
	return redis.String(j.c.Do(args...))
}

func (j *job) RemoveTags(tags ...interface{}) (string, error) {
	args := []interface{}{"tag", timestamp(), "remove", j.jd.JID}
	args = append(args, tags...)
	return redis.String(j.c.Do(args...))
}

func (j *job) Retry(delay int) (int, error) {
	return redis.Int(j.c.Do("retry", timestamp(), j.jd.JID, j.jd.Queue, j.jd.Worker, delay))
}

func (j *job) Depend(jids ...interface{}) (string, error) {
	args := []interface{}{"depends", timestamp(), j.jd.JID, "on"}
	args = append(args, jids...)
	return redis.String(j.c.Do(args...))
}

func (j *job) Undepend(jids ...interface{}) (string, error) {
	args := []interface{}{"depends", timestamp(), j.jd.JID, "off"}
	args = append(args, jids...)
	return redis.String(j.c.Do(args...))
}
