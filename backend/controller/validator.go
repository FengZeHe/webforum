package controller

import (
	ut "github.com/go-playground/universal-translator"
	"strings"
)

// 定义一个全局翻译器T
var trans ut.Translator

// 定义一个去掉结构体名称前缀的自定义方法：
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
