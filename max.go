package max

func Find(slice []interface{}, greater func(i, j int) bool) interface{} {
	if len(slice) == 0 {
		return nil
	}

	maxI := 0
	for i := range slice {
		if greater(i, maxI) {
			maxI = i
		}
	}
	return slice[maxI]
}
