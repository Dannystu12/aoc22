package day6

import "unicode"

type Signal string

func (s Signal) IsValid() bool {
	for _, c := range s {
		if !unicode.IsLetter(c) || !unicode.IsLower(c) {
			return false
		}
	}

	return true
}

func (s Signal) GetStartOfPacketMarker() (int, bool) {
	return s.GetMarker(4)
}

func (s Signal) GetStartOfMessageMarker() (int, bool) {
	return s.GetMarker(14)
}

func (s Signal) GetMarker(num uint) (int, bool) {

	n := int(num)

	if !s.IsValid() {
		return 0, false
	}

	if len(s) < n {
		return 0, false
	}

	for i := 0; i < len(s)-(n-1); i++ {
		var chars = make(map[byte]bool)
		for j := i; j < i+n; j++ {
			chars[s[j]] = true
		}

		if len(chars) == n {
			return i + n, true
		}
	}

	return 0, false
}
