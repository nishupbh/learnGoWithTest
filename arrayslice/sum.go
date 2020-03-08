package arrayslice

func sum(val []int) int {
	sumValue := 0
	for i := 0; i < len(val); i++ {
		sumValue += val[i]
	}
	return sumValue
}

func sumAll(val ...[]int) []int {
	lengthOfNumbers := len(val)
	sums := make([]int, lengthOfNumbers)
	for i, number := range val {
		sums[i] = sum(number)
	}
	return sums
}

func sumAllTrails(numbersToSum ...[]int) []int {
	sums := []int{}

	for _, number := range numbersToSum {

		if len(number) == 0 {
			sums = append(sums, 0)
		} else {
			tail := number[1:]
			sums = append(sums, sum(tail))
		}
	}
	return sums

}
