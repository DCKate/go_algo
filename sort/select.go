package sort

func selectArraySwap(array []SortInterf) []SortInterf {
	smallidx := 0
	for ii := 0; ii < len(array); ii++ {
		if array[smallidx].GetCompareValue() > array[ii].GetCompareValue() {
			smallidx = ii
		}
	}
	array[0], array[smallidx] = array[smallidx], array[0]
	return array
}

//GoSelectSort input the buble interface array and output the dorted array
func GoSelectSort(array []SortInterf) []SortInterf {
	tlen := len(array)
	var final []SortInterf
	oparray := array
	for len(final) < tlen {
		barr := selectArraySwap(oparray)
		final = append(final, barr[0])
		if len(barr) > 1 {
			oparray = barr[1:]
		} else {
			break
		}
	}
	return final
}
