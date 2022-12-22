package filter

// Liefert eine Liste mit allen Elementen aus list, die kleiner oder gleich key sind.
func FilterLess(list []int, key int) []int {
	result := []int{}
	if len(list) == 0 {
		return result
	}
	head := list[0]
	tail := list[1:]
	if head <= key {
		result = append(result, head)
	}
	return append(result, FilterLess(tail, key)...)
}

// Liefert eine Liste mit allen Elementen aus list, die größer als key sind.
func FilterGreater(list []int, key int) []int {
	//TODO programmiere die Funktion FilterGreater
	result := []int{}
	return result
}
