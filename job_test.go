package qless

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

var _ = fmt.Println

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
	assert.Condition(t, func() bool {
		return j1.TTL() > 55
	})
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

func TestJob_Fail(t *testing.T) {
	c := newClientFlush()
	q := c.Queue("test_queue")
	q.Put("class", "data", WithJID("jobTestDEF"))

	workerNameStr = "worker-1"
	j1, err := q.PopOne()
	assert.NoError(t, err)
	res, err := j1.Fail("type", "message")
	assert.NoError(t, err)
	assert.Equal(t, true, res)
	j1, err = c.GetJob("jobTestDEF")
	assert.NoError(t, err)
	assert.NotNil(t, j1)
	failure := j1.Failure()
	failure.When = 0
	assert.Equal(t, &Failure{Group: "type", Message: "message", Worker: workerNameStr}, failure)
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

func TestJob_UnmarshalDataOfStruct(t *testing.T) {

	type testData struct {
		String      string
		Int         int
		StringArray []string
	}

	inData := testData{String: "a string", Int: 0xbadf00d, StringArray: []string{"the", "fox"}}
	c := newClientFlush()
	q := c.Queue("test_queue")
	q.Put("class", inData, WithJID("jobTestDEF"))
	workerNameStr = "worker-1"
	j1, err := q.PopOne()
	assert.NoError(t, err)

	var outData testData
	err = j1.UnmarshalData(&outData)
	assert.NoError(t, err)
	assert.Equal(t, inData, outData)
}

func TestJob_UnmarshalDataOfString(t *testing.T) {
	inData := "foo bar"
	c := newClientFlush()
	q := c.Queue("test_queue")
	q.Put("class", inData, WithJID("jobTestDEF"))
	workerNameStr = "worker-1"
	j1, err := q.PopOne()
	assert.NoError(t, err)

	var outData string
	err = j1.UnmarshalData(&outData)
	assert.NoError(t, err)
	assert.Equal(t, inData, outData)
}
