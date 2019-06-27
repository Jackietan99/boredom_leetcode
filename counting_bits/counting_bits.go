package counting_bits

import "strconv"

func decimal2Binary(n int) string {
	b := ""
	for {
		// remain
		r := 0
		n, r = n>>1, n%2
		b = strconv.Itoa(r) + b
		if 0 == n {
			return b
		}
	}
}

func countOne(s string) int {
	c := 0
	for _, i := range []rune(s) {
		if '1' == i {
			c++
		}
	}
	return c
}

/**
避免对递增数组中的每个数值作计算，将4位看做一个单元，单元内0-15的二进制数中1的个数是确定的。
这样采用16进制去计算，给定数值，每除以16所得的余数就是落在该单元内的数值，直至被除数为0，将每个单元中1的个数累加既可。
*/
func countBits(num int) []int {
	s := make([]int, num+1)
	for i := 0; i <= num; i++ {
		s[i] = countOne(decimal2Binary(i))
	}
	return s
}
