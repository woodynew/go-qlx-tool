package utils

import "sort"

func InArray(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	if index < len(str_array) && str_array[index] == target {
		return true
	}
	return false
}

func AddArrayEx(str_array []string, target string) []string {
	if !InArray(target, str_array) {
		str_array = append(str_array, target)
	}
	return str_array
}
