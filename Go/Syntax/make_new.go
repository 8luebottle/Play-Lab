package main

import (
	"fmt"
	"reflect"
)

type Item struct {
	ID       int64
	Name     string
	Price    float64
	HasColor bool
}

// Builtin function(内建函数): make & new
func main() {
	/*
		new()
			return *T (a pointer to the value's address)
			returns a pointer 返回指针
	*/
	n := new(Item) // zero values
	var nn Item
	var nnn *Item

	ShowDataInfo(n)   // [TYPE: ptr]      *main.Item      &{ID:0 Name: Price:0 HasColor:false}
	ShowDataInfo(nn)  // [TYPE: struct]   main.Item       {ID:0 Name: Price:0 HasColor:false}
	ShowDataInfo(nnn) // [TYPE: ptr]      *main.New       <nil>

	/*
		make()
			used to initialize slice, map, and chan.
			does not return a pointer 不返回指针
	*/
	m := make([]Item, 0)
	mm := make([]Item, 3)
	var mmm []Item
	var mmmm []*Item
	mmmmm := make([]*Item, 3)

	ShowDataInfo(m)     // [TYPE: slice]    []main.Item     []
	ShowDataInfo(mm)    // [TYPE: slice]    []main.Item     [{ID:0 Name: Price:0 HasColor:false} {ID:0 Name: Price:0 HasColor:false} {ID:0 Name: Price:0 HasColor:false}]
	ShowDataInfo(mmm)   // [TYPE: slice]    []main.Item     []
	ShowDataInfo(mmmm)  // [TYPE: slice]    []*main.Item    []
	ShowDataInfo(mmmmm) // [TYPE: slice]    []*main.Item    [<nil> <nil> <nil>]
}

func ShowDataInfo(data interface{}) {
	fmt.Printf("[TYPE: %v] \t %T \t %+v \n", reflect.ValueOf(data).Kind(), data, data)
}
