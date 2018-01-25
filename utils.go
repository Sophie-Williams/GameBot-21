package main

import "strings"

func stringInSlice(str string, list []string) bool {
	for _, elem := range list {
		if strings.Compare(elem, str) == 0 {
			return true
		}
	}
	return false
}
