package stringse

func Diff(listA, listB []string) []string {
	listD := make([]string, 0)
	for _, sA := range listA {
		contains := false
		for _, sB := range listB {
			if sA == sB {
				contains = true
				break
			}
		}
		if !contains {
			listD = append(listD, sA)
		}
	}

	return listD
}
