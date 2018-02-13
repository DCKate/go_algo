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

func makeSmallData() []sort.SortInterf {
	// a := SortEle{Value: 23}
	b := SortEle{Value: 10}
	c := SortEle{Value: 2}
	d := SortEle{Value: 3}
	e := SortEle{Value: 18}
	f := SortEle{Value: 5}
	var arr []sort.SortInterf
	return append(arr, b, c, d, e, f)
}
func makeData() []sort.SortInterf {
	a := SortEle{Value: 23}
	b := SortEle{Value: 10}
	c := SortEle{Value: 2}
	d := SortEle{Value: 3}
	e := SortEle{Value: 18}
	f := SortEle{Value: 5}
	g := SortEle{Value: 70}
	h := SortEle{Value: 1}
	i := SortEle{Value: 15}
	j := SortEle{Value: 29}
	k := SortEle{Value: 37}
	l := SortEle{Value: 8}
	m := SortEle{Value: 70}
	n := SortEle{Value: 1}
	o := SortEle{Value: 15}
	p := SortEle{Value: 29}
	q := SortEle{Value: 37}
	r := SortEle{Value: 8}
	s := SortEle{Value: 2}
	t := SortEle{Value: 58}
	u := SortEle{Value: 62}
	v := SortEle{Value: 5}
	w := SortEle{Value: 10}
	x := SortEle{Value: 11}
	y := SortEle{Value: 33}
	z := SortEle{Value: 3}
	var arr []sort.SortInterf
	return append(arr, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z)
}

func testBuble() {
	aaa := makeData()
	ooo := sort.GoBubleSort(aaa)
	fmt.Printf("%v\n", ooo)
}

func testSelect() {
	aaa := makeData()
	ooo := sort.GoSelectSort(aaa)
	fmt.Printf("%v\n", ooo)
}

func testQuick() {
	aaa := makeData()
	ooo := sort.GoQuickSort(aaa)
	fmt.Printf("%v\n", ooo)
}

func testInsert() {
	aaa := makeSmallData()
	ooo := sort.GoInsertSort(aaa)
	fmt.Printf("%v\n", ooo)
}

func testMerge() {
	aaa := makeData()
	ooo := sort.GoMergeSort(aaa)
	fmt.Printf("%v\n", ooo)
}

func main() {
	// testBuble()
	// testSelect()
	// testQuick()
	testInsert()
	// testMerge()
}
