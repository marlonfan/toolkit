package util

import (
	"strings"
)

// some code for govalidator helper

// ValidateErrorMessage 验证错误信息格式
type ValidateErrorMessage map[string]interface{}

// ParseValidateString 为验证错误结果格式化信息
func ParseValidateString(err string) ValidateErrorMessage {
	errorMessage := make(ValidateErrorMessage)
	arrArray := strings.Split(err, ";")
	for _, i := range arrArray {
		errItem := strings.Split(i, ": ")
		errorMessage[errItem[0]] = errItem[1]
	}
	return errorMessage
}
