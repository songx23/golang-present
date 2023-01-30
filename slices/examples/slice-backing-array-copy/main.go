package slice_backing_array_copy

import "strings"

func mapWithNilSlice(a []string) []string {
	var b []string
	for _, v := range a {
		b = append(b, strings.ToUpper(v))
	}
	return b
}

func mapWithPredefinedCapacity(a []string) []string {
	b := make([]string, 0, len(a))
	for _, v := range a {
		b = append(b, strings.ToUpper(v))
	}
	return b
}

func mapWithPredefinedLength(a []string) []string {
	b := make([]string, len(a))
	for i := range a {
		b[i] = strings.ToUpper(a[i])
	}
	return b
}
