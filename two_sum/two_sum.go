package two_sum

import "fmt"

func main() {
	given := []int{2, 7, 11, 15}
	result := twoSum(given, 9)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	//var result [2]int
	var num1 int
	var num2 int
	m := make(map[int]int)
	for x := 0; x < len(nums); x++ {
		_, isExist := m[target-nums[x]]
		if isExist {
			num1 = m[target-nums[x]]
			num2 = x
			break
		}
		m[nums[x]] = x
	}
	return []int{num1, num2}
	//for x:=0;x<len(nums);x++{
	//	for k,v :=range nums[x:]{
	//		if (nums[x]+ v) == target{
	//			result[0]=x
	//			result[1]=k
	//			break
	//		}
	//	}
	//}
	//return result
}
