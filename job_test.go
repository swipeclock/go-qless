package qless

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJob_HeartbeatForInvalidReturnsError(t *testing.T) {
	c := newClientFlush()
	c.SetConfig("heartbeat", -10)
	c.SetConfig("grace-period", 0)

	q := c.Queue("test_queue")
	q.Put("class", "data", WithJID("jobTestDEF"))

	workerNameStr = "worker-1"
	j1, err := q.PopOne()
	assert.NoError(t, err)

	workerNameStr = "worker-2"
	q.PopOne()

	_, err = j1.Heartbeat()
	assert.Error(t, err)
	assert.True(t, IsJobLost(err))
}

func TestJob_TTL(t *testing.T) {
	c := newClientFlush()
	q := c.Queue("test_queue")
	q.Put("class", "data", WithJID("jobTestDEF"))

	workerNameStr = "worker-1"
	j1, err := q.PopOne()
	assert.NoError(t, err)
	assert.Condition(t, func() bool { return j1.TTL() > 55 })
}