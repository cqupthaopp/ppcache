package ppcache

import "hash/crc32"

type httpServer struct {
	Url  string
	Name string
}

type httpPool struct {
	Servers map[uint32]*httpServer   //记录节点
	hash    func(data []byte) uint32 //哈希
}

func newHttpPool(hashFunc func(data []byte) uint32) *httpPool {

	res := &httpPool{
		Servers: make(map[uint32]*httpServer),
		hash:    hashFunc,
	}

	if res.hash == nil {
		res.hash = crc32.ChecksumIEEE
	}

	return res

}

func (p *httpPool)Push( url string,name string){
	p.Servers[   ]
}
