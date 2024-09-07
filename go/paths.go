package main

import "strings"

func SplitExt(path string) (string, string) {
	i := strings.LastIndex(path, ".")

	if i == -1 {
		return path, ""
	}
	return path[:i], path[i:]
}
