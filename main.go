package main

import "time"

type Bin struct {
	id string
	private bool
	createdAt time.Time
	name string
}

type BinList = []Bin

func createList() {
	sli := make([]BinList, 2)
}

func main() {

}
