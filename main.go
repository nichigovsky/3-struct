package main

import "time"

type Bin struct {
	id string
	private bool
	createdAt time.Time
	name string
}

type BinList = []Bin

func createBin() *Bin {
	return &Bin{
		createdAt: time.Now(),
	}
}

func createList() BinList {
	sli := make(BinList, 0)

	return sli
}

func main() {

}
