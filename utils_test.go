package qless

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUUID(t *testing.T) {
	b := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	uuidRand = bytes.NewBuffer(b)
	v, err := generateUUID()
	assert.NoError(t, err)
	assert.Equal(t, "000102030405460788090a0b0c0d0e0f", v)
}
