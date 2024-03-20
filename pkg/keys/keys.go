package keys

var base62 = []byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o',
	'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D',
	'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S',
	'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '0', '1', '2', '3', '4', '5', '6', '7',
	'8', '9',
}

type Alphabet struct {
	alphabet []byte
}

func (a Alphabet) Gen(keySize uint, keysOut chan<- string, end chan<- bool) {
	defer func() {
		close(keysOut)
		end <- true
	}()
	if keySize == 0 {
		return
	}

	var chars []byte
	recKey(a.alphabet, chars, keySize, keysOut)
}

func recKey(alphabet []byte, chars []byte, remaining uint, keys chan<- string) {
	if remaining == 0 {
		key := string(chars)
		keys <- key
		return
	}

	for _, char := range alphabet {
		chars = append(chars, char)
		recKey(alphabet, chars, remaining-1, keys)
		chars = chars[:len(chars)-1]
	}
}

func NewAlphabet() Alphabet {
	return Alphabet{
		alphabet: base62,
	}
}
