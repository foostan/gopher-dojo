package try

func swap(n, m int) (int, int) {
	return m, n
}

func swap2(n, m *int) {
	tmp := *m
	*m = *n
	*n = tmp
}

func Swapable() {
	n, m := swap(10, 20)
	println(n, m)

	swap2(&n, &m)
	println(n, m)
}
