package sum

func ArraySum(numbers [5]int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SliceSum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {

	// lengthOfNumbers := len(numbersToSum)
	// sums = make([]int, lengthOfNumbers)
	//
	// for i, numbers := range numbersToSum {
	// 	sums[i] = SliceSum(numbers)
	// }

	for _, numbers := range numbersToSum {
		sums = append(sums, SliceSum(numbers))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:] // take from 1 to the end
			sums = append(sums, SliceSum(tail))
		}
	}
	return
}

func Sum(numbersToSum ...int) (sum int) {
	for _, number := range numbersToSum {
		sum += number
	}
	return
}
