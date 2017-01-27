package algorithm

func BinarySearch(v int, target []int) bool {
	found := false

	start := 0
	end := len(target) - 1

	for start <= end {
		mid := (start + end) / 2
		if target[mid] == v {
			found = true
			break
		} else if target[mid] > v {
			end = mid - 1
		} else if target[mid] < v {
			start = mid + 1
		}
	}
	return found
}
