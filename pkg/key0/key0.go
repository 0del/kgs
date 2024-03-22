package key0

import (
	"crypto/rand"
	"fmt"
)

var defaultBase62 = []byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o',
	'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D',
	'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S',
	'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '0', '1', '2', '3', '4', '5', '6', '7',
	'8', '9',
}

const defaultLen = 8

type Generator interface {
	New() (string, error)
}

type Options func(generator)

func WithLen(l int8) Options {
	return func(g generator) {
		g.len = l
	}
}
func WithChars(chars []byte) Options {
	return func(g generator) {
		g.chars = chars
	}
}

func NewGen(opts ...Options) Generator {
	g := generator{}

	g.len = defaultLen
	g.chars = defaultBase62

	for _, ops := range opts {
		ops(g)
	}
	return g
}

type generator struct {
	len   int8
	chars []byte
}

func (g generator) New() (string, error) {
	return g.newChars()
}

func (g generator) newChars() (string, error) {
	charsLen := len(g.chars)
	uri := make([]byte, g.len)
	uriIndex := make([]byte, g.len)

	_, err := rand.Read(uriIndex)
	if err != nil {
		return "", fmt.Errorf("generate error: %s", err.Error())
	}

	for k, v := range uriIndex {
		index := int(v)
		uri[k] = g.chars[index%charsLen]
	}

	return string(uri), nil
}
