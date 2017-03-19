package compress

import "fmt"

func compress(uncompressed string) string {
	var buf string

	if len(uncompressed) == 0 {
		return buf
	}

	letter := string(uncompressed[0])
	count := 1
	for _, next := range uncompressed[1:] {
		if letter == string(next) {
			count++
		} else {
			buf += fmt.Sprintf("%s%d", letter, count)
			count = 1
		}
		letter = string(next)
	}
	buf += fmt.Sprintf("%s%d", letter, count)

	return buf
}
