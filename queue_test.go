package qless

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
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
	fmt.Println(q.Info())
	jobs, err := q.Pop(10)
	fmt.Println(q.Info())
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
