package main

func quickSortGo(a []int, lo, hi int, done chan struct{}, depth int) {
	if lo >= hi {
		done <- struct{}{}
		return
	}
	depth--
	p := partition(a, lo, hi)
	if depth > 0 {
		childDone := make(chan struct{}, 2)
		go quickSortGo(a, lo, p-1, childDone, depth)
		go quickSortGo(a, p+1, hi, childDone, depth)
		<-childDone
		<-childDone
	} else {
		quickSort(a, lo, p-1)
		quickSort(a, p+1, hi)
	}
	done <- struct{}{}
}

func quickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func partition(data []int, a, b int) int {
	pivot := data[b]
	i := a
	for j := a; j < b; j++ {
		if data[j] < pivot {
			data[j], data[i] = data[i], data[j]
			i++
		}
	}
	data[i], data[b] = data[b], data[i]
	return i
}
