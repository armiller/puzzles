package compress

import "bytes"
import "strconv"

func compress(uncompressed string) string {
	var buf bytes.Buffer

	if len(uncompressed) == 0 {
		return buf.String()
	}

	var prev byte
	letter := uncompressed[0]
	last := 0
	for i := 1; i < len(uncompressed); i++ {
		prev = uncompressed[i-1]
		letter = uncompressed[i]
		if letter != prev {
			buf.WriteByte(prev)
			buf.WriteString(strconv.Itoa(i - last))
			last = i
		}
	}
	buf.WriteByte(letter)
	buf.WriteString(strconv.Itoa(len(uncompressed) - last))

	return buf.String()
}
