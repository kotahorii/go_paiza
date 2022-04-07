package pythonrdy

func OrganizeList(unOrganizedList [][]string) map[string][]string {
	// if not define result here, nil map error occurred
	result := map[string][]string{}
	for _, l := range unOrganizedList {
		_, ok := result[l[0]]
		if ok {
			result[l[0]] = append(result[l[0]], l[1])
		} else {
			result[l[0]] = []string{l[1]}
		}
	}
	return result
}
