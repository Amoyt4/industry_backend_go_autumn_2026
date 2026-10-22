package main

import "strings"

func greet(name string) string {
	str := strings.TrimSpace(name)
	if str == "" {
		return "Hello, World!"
	} else {
		return "Hello, " + str + "!"
	}
}
