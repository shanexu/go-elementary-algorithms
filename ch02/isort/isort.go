package isort

func ISort(xs []int) {
	for i := range xs {
		for j := i - 1; j >= 0 && xs[j+1] < xs[j]; j-- {
			xs[j], xs[j+1] = xs[j+1], xs[j]
		}
	}
}
