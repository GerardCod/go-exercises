package recursion

//MCD regresa el maximo común divisor de dos números.
func MCD(n, m int) int {
	if n == m {
		return n
	}

	if n > m {
		return MCD(n-m, m)
	}
	return MCD(n, m-n)
}
