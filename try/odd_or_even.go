package try

func IsEven(n int) bool {
	return n % 2 == 0
}

func OddOrEven() {
	for i := 1; i <= 100; i++ {
		print(i)
		if IsEven(i) {
			println("-偶数")
		} else {
			println("-奇数")
		}
	}
}
