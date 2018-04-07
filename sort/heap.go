package sort

import (
	"fmt"
)

type heapNode struct {
	Com   SortInterf
	pre   int
	left  int
	right int
}

func decideNodeUp(id int, t []heapNode) []heapNode {
	nid := id
	cid := t[nid].pre
	for cid >= 0 {
		com := t[cid]
		if t[nid].Com.GetCompareValue() > t[cid].Com.GetCompareValue() {
			t[nid].Com, t[cid].Com = t[cid].Com, t[nid].Com
			nid = cid
			cid = com.pre
		} else {
			break
		}
	}
	return t
}

func decideNodeDown(t []heapNode) []heapNode {
	nid := 0
	lid := t[nid].left
	rid := t[nid].right
	for lid >= 0 || rid >= 0 {
		if rid == -1 || (t[lid].Com.GetCompareValue() >= t[rid].Com.GetCompareValue()) {
			if t[nid].Com.GetCompareValue() < t[lid].Com.GetCompareValue() {
				t[nid].Com, t[lid].Com = t[lid].Com, t[nid].Com
				nid = lid
				rid = t[lid].right
				lid = t[lid].left
			} else {
				break
			}

		} else {
			if t[nid].Com.GetCompareValue() < t[rid].Com.GetCompareValue() {
				t[nid].Com, t[rid].Com = t[rid].Com, t[nid].Com
				nid = rid
				lid = t[rid].left
				rid = t[rid].right

			} else {
				break
			}
		}
	}
	return t
}

func insertNode(tree []heapNode, n heapNode, iid int) (int, []heapNode) {
	rid := 0
	tree = append(tree, n)
	if iid >= 0 {
		nowi := len(tree) - 1
		tree[nowi].pre = iid
		if tree[iid].left == -1 {
			tree[iid].left = nowi
			rid = iid
		} else if tree[iid].right == -1 {
			tree[iid].right = nowi
			rid = iid + 1
		} else {
			fmt.Printf("Something ERROR %v with node %v\n", tree, n)
		}
		tree = decideNodeUp(nowi, tree)
	}
	// fmt.Printf("Tree %v\n", tree)
	return rid, tree
}

func removeTopNode(t []heapNode) (heapNode, []heapNode) {
	if len(t) == 1 {
		return t[0], nil
	}
	top := t[0]
	fin := t[len(t)-1]
	if t[fin.pre].left == (len(t) - 1) {
		t[fin.pre].left = -1
	} else if t[fin.pre].right == (len(t) - 1) {
		t[fin.pre].right = -1
	}

	t = t[:len(t)-1]
	t[0].Com = fin.Com
	t = decideNodeDown(t)
	// fmt.Printf("RTree %v\n", t)
	return top, t
}

func GoHeapSort(array []SortInterf) []SortInterf {
	var fin []SortInterf
	var darr []heapNode
	iid := -1
	for ii := 0; ii != len(array); ii++ {
		tmp := heapNode{
			Com:   array[ii],
			pre:   -1,
			left:  -1,
			right: -1,
		}
		iid, darr = insertNode(darr, tmp, iid)
	}
	var tmp heapNode
	for len(darr) > 0 {
		tmp, darr = removeTopNode(darr)
		fin = append(fin, tmp.Com)
	}
	return fin
}
