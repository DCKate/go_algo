package main

import (
	"fmt"

	"../go_algo/sort"
)

//SortEle my sort element
type SortEle struct {
	Value int64
}

//GetCompareValue Inplement the interface
func (th SortEle) GetCompareValue() int64 {
	return th.Value
}

func makeData() []sort.SortInterf {
	a := SortEle{Value: 23}
	b := SortEle{Value: 10}
	c := SortEle{Value: 2}
	d := SortEle{Value: 3}
	e := SortEle{Value: 18}
	f := SortEle{Value: 5}

	var arr []sort.SortInterf
	return append(arr, a, b, c, d, e, f)
}

func testBuble() {
	aaa := makeData()
	ooo := sort.GoBubleSort(aaa)
	fmt.Printf("%v\n", ooo)
}

func main() {

}
