package mathematics

//GCDIterative
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
