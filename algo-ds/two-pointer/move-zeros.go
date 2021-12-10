package main

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func moveZeroes(nums []int) {
	j := 0 //keep track zeros pos

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 2, 3}
	moveZeroes(nums)
}
