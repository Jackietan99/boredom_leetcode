package companition

//    root
//                     /  |  \  \
//                    /   |   \   \
//                   /    |    \    \
//                  1     2     3    4
//                / | \   | \    \    -
//               /  |  \  |  \    \
//              2   3   4  3  4    4
//             /    /    -  /  -   -
//            3    4      -  4
//            |    |
//            4    -

func combine(n int, k int) [][]int {

	var (
		res = make([][]int, 0)
	)

}

//combinations are unordered, i.e., [1,2] and [2,1] are considered to be the same combination.

func Travis(k, n, start int, res *[][]int) {

	if start == k {

	}
	for a := start; a <= n; a++ {
		var temp = make([]int, 0)
		temp = append(temp, a)
		Travis(k, n, a+1, &temp)
	}
}
