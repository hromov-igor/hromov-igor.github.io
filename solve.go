package main


func RemoveEven(input []int) []int {
	res := []int{}
	for _, item := range input {
		if item % 2 == 1 {
			res = append(res, item)
		}
	}
	return res
}
