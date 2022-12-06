package day6

import "unicode"

type Signal string

func (s Signal) IsValid() bool {

	if len(s) < 4 {
		return false
	}

	for _, c := range s {
		if !unicode.IsLetter(c) || !unicode.IsLower(c) {
			return false
		}
	}

	return true
}

func (s Signal) GetMarker() (int, bool) {
	if !s.IsValid() {
		return 0, false
	}

	for i := 0; i < len(s)-3; i++ {
		var chars = make(map[byte]bool)
		for j := i; j < i+4; j++ {
			chars[s[j]] = true
		}
		if len(chars) == 4 {
			return i + 4, true
		}
	}

	return 0, false
}
