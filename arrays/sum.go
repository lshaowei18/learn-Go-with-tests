package main

func main() {

}

func Sum(arr []int) (sum int) {
	for _, n := range arr {
		sum += n
	}
	return
}

func SumAll(slices ...[]int) []int {

	combined := make([]int, len(slices))
	for sliceIndex, slice := range slices {
		combined[sliceIndex] = Sum(slice)
	}
	return combined
}

//Slower because of append
func SumAll2(slices ...[]int) (combined []int) {

	for _, slice := range slices {
		combined = append(combined, Sum(slice))
	}
	return
}

func SumTails(slices ...[]int) []int {

	combined := make([]int, len(slices))
	for i, slice := range slices {
		for j := 1; j < len(slice); j++ {
			combined[i] += slice[j]
		}
	}

	return combined
}

func SumTails2(slices ...[]int) []int {

	combined := make([]int, len(slices))
	for i, slice := range slices {
		tails := slice[1:]
		combined[i] = Sum(tails)
	}

	return combined
}
