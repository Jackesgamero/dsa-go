package movezeroes

import "fmt"

func main() {
	nums := []int{0, 1, 0, 2, 0, 3, 0, 4, 0, 5}

	n := len(nums)
	posFirstZero := n

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			posFirstZero = i
			break
		}
	}

	for i := posFirstZero; i < n; i++ {
		if nums[i] != 0 {
			nums[i], nums[posFirstZero] = nums[posFirstZero], nums[i]
			posFirstZero++
		}
	}

	fmt.Println(nums)
}
