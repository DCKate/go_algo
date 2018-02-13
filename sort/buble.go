package sort

//SortInterf the interface need to be implement for buble sort
type SortInterf interface {
	GetCompareValue() int64
}

func bubleArraySwap(array []SortInterf) []SortInterf {
	for ii := len(array) - 1; ii > 0; ii-- {
		if array[ii].GetCompareValue() < array[ii-1].GetCompareValue() {
			array[ii-1], array[ii] = array[ii], array[ii-1]
		}
	}
	return array
}

//GoBubleSort input the buble interface array and output the dorted array
func GoBubleSort(array []SortInterf) []SortInterf {
	tlen := len(array)
	var final []SortInterf
	oparray := array
	for len(final) < tlen {
		barr := bubleArraySwap(oparray)
		final = append(final, barr[0])
		if len(barr) > 1 {
			oparray = barr[1:]
		} else {
			break
		}
	}
	return final
}
