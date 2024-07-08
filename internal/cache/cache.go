package cache

import "time"

type Value interface{}

type Cache struct {
	Expiry time.Time
	Value  Value
}
