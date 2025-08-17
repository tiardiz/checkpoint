package piscine

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func FindPrevPrime(nb int) int {

	if nb < 2 {
		return 0
	}

	for i := nb; i >= 2; i-- {
		if isPrime(i) {
			return i
		}
	}
	return 0
}
