package encrypt

import (
	"bytes"
	"fmt"
)

type PaddingMode string

const PKCS5 PaddingMode = "PKCS5"
const PKCS7 PaddingMode = "PKCS7"
const ZEROS PaddingMode = "ZEROS"

func Padding(padding PaddingMode, src []byte, blockSize int) []byte {
	switch padding {
	case PKCS5:
		src = PKCS5Padding(src, blockSize)
	case PKCS7:
		src = PKCS7Padding(src, blockSize)
	case ZEROS:
		src = ZerosPadding(src, blockSize)
	}
	return src
}

func UnPadding(padding PaddingMode, src []byte) ([]byte, error) {
	switch padding {
	case PKCS5:
		return PKCS5UnPadding(src)
	case PKCS7:
		return PKCS7UnPadding(src)
	case ZEROS:
		return ZerosUnPadding(src)
	}
	return src, nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	return PKCS7Padding(src, blockSize)
}

func PKCS5UnPadding(src []byte) ([]byte, error) {
	return PKCS7UnPadding(src)
}

func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, paddingText...)
}

func PKCS7UnPadding(src []byte) ([]byte, error) {
	length := len(src)
	if length == 0 {
		return src, fmt.Errorf("src length is 0")
	}
	unPadding := int(src[length-1])
	if length < unPadding {
		return src, fmt.Errorf("src length is less than unPadding")
	}
	return src[:(length - unPadding)], nil
}

func ZerosPadding(src []byte, blockSize int) []byte {
	rem := len(src) % blockSize
	if rem == 0 {
		return src
	}
	return append(src, bytes.Repeat([]byte{0}, blockSize-rem)...)
}

func ZerosUnPadding(src []byte) ([]byte, error) {
	for i := len(src) - 1; ; i-- {
		if src[i] != 0 {
			return src[:i+1], nil
		}
	}
}
