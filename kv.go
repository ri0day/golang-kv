package kv

import (
	"github.com/ri0day/golang-kv/bolt"
)

func Bolt(path string) Bucket {
	return bolt.Open(path)
}
