package qless

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"time"
	"unicode"
	"unicode/utf8"
)

var uuidRand = rand.Reader

func generateUUID() (string, error) {
	var uuid [16]byte

	if n, err := io.ReadFull(uuidRand, uuid[:]); err != nil {
		return "", errors.New("read failed: " + err.Error())
	} else if n != len(uuid) {
		return "", errors.New("read failed")
	}

	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40

	var dst [32]byte
	hex.Encode(dst[:], uuid[:])
	return string(dst[:]), nil
}

// returns a timestamp used in LUA calls
func timestamp() int64 {
	return time.Now().Unix()
}

// makes the first character of a string upper case
func ucfirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}
