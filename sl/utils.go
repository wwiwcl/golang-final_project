package sl

func inSliceString(e []string, slice []string) int {
	i := 0
	for _, s := range slice {
		for _, v := range e {
			if v == s {
				return i
			}
		}
		i++
	}
	return -1
}
