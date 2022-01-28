package book1Lesson27

import (
	"fmt"
	"sort"
)

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool {
			return s[i] < s[j]
		}
		sort.Slice(s, less)
	}
}

func pracSort() {
	fmt.Println("----------")
	fmt.Println("sort.go")
	fmt.Println("----------")

	food := []string{"onion", "carrot", "celery"}
	sortStrings(food, nil)
	fmt.Println(food)

}
