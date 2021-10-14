package mathematics

// FindGreatestCommonDivisor Iterative implementation of Euclidian algorithm.
// Algorithm formula is following
// gcd(num1,num2) = gcd(min(num1,num2),max(num1,num2)−min(num1,num2))
func FindGreatestCommonDivisor(num1, num2 int) int {
	var temp int
	for num1 > 0 {
		if num1 < num2 {
			temp = num1
			num1 = num2
			num2 = temp
		}
		num1 = num1 - num2
	}
	return num2
}
