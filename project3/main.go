package main

import "fmt"

func main() {
	list := NewBinList()
	list.Add(NewBin("1", "data.json", false))
	fmt.Println(list)
}
