package qless

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDial(t *testing.T) {
	c := newClient()
	c.SetConfig("test", "hello world")
	v, _ := c.GetConfig("test")
	fmt.Println(v)
	assert.NotNil(t, c)
	c.conn.Do("FLUSHDB")
}
