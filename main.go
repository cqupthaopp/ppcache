package main

import (
	"errors"
	"fmt"
	"ppCache/ppcache"
)

/*
Easy Example
*/

func getter(key string) ([]byte, error) {
	if key != "Hao_pp" {
		return nil, errors.New("No cache Hit")
	}
	return ([]byte)(key + "Hao_pp"), nil
}

func Test(g *ppcache.CacheGroup) {
	fmt.Println(g.Len())
	fmt.Println(g.UsedMemory())
	fmt.Println(g.MaxMemory())
	fmt.Println(g.Get("Hao_pp"))
	fmt.Println(g.Get("Hao_pp1"))
}

func main() {

	g := ppcache.NewGroup(21, nil)

	g.Add("Hao_pp", []byte("!234"))

	Test(g)

	g.Add("Hao_pp1", []byte("!234"))

	Test(g)

}
