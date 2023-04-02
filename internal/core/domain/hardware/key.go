package hardware

import (
	"strconv"
	"strings"
)

type Key string

const (
	AlphabetKey = "ABCDabcd"
	NumberKey   = "0123456789"
	StarKey     = "*"
	SharpKey    = "#"
	ALPHABET    = "alphabet"
	NUMBER      = "number"
	SHARP       = "sharp"
	STAR        = "star"
	UNKNOWN     = "unknown"
)

func (k Key) Type() string {
	if k.IsAlphabet() {
		return ALPHABET
	} else if k.IsNumber() {
		return NUMBER
	} else if k.IsSharp() {
		return SHARP
	} else if k.IsStar() {
		return STAR
	} else {
		return UNKNOWN
	}
}

func (k Key) ToString() string {
	return strings.ToLower(string(k))
}

func (k Key) ToNumber() int {
	intVar, _ := strconv.Atoi(string(k))
	return intVar
}

func (k Key) IsAlphabet() bool {
	return strings.Contains(AlphabetKey, string(k))
}

func (k Key) IsNumber() bool {
	return strings.Contains(NumberKey, string(k))
}

func (k Key) IsStar() bool {
	return k == StarKey
}

func (k Key) IsSharp() bool {
	return k == SharpKey
}
