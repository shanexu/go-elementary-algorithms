package isort

func ISort(xs []int) {
	for i := range xs {
		for j := i - 1; j >= 0 && xs[j+1] < xs[j]; j-- {
			xs[j], xs[j+1] = xs[j+1], xs[j]
		}
	}
}

func BinarySearch(xs []int, x int) int {
	l := 0
	u := len(xs)
	for l < u {
		m := (l + u) / 2
		if xs[m] == x {
			return m
		} else if xs[m] < x {
			l = m + 1
		} else {
			u = m
		}
	}
	return l
}

func ISortB(xs []int) {
	for i := range xs {
		x := xs[i]
		p := BinarySearch(xs[:i], x)
		for j := i; j >= 1; j-- {
			xs[j] = xs[j-1]
		}
		xs[p] = x
	}
}
