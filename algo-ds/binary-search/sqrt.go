package main

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	l := 1
	r := x / 2
	for l <= r {
		mid := l + (r-l)/2

		if mid*mid == x {
			return mid
		}

		if mid*mid > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l - 1
}

