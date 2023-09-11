package utils

import "golang.org/x/crypto/bcrypt"

// 加密字符串
func Encrypt(str string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(result), err
}
