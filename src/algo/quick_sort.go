package algo

import "fmt"

func Quick_sort(ulist []int) []int {
	if len(ulist) <= 1 {
		return ulist
	}
	mid := len(ulist) / 2
	fmt.Println(mid)
	left := []int{}
	right := []int{}
	equals := []int{}

	for _, i := range ulist {
		if i < ulist[mid] {
			left = append(left, i)
		} else if i > ulist[mid] {
			right = append(right, i)
		} else {
			equals = append(equals, i)
		}
	}

	return append(append(Quick_sort(left), equals...), Quick_sort(right)...)
}
