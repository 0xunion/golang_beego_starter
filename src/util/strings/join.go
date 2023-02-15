package strings

import (
	"fmt"
	"strings"
)

func StringJoin(strs ...string) string {
	n := 0
	for i := 0; i < len(strs); i++ {
		n += len(strs[i])
	}

	var b strings.Builder
	b.Grow(n)
	for _, s := range strs {
		b.WriteString(s)
	}
	return b.String()
}

func StringJoinWithDelim(delim string, strs ...string) string {
	n := 0
	for i := 0; i < len(strs); i++ {
		n += len(strs[i]) + len(delim)
	}

	var b strings.Builder
	b.Grow(n)
	for i, s := range strs {
		if i > 0 {
			b.WriteString(delim)
		}
		b.WriteString(s)
	}
	return b.String()
}

type NumberConstraint interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func NumberJoin[T NumberConstraint](delim string, slice []T) string {
	length := len(slice)
	if length == 0 {
		return ""
	}

	var finallen int
	var b strings.Builder
	string_slice := make([]string, length)
	string_slice[0] = fmt.Sprintf("%v", slice[0])
	for i := 0; i < length; i++ {
		string_slice[i] = fmt.Sprintf("%v", slice[i])
		finallen += len(string_slice[i]) + len(delim)
	}

	b.Grow(finallen)
	// write first element
	b.WriteString(string_slice[0])
	for i := 1; i < length; i++ {
		b.WriteString(delim)
		b.WriteString(string_slice[i])
	}

	return b.String()
}
