package sort

func findInserIndex(tar SortInterf, array []SortInterf) int {
	if len(array) == 0 {
		return 0
	}
	if len(array) == 1 {
		if tar.GetCompareValue() <= array[0].GetCompareValue() {
			return 0
		}
		return 1
	}
	refi := (len(array) - 1) / 2
	ref := array[refi]
	if tar.GetCompareValue() <= ref.GetCompareValue() {
		ff := findInserIndex(tar, array[:refi])
		return ff
	}
	final := refi + findInserIndex(tar, array[refi+1:]) + 1
	return final
}

func insertArray(tar SortInterf, array []SortInterf) []SortInterf {
	var base []SortInterf
	ii := findInserIndex(tar, array)
	if ii == 0 {
		base = append(base, tar)
	} else {
		for _, e := range array[:ii] {
			base = append(base, e)
		}
		base = append(base, tar)
	}

	for _, e := range array[ii:] {
		base = append(base, e)
	}
	return base
}

//GoInsertSort input array and ouput he sorted array
func GoInsertSort(array []SortInterf) []SortInterf {
	var order []SortInterf
	for _, e := range array {
		order = insertArray(e, order)
	}
	return order
}
