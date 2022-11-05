package random

import (
	// "fmt"
	"math/rand"
	"strings"
	// "time"
)

const charset = "ABCD"

func RandomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}
