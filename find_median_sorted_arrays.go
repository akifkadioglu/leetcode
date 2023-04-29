package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var list []int
	var result float64
	list = append(list, nums1...)
	list = append(list, nums2...)
	for i, _ := range list {
		for j := i; j < len(list); j++ {
			if list[i] > list[j] {
				tmp := list[j]
				list[j] = list[i]
				list[i] = tmp
			}
		}
	}
	lenght := len(list)
	if lenght%2 == 1 {
		result = float64(list[lenght/2])
	} else {
		result = float64(list[lenght/2]+list[(lenght/2)-1]) / 2
	}
	return result
}
