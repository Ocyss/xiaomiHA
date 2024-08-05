package utils

import (
	"time"
)

func SliceGroup[T any](slice []T, avgTime *time.Duration, cell func(i, j int, t int64) bool) [][]T {
	n := len(slice)
	group := make([][]T, 0, n/2)
	var (
		i, j int
		t    int64
	)
	if avgTime != nil {
		t = int64(*avgTime / time.Second)
	} else {
		t = int64(time.Hour * 24 / time.Second)
	}
	for {
		if i >= n {
			break
		}
		for j = i + 1; j < n && cell(i, j, t); j++ {

		}
		group = append(group, slice[i:j])
		i = j
	}
	return group
}
