package tools

import (
	crand "crypto/rand"
	"fmt"
	mrand "math/rand"
	"time"
)

func RandomUUID() string {
	//source:github.com/sluu99/uuid
	var x [16]byte

	length := 16
	n, err := crand.Read(x[:])

	if n != length || err != nil {
		mrand.Seed(time.Now().UnixNano())

		for length > 0 {
			length--
			x[length] = byte(mrand.Int31n(256))
		}
	}

	x[6] = (x[6] & 0x0F) | 0x40
	x[8] = (x[8] & 0x3F) | 0x80

	return fmt.Sprintf("%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x",
		x[0], x[1], x[2], x[3], x[4],
		x[5], x[6],
		x[7], x[8],
		x[9], x[10], x[11], x[12], x[13], x[14], x[15])
}
