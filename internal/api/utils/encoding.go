package utils

import (
	"strings"
)

func EncodeToBase62(deci int) string {
	s := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	hashStr := ""
	for deci > 0 {
		hashStr = string(s[deci%62]) + hashStr
		deci /= 62
	}
	return hashStr
}

func DecodeFromBase62(hashStr string) int {
	s := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	deci := 0
	for _, c := range hashStr {
		deci = deci*62 + strings.Index(s, string(c))
	}
	return deci
}

func DecodeFromURL(shortURL string) int {
	return DecodeFromBase62(shortURL)
}
