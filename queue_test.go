package qless

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestQueue_PushAndPop(t *testing.T) {
	c := newClient()
	q := c.Queue("test_queue")
	j, err := q.Put("class", "data", PutJID("jid"))
	assert.NoError(t, err)
	fmt.Println(j)
	job, err := q.PopOne()
	assert.NoError(t, err)
	assert.NotNil(t, job)
	fmt.Println(job.Data())
}
