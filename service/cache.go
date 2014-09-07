package service

import (
	"github.com/feedlabs/feedify/memcache"
)

func NewCache() *memcache.MemcacheClient {
	return memcache.NewMemcacheClient()
}
