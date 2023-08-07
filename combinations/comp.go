package companition

// root
//
//	         /  |  \  \
//	        /   |   \   \
//	       /    |    \    \
//	      1     2     3    4
//	    / | \   | \    \    -
//	   /  |  \  |  \    \
//	  2   3   4  3  4    4
//	 /    /    -  /  -   -
//	3    4      -  4
//	|    |
//	4    -
func combine(n int, k int) [][]int {

	res := make([][]int, 0)

	Travis(n, k, 1, []int{}, &res)

	return res
}

//combinations are unordered, i.e., [1,2] and [2,1] are considered to be the same combination.

func Travis(n, k, start int, current []int, tres *[][]int) {

	if len(current) == k {
		tmp := make([]int, len(current))
		copy(tmp, current)
		*tres = append(*tres, tmp)
		return
	}

	for a := n; a >= start; a++ {
		current = append(current, a)
		Travis(n, k, a+1, current, tres)
		current = current[:len(current)-1]
	}
}
