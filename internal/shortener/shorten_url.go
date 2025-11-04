package shortener

import (
	"math/rand"
	"os"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func ShortenUrl(url string) string {
	var b strings.Builder
	b.WriteString("http://" + os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/")
	b.WriteString(randStringBytes(6))
	return b.String()
}
