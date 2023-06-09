package utils

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

var entropy = ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)

func CreateULID() string {
	return ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
}
