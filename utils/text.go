package utils

import (
	"regexp"
	"strings"
	"unicode"
)

func IsChineseChar(str string) bool {
	//https://en.wikipedia.org/wiki/Script_(Unicode)
	//。；，：“”（）、？《》
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}

func IsEnglishChar(text string) bool {
	// * ASCII码
	// * A-Z 65-90 91 [ 92 \ 93 ] 94^ 95_96` a-z 97-122
	// * 33-47 标点符号 0-9 48-64
	// https://tool.oschina.net/commons?type=4
	var result []bool
	var i = 0
	for _, r := range compressStr(strings.TrimSpace(text)) {
		if r > 32 && r <= 126 {
			result = append(result, true)
			i++
			continue
		}
		result = append(result, false)
	}
	if i == len(result) {
		return true
	}
	return false

}

func compressStr(str string) string {
	if str == "" {
		return ""
	}
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}
