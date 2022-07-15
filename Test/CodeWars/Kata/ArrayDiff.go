package main

func ArrayDiff(a, b []int) []int {

	for _, vb := range b {
		a = removeSliceElement(vb, a)
	}
	return a
}

func removeSliceElement(element int, slice []int) []int {
	for i := range slice {
		if slice[i] == element {
			copy(slice[i:], slice[i+1:])
			slice[len(slice)-1] = 0
			slice = slice[:len(slice)-1]
			return removeSliceElement(element, slice)
		}
	}
	return slice
}
