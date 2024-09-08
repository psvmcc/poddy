package helpers

type ByKey [][]string

func (a ByKey) Len() int {
	return len(a)
}

func (a ByKey) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}

func (a ByKey) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
