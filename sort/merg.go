package sort

func mergeArray(a []SortInterf, b []SortInterf) []SortInterf {
	var odr []SortInterf
	for len(a) > 0 && len(b) > 0 {
		if a[0].GetCompareValue() <= b[0].GetCompareValue() {
			odr = append(odr, a[0])
			a = a[1:]
			continue
		}

		odr = append(odr, b[0])
		b = b[1:]

	}
	if len(a) > 0 {
		for _, e := range a {
			odr = append(odr, e)
		}
	}
	if len(b) > 0 {
		for _, e := range b {
			odr = append(odr, e)
		}
	}
	return odr
}

func recurrMerge(array []SortInterf) []SortInterf {
	al := len(array)
	if al == 1 {
		return array
	}
	splii := al / 2
	left := recurrMerge(array[:splii])
	right := recurrMerge(array[splii:])
	return mergeArray(left, right)
}

//GoMergeSort input array and ouput he sorted array
func GoMergeSort(array []SortInterf) []SortInterf {
	return recurrMerge(array)
}
