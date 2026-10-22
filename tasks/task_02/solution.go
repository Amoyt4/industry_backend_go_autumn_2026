package main

func rotateRunes(s string, shift int) string {

	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	n := len(runes)

	shiftCount := shift % n
	if shiftCount < 0 {
		shiftCount += n
	}

	if shiftCount == 0 {
		return string(runes)
	}

	result := make([]rune, n)

	copy(result, runes[shiftCount:])
	copy(result[n-shiftCount:], runes[:shiftCount])

	return string(result)
}
