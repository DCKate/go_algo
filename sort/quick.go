package sort

func combineArray(ref []SortInterf, a []SortInterf, b []SortInterf) []SortInterf {
	for _, r := range ref {
		a = append(a, r)
	}
	for _, e := range b {
		a = append(a, e)
	}
	return a
}

func quickArraySwap(array []SortInterf) ([]SortInterf, []SortInterf, []SortInterf) {
	refer := array[0]
	var reArr []SortInterf
	var lsArr []SortInterf
	var bgArr []SortInterf
	for _, e := range array {
		if e.GetCompareValue() < refer.GetCompareValue() {
			lsArr = append(lsArr, e)
		} else if e.GetCompareValue() > refer.GetCompareValue() {
			bgArr = append(bgArr, e)
		} else {
			reArr = append(reArr, e)
		}
	}
	return reArr, lsArr, bgArr
}

func recurrQuick(array []SortInterf) []SortInterf {
	ref, ls, bg := quickArraySwap(array)
	ll := len(ls)
	lb := len(bg)
	if ll > 1 {
		// fmt.Printf("Left %v\n", ls)
		ls = recurrQuick(ls)
	}
	if lb > 1 {
		// fmt.Printf("Right %v\n", bg)
		bg = recurrQuick(bg)
	}
	cc := combineArray(ref, ls, bg)
	// fmt.Printf("Combine %v\n", cc)
	return cc
}

//GoQuickSort input array and ouput he sorted array
func GoQuickSort(array []SortInterf) []SortInterf {
	return recurrQuick(array)
}
