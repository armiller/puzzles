package compress

import "bytes"
import "strconv"

func compress(uncompressed string) string {
	var buf bytes.Buffer

	switch len(uncompressed) {
	case 0:
		break
	case 1:
		buf.WriteByte(uncompressed[0])
		buf.WriteString("1")
	default:
		var prev byte
		var letter byte
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
	}

	return buf.String()
}
