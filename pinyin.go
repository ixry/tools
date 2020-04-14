package tools

import (
	"bytes"
	"strings"

	pinyin "github.com/mozillazg/go-pinyin"
)

// Pinyin 拼音
func Pinyin(src string) string {

	return strings.Join(pinyin.LazyConvert(src, nil), "")
}

// PinyinFirst 只取拼音首字母
func PinyinFirst(src string) string {
	py := pinyin.LazyConvert(src, nil)

	var buf bytes.Buffer
	for i, l := 0, len(py); i < l; i++ {
		buf.WriteString(py[i][:1])
	}

	return buf.String()
}

// StringToPy 汉字数字字母串转拼音串
func StringToPy(src string) string {
	a := pinyin.NewArgs()
	a.Fallback = func(r rune, a pinyin.Args) []string {
		return []string{string(r)}
	}

	arr := pinyin.Pinyin(src, a)

	var ret bytes.Buffer
	for k := range arr {
		ret.WriteString(arr[k][0])
	}

	return ret.String()
}
