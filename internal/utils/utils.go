package utils

func ContainsTag(tags []string, tag string) bool {
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return false
}

func RemoveDuplicates(s []string) []string {
	m := make(map[string]struct{})
	result := []string{}

	for _, v := range s {
		if _, value := m[v]; !value {
			m[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

func In[T comparable](target T, arr []T) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}
