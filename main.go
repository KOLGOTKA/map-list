package main

import (
	"map-list/storage"
	"map-list/storage/list"
	"map-list/storage/mp"
)

func main() {
	var Liist storage.Storage = &list.List{}
	var Mpp storage.Storage = &mp.Mp{}
	Liist.Add(5)
	Liist.Add(6)
	Liist.Add(7)
	Liist.Add(8)
	Mpp.Add(5)
	Mpp.Add(6)
	Mpp.Add(7)
	Mpp.Add(8)
	Liist.Print()
	Mpp.Print()

}
