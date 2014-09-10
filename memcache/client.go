package memcache

import (
	"github.com/bradfitz/gomemcache/memcache"

	"github.com/feedlabs/feedify/config"
)

type MemcacheClient struct {
	host	string
	port	string
	conn	*memcache.Client
}

func (m *MemcacheClient) Connect() {
	m.conn = memcache.New(m.host + ":" + m.port)
}

func (m *MemcacheClient) Set(key string, value string) {
	m.conn.Set(&memcache.Item{Key: key, Value: []byte(value)})
}

func (m *MemcacheClient) Get(key string) (value string, err error) {
	_item, err := m.conn.Get(key)
	return string(_item.Value), err
}

func (m *MemcacheClient) GetMulti(keys []string) map[string]string {
	_items, _ := m.conn.GetMulti(keys)

	items := make(map[string]string, len(_items))

	for _, key := range(keys)  {
		items[key] = string(_items[key].Value)
	}

	return items
}

func (m *MemcacheClient) Delete(key string) error {
	return m.conn.Delete(key)
}

func (m *MemcacheClient) DeleteAll() error {
	return m.conn.DeleteAll()
}

func NewMemcacheClient() *MemcacheClient {
	host := config.GetConfigKey("memcache::host")
	port := config.GetConfigKey("memcache::port")

	return &MemcacheClient{host, port, nil}
}
