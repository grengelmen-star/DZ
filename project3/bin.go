package main

import "time"

type Bin struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}
type BinList []Bin

func NewBin(id, name string, private bool) Bin {
	return Bin{
		ID:        id,
		Name:      name,
		Private:   private,
		CreatedAt: time.Now(),
	}
}
func NewBinList() BinList {
	return BinList{}
}
func (a *BinList) Add(b Bin) {
	*a = append(*a, b)
}
