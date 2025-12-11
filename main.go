package main

import "time"

type Bin struct {
	Id string `json:"id"`
	Private bool `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name string `json:"name"`
}

type BinList = []Bin

func createBin() *Bin {
	return &Bin{
		CreatedAt: time.Now(),
	}
}

func createList() BinList {
	sli := make(BinList, 0)

	return sli
}

func main() {

}
