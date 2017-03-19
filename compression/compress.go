package compress

import "strconv"

func compress(uncompressed string) string {
	var buf string

	if len(uncompressed) == 0 {
		return buf
	}

	var prev byte
	letter := uncompressed[0]
	last := 0
	for i := 1; i < len(uncompressed); i++ {
		prev = uncompressed[i-1]
		letter = uncompressed[i]
		if letter != prev {
			buf += string(prev)
			buf += strconv.Itoa(i - last)
			last = i
		}
	}
	buf += string(letter)
	buf += strconv.Itoa(len(uncompressed) - last)

	return buf
}
