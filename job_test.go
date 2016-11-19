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

func TestJob_Complete(t *testing.T) {
	c := newClientFlush()
	q := c.Queue("test_queue")
	q.Put("class", "data", WithJID("jobTestDEF"))

	workerNameStr = "worker-1"
	j1, err := q.PopOne()
	assert.NoError(t, err)
	res, err := j1.Complete()
	assert.NoError(t, err)
	assert.Equal(t, "complete", res)
}

func TestJob_Retry(t *testing.T) {
	c := newClientFlush()
	q := c.Queue("test_queue")
	q.Put("class", "data", WithJID("jobTestDEF"))

	workerNameStr = "worker-1"
	j1, err := q.PopOne()
	assert.NoError(t, err)
	remain, err := j1.Retry(0)
	assert.NoError(t, err)
	assert.Equal(t, 4, remain)
}

func TestJob_RetryRespectsPutArgument(t *testing.T) {
	c := newClientFlush()
	q := c.Queue("test_queue")
	q.Put("class", "data", WithJID("jobTestDEF"), WithRetries(1))

	workerNameStr = "worker-1"
	j1, err := q.PopOne()
	assert.NoError(t, err)
	remain, err := j1.Retry(0)
	assert.NoError(t, err)
	assert.Equal(t, 0, remain)
}

func TestJob_RetryReturnsNegativeOneWhenNoMoreAvailable(t *testing.T) {
	c := newClientFlush()
	q := c.Queue("test_queue")
	q.Put("class", "data", WithJID("jobTestDEF"), WithRetries(0))

	workerNameStr = "worker-1"
	j1, err := q.PopOne()
	assert.NoError(t, err)
	remain, err := j1.Retry(0)
	assert.NoError(t, err)
	assert.Equal(t, -1, remain)
}

func TestJob_Cancel(t *testing.T) {
	c := newClientFlush()
	q := c.Queue("test_queue")
	q.Put("class", "data", WithJID("jobTestDEF"))

	workerNameStr = "worker-1"
	j1, err := q.PopOne()
	assert.NoError(t, err)
	j1.Cancel()
	j1, err = q.PopOne()
	assert.NoError(t, err)
	assert.Nil(t, j1)
}
