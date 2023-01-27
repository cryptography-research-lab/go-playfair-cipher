package playfair_cipher

import (
	"github.com/golang-infrastructure/go-tuple"
)

type Options struct {
	BlackCharacter rune
}

func Encrypt(plaintext, key string) string {
	panic("")
}

// GroupToPairForEncrypt 把字符串分组为一对一对的
func GroupToPairForEncrypt(s string) []*tuple.Tuple2[rune, rune] {
	//
}
