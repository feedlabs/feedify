package service

import (
	"github.com/feedlabs/feedify/memcache"
)

func NewMemcache() *memcache.MemcacheClient {
	return memcache.NewMemcacheClient()
}
