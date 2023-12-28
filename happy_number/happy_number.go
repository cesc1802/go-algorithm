package main

func digit(num int) int {
	if num < 10 {
		return num
	}
	return num % 10
}

func sumSquareOfNum(num int) int {
	sum := 0
	for num > 0 {
		d := digit(num)
		sum += d * d
		num = num / 10
	}
	return sum
}
func isHappy(num int) bool {
	slow := num
	fast := sumSquareOfNum(num)
	for fast != 1 && slow != fast {
		slow = sumSquareOfNum(slow)
		fast = sumSquareOfNum(sumSquareOfNum(fast))
	}
	if fast == 1 {
		return true
	}
	return false
}
