package mbutils

func EnumMapToList(mp map[int32]string) []string {
	arr := []string{}
	for _, v := range mp {
		arr = append(arr, v)
	}

	return arr
}
