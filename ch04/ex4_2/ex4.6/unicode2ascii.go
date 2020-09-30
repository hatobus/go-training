package u2a

import (
	"unicode"
)

func Unicode2AsciiTrimSpace(unicodeBytes []byte) []byte {
	ascii := 0
	var pSpace bool

	for _, b := range unicodeBytes {
		space := unicode.IsSpace(rune(b))

		if !space {
			unicodeBytes[ascii] = b
			ascii++
		} else if !pSpace {
			unicodeBytes[ascii] = byte(' ')
			ascii++
		}
		pSpace = space
	}
	return unicodeBytes[:ascii]
}
