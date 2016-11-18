package qless

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ = fmt.Println

func TestQueue_PushAndPop(t *testing.T) {
	c := newClientFlush()
	q := c.Queue("test_queue")
	j, err := q.Put("class", "data")
	assert.True(t, len(j) > 0)
	assert.NoError(t, err)
	j1, err := q.PopOne()
	assert.NoError(t, err)
	assert.Equal(t, j, j1.JID())
}

func TestQueue_PopNoJobs(t *testing.T) {
	c := newClientFlush()
	q := c.Queue("test_queue")
	job, err := q.PopOne()
	assert.NoError(t, err)
	assert.Nil(t, job)
}

func TestQueue_PopMultiple(t *testing.T) {
	c := newClientFlush()
	q := c.Queue("test_queue")
	for i := 0; i < 10; i++ {
		q.Put("class", "data")
	}
	jobs, err := q.Pop(10)
	assert.NoError(t, err)
	assert.Len(t, jobs, 10)
}

func TestQueue_Length(t *testing.T) {
	c := newClientFlush()
	q := c.Queue("test_queue")
	for i := 0; i < 10; i++ {
		q.Put("class", "data")
	}
	l, err := q.Len()
	assert.NoError(t, err)
	assert.EqualValues(t, 10, l)
}

func TestQueue_JobWithIntervalIsThrottled(t *testing.T) {
	c := newClientFlush()
	q := c.Queue("test_queue")
	_, err := q.Put("class", "data", WithJID("jid-1"), WithInterval(60))
	assert.NoError(t, err)
	j1, err := q.PopOne()
	assert.NoError(t, err)
	assert.Equal(t, "jid-1", j1.JID())
	_, err = j1.Complete()
	assert.NoError(t, err)

	q.Put("class", "data", WithJID("jid-1"), WithInterval(60))
	j1, err = q.PopOne()
	assert.NoError(t, err)
	assert.Nil(t, j1)
}

func TestQueue_JobIsNotReplaced(t *testing.T) {
	c := newClientFlush()
	q := c.Queue("test_queue")
	tl, err := q.PutOrReplace("class", "jid-1", "data")
	assert.NoError(t, err)

	j1, err := q.PopOne()
	assert.NoError(t, err)
	assert.Equal(t, "jid-1", j1.JID())

	tl, err = q.PutOrReplace("class", "jid-1", "data")
	assert.NoError(t, err)
	assert.True(t, tl > 0)
}
