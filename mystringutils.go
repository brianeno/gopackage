package util

import (
	strUtil "github.com/agrison/go-commons-lang/stringUtils"
)

func ReverseString(s string) string {
	return strUtil.Reverse(s)
}

func UpperCase(s string) string {
	return strUtil.UpperCase(s)
}
