package mylib

func Average(s []int) int {
	sum := 0
	for _, val := range s {
		sum += val
	}
	return int(sum / len(s))
}
