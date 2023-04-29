package main

func longestPalindrome(s string) string {
	var polyndrome []string
	isPolyndrome := func(tmp string, x int, y int) bool {
		for a := 0; a < len(tmp); a++ {
			if tmp[x] == tmp[y] {
				x += 1
				y -= 1
			} else {
				return false
			}
		}
		return true
	}
	for i := 0; i < len(s); i++ {
		for j := i; j <= len(s); j++ {
			tmp := s[i:j]
			x := 0
			y := len(tmp) - 1
			if isPolyndrome(tmp, x, y) {
				polyndrome = append(polyndrome, tmp)
			}
		}
	}
	result := ""
	longest := 0
	for _, v := range polyndrome {
		if longest < len(v) {
			result = v
			longest = len(v)
		}
	}
	return result
}