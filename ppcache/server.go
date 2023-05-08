package ppcache

import "fmt"

type Server struct {
	URL   string
	Name  string
	Pool  *httpPool
	cache *CacheGroup
}

const (
	DefaultGetCachePrefix = "/ppcache/get"
	DefaultGetAllNodes    = "/ppcache/getNodes"
)

func NewServer(url string, name string, hashFunc func([]byte) uint32, cacheSize int, getter func(key string) (value []byte, err error)) *Server {
	return &Server{
		URL:   url,
		Name:  name,
		Pool:  newHttpPool(hashFunc),
		cache: NewGroup(cacheSize, getter),
	}
}

func (s *Server) SetGetter(getter func(key string) (value []byte, err error)) error {
	if getter != nil {
		s.cache.SetGetter(getter)
		return nil
	}
	return fmt.Errorf("[ppCache] getter function can not is nil")
}

func (s *Server) Connection(url string, name string) error {

}
