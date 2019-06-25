package reverse_integer

import "fmt"

func main() {

	num := -2147483412

	result := reverse(num)

	fmt.Println(result)

}

func reverse(x int) int {
	var min int = 0x8000000
	var max int = 0x7FFFFFF
	if x == 0 {
		return 0
	}
	sum := 0
	for {
		leftDigits := x / 10
		lastDigits := x % 10
		x = leftDigits

		sum = sum*10 + lastDigits

		if leftDigits == 0 {
			break
		}
	}

	if sum < -min || sum > max {
		sum = 0
	}

	return sum

}
