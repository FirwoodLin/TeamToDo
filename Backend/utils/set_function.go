package utils

func Intersect(slice1, slice2 []uint) []uint {
	m := make(map[uint]int)
	nn := make([]uint, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}
